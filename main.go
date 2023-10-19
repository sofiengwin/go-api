package main

import (
	"fmt"

	"github.com/sofiengwin/cryptomasters/datatypes/api"
)

func main() {
	rate, _ := api.GetRate("BTC")
	r, _ := api.GetRate("ETH")
	fmt.Printf("%v %v \n", rate.Price, rate.Currency)
	fmt.Printf("%v  %v \n", r.Price, r.Currency)
}
