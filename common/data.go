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
	m        *sync.RWMutex
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
		fmt.Println("===================")
		log.Fatal("000000000000", err)
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &vegetables); err != nil {
		fmt.Println("===================")
		fmt.Println(err)
		// return nil
	}
	return
}

func NewMarket() *Market {
	var m sync.RWMutex
	return &Market{
		database: readCsvFile("data.csv", &m),
		m:        &m,
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

func (market *Market) GetAll(playload string, reply *[]string) error {
	for _, v := range market.database {
		*reply = append(*reply, v.Name)

	}
	return nil
}

func (market *Market) Add(playload Vegetable, reply *Vegetable) error {
	if ok, _ := vegetableAreadyExists(playload.Name, market.database); ok {
		return fmt.Errorf("vegetable with name %s already exists", playload.Name)
	}
	market.database = append(market.database, &playload)
	newVegetableToCsv(&playload)
	*reply = playload
	return nil
}

func vegetableAreadyExists(vegName string, array []*Vegetable) (bool, *Vegetable) {
	for _, veg := range array {
		if veg.Name == vegName {
			return true, veg
		}
	}
	return false, nil
}

func newVegetableToCsv(vegetable *Vegetable) {
	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("unable to read input file", err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()
	fmt.Printf("vegtable db => %v", records)
	writer := csv.NewWriter(f)
	if len(records) == 0 {
		writer.Write([]string{"Name", "PricePerKg", "AvaiableAmountofkg"})

	}
	var row []string
	row = append(row, vegetable.Name)
	row = append(row, fmt.Sprintf("%f", vegetable.PricePerKg))
	row = append(row, fmt.Sprintf("%f", vegetable.Amount))
	writer.Write(row)
	writer.Flush()

}

func (market *Market) Update(playload Vegetable, reply *Vegetable) error {

	exists, currentVeg := vegetableAreadyExists(playload.Name, market.database)
	if !exists {
		reply = nil
		return fmt.Errorf("vegetable with name %s is not exists", playload.Name)
	} else {
		currentVeg.Name = playload.Name
		currentVeg.Amount = playload.Amount
		currentVeg.PricePerKg = playload.PricePerKg
		res := writeCsvFile(market.database, market.m)
		fmt.Printf("new veg database %v\n", res)
		*reply = playload
	}
	return nil

}
