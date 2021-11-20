package main

import (
	"github/ReigenDing/sales-system/common"
	"io"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	market := common.NewMarket()

	rpc.Register(market)

	rpc.HandleHTTP()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "RPC Server live!")
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
