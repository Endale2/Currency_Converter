package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type ExchangeResponse struct {
    Amount float64            `json:"amount"`
    Base   string             `json:"base"`
    Date   string             `json:"date"`
    Rates  map[string]float64 `json:"rates"`
}

func main() {
    var base, target string
    var amount float64

    fmt.Print("Enter base currency (e.g., USD): ")
    fmt.Scanln(&base)

    fmt.Print("Enter target currency (e.g., EUR): ")
    fmt.Scanln(&target)

    fmt.Print("Enter amount to convert: ")
    fmt.Scanln(&amount)

    url := fmt.Sprintf("https://api.frankfurter.app/latest?amount=%f&from=%s&to=%s", amount, base, target)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal("Error fetching API:", err)
    }
    defer resp.Body.Close()

    var data ExchangeResponse
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        log.Fatal("Error decoding JSON:", err)
    }

    converted := data.Rates[target]
	fmt.Println("========================================================")
    fmt.Printf("%.2f %s = %.2f %s\n", amount, base, converted, target)
	fmt.Println("========================================================")
}
