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
	api "github.com/riussi/cryptovalues/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
	"os"
)

// listCmd represents the list command
var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "show details of a cryptocurrency",
	Long:  `Show the details of a single cryptocurrency.`,
	Run: func(cmd *cobra.Command, args []string) {
		symbol := viper.GetString("details.symbol")
		details := api.GetCurrencyDetails(symbol)
		if details.Symbol == "" {
			fmt.Printf("No details available for %s\n\n", symbol)
			os.Exit(1)
		}
		fmt.Printf("\nDetails for %s\n", symbol)
		printOutDetails(&details)
	},
}

func printOutDetails(data *api.Datum) {
	fmt.Printf("- Symbol:            %s\n", data.Symbol)
	fmt.Printf("- Coin name:         %s\n", data.CoinName)
	fmt.Printf("- Full name:         %s\n", data.FullName)
	fmt.Printf("- Algorithm:         %s\n", data.Algorithm)
	fmt.Printf("- Proof type:        %s\n", data.ProofType)
}

func init() {
	RootCmd.AddCommand(detailsCmd)
	detailsCmd.Flags().StringP("symbol", "s", "", "currency symbol (required)")
	detailsCmd.MarkFlagRequired("symbol")
	viper.BindPFlag("details.symbol", detailsCmd.Flags().Lookup("symbol"))
}
