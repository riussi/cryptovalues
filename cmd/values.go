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
	"strings"

	"time"

	"github.com/riussi/cryptovalues/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO add all supported fiat currency symbols
var AcceptedFiatCurrencies = []string{
	"EUR", "USD", "GBP",
}

var valuesCmd = &cobra.Command{
	Use:   "values",
	Short: "Get current values",
	Run: func(cmd *cobra.Command, args []string) {
		from := viper.GetString("values.from")
		to := viper.GetString("values.to")

		if len(from) == 0 || len(to) == 0 {
			cmd.Usage()
			return
		}
		from = strings.Replace(from, " ", "", -1)
		to = strings.Replace(to, " ", "", -1)
		fromCurrencies := strings.Split(from, ",")
		toCurrencies := strings.Split(to, ",")

		coinlist := api.GetCoinlist()

		keys := make([]string, len(coinlist.Data))

		i := 0
		for k := range coinlist.Data {
			keys[i] = k
			i++
		}
		keys = append(keys, AcceptedFiatCurrencies...)

		isValid1 := validateCurrencies(&fromCurrencies, &keys)
		isValid2 := validateCurrencies(&toCurrencies, &keys)
		if !isValid1 || !isValid2 {
			return
		}

		start := time.Now()
		fmt.Printf("Getting currency values for %s in %s from CryptoCompare\n\n%s\n\n", strings.Join(fromCurrencies, ", "), strings.Join(toCurrencies, ", "), start)

		// Get the actual values from the API
		api.GetCurrencyValues(&fromCurrencies, &toCurrencies)
	},
}

func init() {
	RootCmd.AddCommand(valuesCmd)

	// Setup the from and to flags and also bind to config-file values with viper.
	valuesCmd.Flags().StringP("from", "f", "", "comma-separated source currency list. Example: ETH,BTC")
	valuesCmd.Flags().StringP("to", "t", "", "comma-separated target currency list. Example: EUR,USD")
	viper.BindPFlag("values.from", valuesCmd.Flags().Lookup("from"))
	viper.BindPFlag("values.to", valuesCmd.Flags().Lookup("to"))
}

// Check that all given currency symbols are valid
func validateCurrencies(currencies *[]string, acceptedList *[]string) bool {
	var isValid bool = false
	for _, value := range *currencies {
		isValid = validateCurrency(value, acceptedList)
		if !isValid {
			fmt.Printf("%s is not a valid symbol.\n", value)
		}
	}
	return isValid
}

func validateCurrency(symbol string, acceptedList *[]string) bool {
	var isValid bool = false
	for _, value := range *acceptedList {
		//fmt.Printf("Checking %s against %s\n", symbol, value)
		if value == symbol {
			isValid = true
			break
		}
	}
	return isValid
}
