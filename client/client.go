package main

import (
	"fmt"
	"github/ReigenDing/sales-system/common"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
)

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000")
	var veg common.Vegetable
	var vegs []string
	// var vegName string
	// var vegAvailableAmount float32
	// var vegPricePerKg float32

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "this is client sever")
	})

	mux.HandleFunc("/vegetables/add", func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprint(rw, "unable to parse form: ", err)
			log.Print("Unable to parse form: ", err)
			return
		}
		fmt.Printf("request data => %v\n", r.Body)
		fmt.Printf("request data => %v\n", r.Form)
		fmt.Printf("request data => %v\n", r.PostForm)
		fmt.Printf("request data => %v\n", r.RemoteAddr)
		pricePerKgValue, err := strconv.ParseFloat(r.FormValue("pricePerKg"), 32)
		if err != nil {
			fmt.Fprintln(rw, "unable to read input pricePerKg: Please enter valid value (float/int)")
			log.Print("unable to read input pricePerKg: ", err)
			return
		}
		pricePerkgFloat32 := float32(pricePerKgValue)

		availableAmountOfKgValue, err := strconv.ParseFloat(r.FormValue("amount"), 32)

		if err != nil {
			fmt.Fprint(rw, "unable to read the amount of vegetable: please enter valid amount value (float/int)")
			log.Print("unbale to read input amount of vegetable: ", err)
		}
		availableAmount := float32(availableAmountOfKgValue)
		if err := client.Call("Market.Add", common.Vegetable{
			Name:       r.FormValue("name"),
			PricePerKg: pricePerkgFloat32,
			Amount:     availableAmount,
		}, &veg); err != nil {
			fmt.Println(err)
			fmt.Fprint(rw, err)
			return
		}
		fmt.Printf("vegetable %s created\n", veg.Name)
		fmt.Fprintf(rw, "vegertable %s created\n", veg.Name)

	})

	mux.HandleFunc("/vegetables/update", func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("unable to parse form, err: %v", err)
			return
		}
		pricePerKgValue, err := strconv.ParseFloat(r.FormValue("pricePerKg"), 32)
		if err != nil {
			log.Printf("unable to read the input of pricePerKg, err: %v", err)
			return
		}
		pricePerkgFloat32 := float32(pricePerKgValue)

		availableAmountOfKgValue, err := strconv.ParseFloat(r.FormValue("amount"), 32)
		if err != nil {
			log.Printf("unable to read the input of amount, err: %v", err)
			return
		}
		availableAmountFloat32 := float32(availableAmountOfKgValue)

		if err := client.Call("Market.Update", common.Vegetable{
			Name:       r.FormValue("name"),
			PricePerKg: pricePerkgFloat32,
			Amount:     availableAmountFloat32,
		}, &veg); err != nil {
			fmt.Fprint(rw, err)
			return
		}
		fmt.Printf("vegetable %s updated\n", veg.Name)
		fmt.Fprintf(rw, "vegetable name %s update\n", veg.Name)
		fmt.Fprintf(rw, "vegetable price %f update\n", veg.PricePerKg)
		fmt.Fprintf(rw, "vegetable amount %f update\n", veg.Amount)

	})

	mux.HandleFunc("/vegetables/get", func(rw http.ResponseWriter, r *http.Request) {
		if err := client.Call("Market.Get", r.FormValue("name"), &veg); err != nil {
			fmt.Fprint(rw, err)
			return
		}
		fmt.Fprintf(rw, "vegtable %s found\n", veg.Name)
		fmt.Fprintf(rw, "vegtable %s price is %f \n", veg.Name, veg.PricePerKg)
		fmt.Fprintf(rw, "vegtable %s amount is %f \n", veg.Name, veg.Amount)
	})

	mux.HandleFunc("/vegetables/get/all", func(rw http.ResponseWriter, r *http.Request) {
		if err := client.Call("Market.GetAll", "", &vegs); err != nil {
			fmt.Fprint(rw, err)
			return
		}
		fmt.Printf("vegs => %v", vegs)
		fmt.Fprint(rw, vegs)
	})

	http.ListenAndServe(":9001", mux)
}
