package common

import (
	"fmt"
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

func TestGet(t *testing.T) {
	market := NewMarket()
	var reply Vegetable
	t.Run("get vegetable by name", func(t *testing.T) {
		test_case := []struct {
			name string
			want string
		}{
			{"veg1", "veg1"},
			{"veg2", "veg2"},
			{"veg3", "veg3"},
			{"veg4", "vegetables veg4 is not exists"},
		}
		for _, v := range test_case {
			err := market.Get(v.name, &reply)
			if v.name == "veg4" && err.Error() != v.want {
				t.Errorf("want '%s' but got '%s'", err, fmt.Errorf(v.want))
			}
			if v.name != "veg4" && v.name != reply.Name {
				t.Errorf("want %s but got %s", v.name, reply.Name)
			}
		}
	})
}

func TestGetAmount(t *testing.T) {
	market := NewMarket()
	var reply float32
	t.Run("get vegetable amount by name", func(t *testing.T) {
		test_case := []struct {
			name string
			want float32
		}{
			{"veg1", 19.0},
			{"veg2", 90.2},
			{"veg3", 9.8},
			{"veg4", 0},
		}
		for _, v := range test_case {
			err := market.GetAmount(v.name, &reply)
			if err != nil && err.Error() != "vegetables veg4 is not exists" {
				t.Errorf("want '%s' but got '%s'", "vegetables veg4 is not exists", err.Error())
			} else if err == nil && v.want != reply {
				t.Errorf("want '%f' but got '%f'", v.want, reply)
			}

		}
	})
}

func TestGetPricePerKg(t *testing.T) {
	market := NewMarket()
	var reply float32
	t.Run("get vegetable price per kg by name", func(t *testing.T) {
		test_case := []struct {
			name string
			want float32
		}{
			{"veg1", 1.2},
			{"veg2", 2.4},
			{"veg3", 5.2},
			{"veg4", 0},
		}
		for _, v := range test_case {
			err := market.GetPricePerKg(v.name, &reply)
			if err != nil && err.Error() != "vegetables veg4 is not exists" {
				t.Errorf("want '%s' but got '%s'", "vegetables veg4 is not exists", err.Error())

			} else if err == nil && v.want != reply {
				t.Errorf("want '%f' but got '%f'", v.want, reply)

			}
		}
	})
}
