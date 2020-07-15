package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
    "fmt"
	"github.com/gorilla/mux"
)


type Stock struct {
	Ticker       string `json:"ticker,omitempty"`
	Co           string `json:"co,omitempty"`
	Endprice     string `json:"endprice,omitempty"`
	Openprice    string `json:"openprice,omitempty"`
	Currentprice string `json:"currentprice,omitempty"`
	Varreais     string `json:"varreais,omitempty"`
	Varcent      string `json:"varcent,omitempty"`
}

var stocks []Stock

func ReadCsv() ([][]string, error) {

    f, err := os.Open("input.csv")
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    s, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return s, nil
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stocks)
}


func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range stocks {
		if item.Ticker == params["ticker"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Stock{})
}

func main() {
	router := mux.NewRouter()

	s, err := ReadCsv()
    if err != nil {
        panic(err)
    }

    for _, e := range s {
        data := Stock{
            Ticker: e[0],
			Co: e[1],
			Endprice: e[2],
			Openprice: e[3],
			Currentprice: e[4],
			Varreais: e[5],
			Varcent: e[6],			
        }

        fmt.Println(data.Ticker + " " + data.Co + " " + data.Endprice + " " + data.Currentprice + " " + data.Varreais + " " + data.Varcent)

        stocks = append(stocks, data)
    }



	router.HandleFunc("/stocks", GetStocks).Methods("GET")
	router.HandleFunc("/stocks/{ticker}", GetStock).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
