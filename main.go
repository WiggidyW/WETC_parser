package main

import (
	"github.com/evepraisal/go-evepraisal/parsers"
	"encoding/json"
	"fmt"
	"os"
	"io"
)

type OutItem struct {
	Name	string	`json:"name"`
	Quantity	int64	`json:"quantity"`
}

func main() {
	itemsMap := make(map[string]int64)
	var items = make([]OutItem, 0)

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	result, _ := parsers.AllParser(parsers.StringToInput(string(input)))

	for _, sub_result := range result.(*parsers.MultiParserResult).Results {
		switch r := sub_result.(type) {
			default:
				fmt.Println("Unreachable")
			case *parsers.AssetList:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.CargoScan:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.Contract:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.EFT:
				itemsMap[r.Ship] += 1
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.Fitting:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.Industry:
				for _, item := range r.Items {
					if item.Quantity == 1 && item.BPCRuns > 1 {
						itemsMap[item.Name] += item.BPCRuns
					}
					else {
						itemsMap[item.Name] += item.Quantity
					}
				}
			case *parsers.Listing:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.LootHistory:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.PI:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.SurveyScan:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.ViewContents:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.MiningLedger:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.MoonLedger:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.HeuristicResult:
				for _, item := range r.Items {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.DScan:
				for _, item := range r.Items {
					itemsMap[item.Name] += 1
				}
			case *parsers.Compare:
				for _, item := range r.Items {
					itemsMap[item.Name] += 1
				}
			case *parsers.Wallet:
				for _, item := range r.ItemizedTransactions {
					itemsMap[item.Name] += item.Quantity
				}
			case *parsers.Killmail:
				for _, item := range r.Dropped {
					itemsMap[item.Name] += item.Quantity
				}
				for _, item := range r.Destroyed {
					itemsMap[item.Name] += item.Quantity
				}
		}
	}

	for name, quantity := range itemsMap {
		if quantity > 0 {
			items = append(items, OutItem{Name: name, Quantity: quantity})
		}
	}

	if err := json.NewEncoder(os.Stdout).Encode(items); err != nil {
		panic(err)
	}
}
