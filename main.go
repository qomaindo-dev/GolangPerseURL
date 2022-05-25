package main

import (
	"fmt"
	"net/url"
	"reflect"
)

func main() {
	address := "https://www.toko-kita.com/dataproduk?databrand_id=181808782&variations%5B%5D=302008457&variations%5B%5D=302008463"

	u, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	fmt.Println("query: ", u.RawQuery)

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}

	brandID := q.Get("databrand_id")
	if brandID == "" {
		fmt.Println("Kosong")
	} else {
		fmt.Println("brandID: ", brandID)
	}

	variations := q.Get("&variations[]")
	if variations == "" {
		fmt.Println("Kosong variations")
	} else {
		fmt.Println("variations", variations)
	}

	getQuery := u.Query()
	dataApa := getQuery["sort"]
	fmt.Println("dataApa: ", reflect.TypeOf(dataApa), ": ", dataApa)
	fmt.Println("getQuery: ", getQuery)
	fmt.Println(reflect.TypeOf(getQuery))
	fmt.Println(getQuery.Get("sort"))

}
