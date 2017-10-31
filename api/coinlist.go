// Copyright © 2017 Juha Ristolainen <juha.ristolainen@iki.fi>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs"
)

const ApiBaseUrl = "https://min-api.cryptocompare.com"

func GetCoinlist() CoinList {
	// Construct the API URL
	var url = ApiBaseUrl + "/data/all/coinlist?extraParams=cryptovalues"

	// Get the API response#
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: connection to %s failed: %s\n\n", url, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Parse the JSON
	list, err := UnmarshalCoinList([]byte(body))
	if err != nil {
		fmt.Printf("ERROR: failed to parse JSON: %s\n\n", err.Error())
		os.Exit(1)
	}
	return list
}

func GetCurrencyDetails(symbol string) Datum {
	coinlist := GetCoinlist()

	var data Datum
	for key := range coinlist.Data {
		if key == symbol {
			data = coinlist.Data[key]
			break
		}
	}
	return data
}

func GetCurrencyValues(fromCurrencies *[]string, toCurrencies *[]string) {
	// Construct the API URL
	var url = ApiBaseUrl + "/data/pricemulti?fsyms=" + strings.Join(*fromCurrencies, ",") + "&tsyms=" + strings.Join(*toCurrencies, ",") + "&extraParams=cryptovalues"

	// Get the API response#
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: connection to %s failed: %s\n\n", url, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Parse the JSON
	jsonParsed, err := gabs.ParseJSON([]byte(body))
	if err != nil {
		fmt.Printf("ERROR: failed to parse JSON: %s\n\n", err.Error())
		os.Exit(1)
	}

	// Print out the results
	for _, fromCurrency := range *fromCurrencies {
		children, _ := jsonParsed.S(fromCurrency).ChildrenMap()
		for key, child := range children {
			fmt.Printf("1.0 %s -> %v %s\n", fromCurrency, child, key)
		}
		fmt.Println()
	}
}

func UnmarshalCoinList(data []byte) (CoinList, error) {
	var r CoinList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CoinList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CoinList struct {
	BaseImageUrl     string           `json:"BaseImageUrl"`
	BaseLinkUrl      string           `json:"BaseLinkUrl"`
	Data             map[string]Datum `json:"Data"`
	DefaultWatchlist DefaultWatchlist `json:"DefaultWatchlist"`
	Message          string           `json:"Message"`
	Response         string           `json:"Response"`
	Type             int64            `json:"Type"`
}

type DefaultWatchlist struct {
	CoinIs    string `json:"CoinIs"`
	Sponsored string `json:"Sponsored"`
}

type Datum struct {
	Algorithm           string  `json:"Algorithm"`
	CoinName            string  `json:"CoinName"`
	FullName            string  `json:"FullName"`
	FullyPremined       string  `json:"FullyPremined"`
	Id                  string  `json:"Id"`
	ImageUrl            *string `json:"ImageUrl"`
	Name                string  `json:"Name"`
	PreMinedValue       string  `json:"PreMinedValue"`
	ProofType           string  `json:"ProofType"`
	SortOrder           string  `json:"SortOrder"`
	Sponsored           bool    `json:"Sponsored"`
	Symbol              string  `json:"Symbol"`
	TotalCoinSupply     string  `json:"TotalCoinSupply"`
	TotalCoinsFreeFloat string  `json:"TotalCoinsFreeFloat"`
	Url                 string  `json:"Url"`
}
