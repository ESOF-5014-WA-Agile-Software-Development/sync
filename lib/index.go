package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Index struct {
	Cur int64 `json:"cur"`
}

func WriteIndexTo(i int64, f string) error {
	fp, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer func() {
		_ = fp.Close()
	}()
	if err != nil {
		return err
	}

	index := Index{Cur: i}
	data, err := json.Marshal(index)
	if err != nil {
		return err
	}

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadIndexFrom(f string) (int64, error) {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return 0, err
	}

	var index Index
	err = json.Unmarshal(content, &index)
	if err != nil {
		return 0, err
	}

	return index.Cur, nil
}
