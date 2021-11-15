package common

import (
	"os"
	"sync"
	"testing"
)

func TestVegetableVtruct(t *testing.T) {
	vg := Vegetable{
		Name:       "vegetable1",
		PricePerKg: 1.5,
		Amount:     10.0,
	}
	if vg.Amount != 10.0 {
		t.Errorf("got '%f' want '%f'", vg.Amount, 10.0)
	}
	if vg.Name != "vegetable1" {
		t.Errorf("got %s want %s", vg.Name, "vegetable1")
	}
	if vg.PricePerKg != 1.5 {
		t.Errorf("got %f want %f", vg.PricePerKg, 1.5)
	}

}

func TestWriteCsvFile(t *testing.T) {
	var m sync.RWMutex
	vegs := []*Vegetable{
		{
			Name:       "veg1",
			PricePerKg: 1.2,
			Amount:     19.0,
		},
		{
			Name:       "veg2",
			PricePerKg: 2.4,
			Amount:     90.2,
		},
		{
			Name:       "veg3",
			PricePerKg: 5.2,
			Amount:     9.8,
		},
	}
	got := writeCsvFile(vegs, &m)
	if _, err := os.Stat("data.csv"); err == os.ErrExist {
		t.Fatalf("file create faild got %v", got)
	}

}

func TestReadCsvFile(t *testing.T) {
	vegs := []*Vegetable{}
	var m sync.RWMutex
	filePath := "data.csv"

	t.Run("read csv file", func(t *testing.T) {
		vegs = readCsvFile(filePath, &m)
		if len(vegs) != 3 {
			t.Errorf("want %d got %d", 3, len(vegs))
		}
	})
}
