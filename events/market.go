package events

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mylakehead/sync/abis/market"
	"github.com/mylakehead/sync/lib"
	"github.com/mylakehead/sync/models"
	"github.com/mylakehead/sync/runtime"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContractEvent struct {
	runTime               *runtime.Runtime
	abi                   abi.ABI
	wssEndpoint           string
	wg                    *sync.WaitGroup
	dialInterval          time.Duration
	pullInterval          time.Duration
	requestTimeout        time.Duration
	MarketContractAddress common.Address
	indexFilePath         string
	step                  int64
	marketFilter          *market.MarketFilterer
}

func NewMarketEvent(wg *sync.WaitGroup, rt *runtime.Runtime, abs string) (*ContractEvent, error) {
	Abi, err := abi.JSON(strings.NewReader(market.MarketMetaData.ABI))
	if err != nil {
		return nil, err
	}

	// index file and cur
	indexFilePath := filepath.Join(abs, rt.Config.Contract.IndexFile)
	cur := rt.Config.Contract.FallbackIndex
	if lib.IsFileExists(indexFilePath) {
		content, err := os.ReadFile(indexFilePath)
		if err != nil {
			return nil, err
		}
		var index lib.Index
		err = json.Unmarshal(content, &index)
		if err != nil {
			return nil, err
		}
		cur = index.Cur
	} else {
		err := lib.WriteIndexTo(cur, indexFilePath)
		if err != nil {
			return nil, err
		}
		fmt.Printf("[event][market] record index: %v \n", cur)
	}

	// filter & parser
	marketFilter, err := market.NewMarketFilterer(
		common.HexToAddress(rt.Config.Contract.MarketAddress),
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &ContractEvent{
		runTime:               rt,
		abi:                   Abi,
		wssEndpoint:           rt.Config.Infura.Wss,
		wg:                    wg,
		dialInterval:          time.Duration(rt.Config.Contract.DialInterval) * time.Second,
		pullInterval:          time.Duration(rt.Config.Contract.PullInterval) * time.Second,
		requestTimeout:        time.Duration(rt.Config.Contract.RequestTimeout) * time.Second,
		MarketContractAddress: common.HexToAddress(rt.Config.Contract.MarketAddress),
		indexFilePath:         indexFilePath,
		step:                  rt.Config.Contract.Step,
		marketFilter:          marketFilter,
	}, nil
}

func (c *ContractEvent) ListenEvents() {
	defer c.wg.Done()

	for {
		// auto reconnect, don't worry
		client, err := ethclient.Dial(c.wssEndpoint)
		if err != nil {
			fmt.Printf("%v\n", err)
			time.Sleep(c.dialInterval)
			continue
		}

		for {
			time.Sleep(c.pullInterval)

			// get block header
			ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
			header, err := client.HeaderByNumber(ctx, nil)
			if err != nil {
				fmt.Printf("%v\n", err)
				cancel()
				continue
			}
			cancel()
			fmt.Printf("[event][market] block header: %v \n", header.Number.Int64())

			// get cur from file
			cur, err := lib.ReadIndexFrom(c.indexFilePath)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			// calculate <FromBlock> <ToBlock>
			var from *big.Int
			var to *big.Int
			h := header.Number.Int64()
			if cur == h {
				from = big.NewInt(cur)
				to = nil
				fmt.Printf("[event][market] pulling logs from %v to %v \n", from.Int64(), "latest")
			} else if cur+1+c.step > h {
				from = big.NewInt(cur + 1)
				to = nil
				fmt.Printf("[event][market] pulling logs from %v to %v \n", from.Int64(), "latest")
			} else {
				from = big.NewInt(cur + 1)
				to = big.NewInt(cur + c.step)
				fmt.Printf("[event][market] pulling logs from %v to %v \n", from.Int64(), to.Int64())
			}

			// pull logs
			query := ethereum.FilterQuery{
				FromBlock: from,
				ToBlock:   to,
				Addresses: []common.Address{
					c.MarketContractAddress,
				},
				Topics: [][]common.Hash{{
					common.HexToHash(c.runTime.Config.Contract.PurchasedEvent),
				}},
			}
			ctx, cancel = context.WithTimeout(context.Background(), c.requestTimeout)
			logs, err := client.FilterLogs(ctx, query)
			if err != nil {
				fmt.Printf("%v", err)
				cancel()
				continue
			}
			cancel()
			fmt.Printf("[event][market] %v logs pulled \n", len(logs))

			// no record pulled, go step
			if len(logs) <= 0 && to != nil {
				err = lib.WriteIndexTo(to.Int64(), c.indexFilePath)
				if err != nil {
					fmt.Printf("%v\n", err)
					continue
				}
				fmt.Printf("[event][market] record index: %v \n", to.Int64())
				continue
			}

			// handle and record the block until end
			for _, l := range logs {
				if cur == h && int64(l.BlockNumber) == h {
					fmt.Printf("[event][market] passing repetition header \n")
					continue
				}

				err := c.handleLog(&l, client)
				if err != nil {
					fmt.Printf("%v\n", err)
					break
				}
				// write cur to file
				err = lib.WriteIndexTo(int64(l.BlockNumber), c.indexFilePath)
				if err != nil {
					fmt.Printf("%v\n", err)
					break
				}
				fmt.Printf("[event][market] record index: %v \n", int64(l.BlockNumber))
			}
			// another loop
		}
	}
}

func (c *ContractEvent) handleLog(log *types.Log, client *ethclient.Client) error {
	if len(log.Topics) <= 0 {
		return nil
	}

	t := log.Topics[0].Hex()
	switch t {
	case c.runTime.Config.Contract.PurchasedEvent:
		// purchase event
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(log.BlockNumber)))
		if err != nil {
			return err
		}
		timestamp := block.Time()

		l, err := c.marketFilter.ParsePurchased(*log)
		if err != nil {
			return err
		}

		err = c.syncPurchased(l, log.BlockNumber, timestamp)
		if err != nil {
			return err
		}

		fmt.Printf("[event][market][purchased] block: %v, seller: %v, buyer: %v, offer: %v, amout: %v, total price: %v \n",
			log.BlockNumber, l.Seller, l.Buyer, l.OfferId, l.Amount, l.TotalPrice,
		)
	}

	return nil
}

func (c *ContractEvent) syncPurchased(event *market.MarketPurchased, block uint64, t uint64) error {
	amount, _ := event.Amount.Float64()

	err := c.runTime.Mysql.Transaction(func(tx *gorm.DB) error {
		purchased := models.Purchased{
			BlockID:   block,
			OfferID:   event.OfferId.Uint64(),
			Seller:    event.Seller.String(),
			Buyer:     event.Buyer.String(),
			Amount:    amount,
			Timestamp: t,
		}

		if e := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&purchased).Error; e != nil {
			return e
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
