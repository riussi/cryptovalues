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

package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	api "gitlab.com/juha.ristolainen/cryptovalues/api"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all available cryptocurrencies",
	Long:  `Lists all the available cryptocurrency symbols.`,
	Run: func(cmd *cobra.Command, args []string) {
		coinlist := api.GetCoinlist()
		printOutCoinlist(&coinlist)
	},
}

func printOutCoinlist(coinlist *api.CoinList) {
	keys := make([]string, 0, len(coinlist.Data))
	for key := range coinlist.Data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Printf("There are %d currencies available in the API\n", len(keys))
	fmt.Printf("%-10s- %-20s - %-10s- %-10s\n", "Symbol", "Coin name", "Prooftype", "Algorithm")
	for _, value := range keys {
		fmt.Printf("%-10s- %-20s - %-10s- %-10s\n", coinlist.Data[value].Symbol, coinlist.Data[value].CoinName, coinlist.Data[value].ProofType, coinlist.Data[value].Algorithm)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
