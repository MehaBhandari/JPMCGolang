package main

import (
	"net/http"
	"text/template"
)

type StockResponse struct {
	Data StockDetails
}

type StockDetails struct {
	PriceCurrent   string
	PricePrevClose string
}

func responseHandler(response http.ResponseWriter, request *http.Request) {

	MyStockData := StockDetails{
		PricePrevClose: "500",
		PriceCurrent:   "540",
	}
	stockTemplate := template.Must(template.New("stockTemplate").Parse(`
		<h1>The Stock Value for SBI<h1>
		<h2>The Current Price of Stock is {{.PriceCurrent}}
		<h2>The Previous Price of Stock is {{.PricePrevClose}}
	`))

	stockTemplate.Execute(response, MyStockData)
}

func simpleServer() {
	http.HandleFunc("/stockprice", responseHandler)
	http.ListenAndServe(":8000", nil)
}
