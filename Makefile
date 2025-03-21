.PHONY: abis

default: abis

abis:
	cd ./abis/market && abigen --abi=market.abi.json --pkg=market --out=market.go

linux:
	GOOS=linux GOARCH=amd64 go build -o sync
