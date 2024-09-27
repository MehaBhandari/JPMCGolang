package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type StockResponseData struct {
	Data StockDetailsData
}

type StockDetailsData struct {
	PriceCurrent   string
	PricePrevClose string
}

type Counter struct {
	userCount  int
	stockValue string
}

func getInitialApplicationData() Counter {
	return Counter{
		userCount:  10,
		stockValue: "0",
	}
}

func (c Counter) Init() tea.Cmd {
	return tick()
}

func quitApplication(c Counter) (tea.Model, tea.Cmd) {
	return c, tea.Quit
}

func getUpdatedStockInfo(c Counter) (tea.Model, tea.Cmd) {
	resp, _ := http.Get("https://priceapi.moneycontrol.com/pricefeed/nse/equitycash/SBI")
	defer resp.Body.Close()
	var stockResponseData = new(StockResponseData)
	dataBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(dataBytes, stockResponseData)
	c.stockValue = stockResponseData.Data.PriceCurrent
	return c, nil
}

func (c Counter) Update(inputMsg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := inputMsg.(type) {
	case time.Time:
		c.userCount++
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			quitApplication(c)
		case "+":
			c.userCount++
		case "-":
			c.userCount--
		case "p":
			getUpdatedStockInfo(c)
		}

	}
	return c, nil
}

func tick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return time.Time(t)
	})
}

func (c Counter) View() string {
	return fmt.Sprintf("The Count Value is: %d ", c.userCount)
}

func main() {
	myApplication := tea.NewProgram(getInitialApplicationData())
	myApplication.Run()
}
