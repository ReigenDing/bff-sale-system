package common

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gocarina/gocsv"
)

type Vegetable struct {
	Name       string  `csv:"Name"`
	PricePerKg float32 `csv:"PricePerKg"`
	Amount     float32 `csv:"Amount"`
}

type Market struct {
	database []*Vegetable
}

func writeCsvFile(newVegs []*Vegetable, m *sync.RWMutex) []*Vegetable {
	m.Lock()

	defer m.Unlock()

	err := os.Remove("data.csv")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	f, err := os.OpenFile("data.csv", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Write([]string{"Name", "PricePerKg", "Amount"})
	for _, veg := range newVegs {
		row := []string{veg.Name, fmt.Sprintf("%.2f", veg.PricePerKg), fmt.Sprintf("%.2f", veg.Amount)}
		writer.Write(row)
	}
	writer.Flush()
	return newVegs
}

func readCsvFile(filePath string, m *sync.RWMutex) (vegetables []*Vegetable) {
	m.RLock()
	defer m.RUnlock()

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &vegetables); err != nil {
		log.Fatal(err)
		return nil
	}
	return
}

func NewMarket() *Market {
	var m sync.RWMutex
	return &Market{
		database: readCsvFile("data.csv", &m),
	}
}

func (market *Market) Get(playload string, reply *Vegetable) error {
	for _, v := range market.database {
		if v.Name == playload {
			*reply = *v
			return nil
		}
	}
	return fmt.Errorf("vegetables %s is not exists", playload)
}

func (market *Market) GetAmount(playload string, reply *float32) error {
	for _, v := range market.database {
		if v.Name == playload {
			*reply = v.Amount
			return nil
		}
	}
	return fmt.Errorf("vegetables %s is not exists", playload)
}

func (market *Market) GetPricePerKg(name string, reply *float32) error {
	for _, v := range market.database {
		if v.Name == name {
			*reply = v.PricePerKg
			return nil
		}
	}
	return fmt.Errorf("vegetables %s is not exists", name)
}
