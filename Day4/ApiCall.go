package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type StockResponseData struct {
	Data StockDetails
}

type StockDetailsData struct {
	PriceCurrent   string
	PricePrevClose string
}

func responseHandlerApi(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("stock") == "SBI" {
		resp, err := http.Get("https://priceapi.moneycontrol.com/pricefeed/nse/equitycash/SBI")
		defer resp.Body.Close()
		var stockResponseData = new(StockResponse)
		if err == nil {
			dataBytes, readErr := io.ReadAll(resp.Body)

			if readErr == nil {
				json.Unmarshal(dataBytes, stockResponseData)
				fmt.Fprintf(response, "The Value for the Stock SBI is %s", stockResponseData.Data.PriceCurrent)
			} else {
				fmt.Fprintf(response, "Error Reading from the Apis")
			}
		} else {
			fmt.Fprintf(response, "Invalid Url")
		}
	}
}

func apiCall() {
	http.HandleFunc("/stockprice", responseHandler)
	http.ListenAndServe(":8000", nil)
}
