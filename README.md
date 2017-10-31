# cryptovalues
CLI-tool to get cryptocurrency info from the [CryptoCompare API](https://min-api.cryptocompare.com/).

## Usage

```
$./cryptovalues
Get latest cryptocurrency info from CryptoCompare

Usage:
  cryptovalues [command]

Available Commands:
  details     show details of a cryptocurrency
  help        Help about any command
  list        list all available cryptocurrencies
  values      Get current values

Flags:
      --config string   config file (default is $HOME/.cryptovalues.yaml)
  -h, --help            help for cryptovalues

Use "cryptovalues [command] --help" for more information about a command.
```

## list

List command returns the list of available cryptocurrencies in the API.

**Usage:**

```
$./cryptovalues list --help
Lists all the available cryptocurrency symbols.

Usage:
  cryptovalues list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --config string   config file (default is $HOME/.cryptovalues.yaml)
```



## details

Shows the information available in the API for a single currency.

**Usage:**

```
$./cryptovalues details --help
Show the details of a single cryptocurrency.

Usage:
  cryptovalues details [flags]

Flags:
  -h, --help            help for details
  -s, --symbol string   currency symbol (required)

Global Flags:
      --config string   config file (default is $HOME/.cryptovalues.yaml)
```

**Example:**

```
$./cryptovalues details -s ETH
Details for ETH
- Symbol:            ETH
- Coin name:         Ethereum
- Full name:         Ethereum  (ETH)
- Algorithm:         Ethash
- Proof type:        PoW
```



## values

**Usage:**

```
$./cryptovalues values --help
Get current values

Usage:
  cryptovalues values [flags]

Flags:
  -f, --from string   comma-separated source currency list. Example: ETH,BTC
  -h, --help          help for values
  -t, --to string     comma-separated target currency list. Example: EUR,USD

Global Flags:
      --config string   config file (default is $HOME/.cryptovalues.yaml)
```



**Example:**

```
$./cryptovalues values -f BTC -t "USD, EUR, GBP"
Getting currency values for BTC in USD, EUR, GBP from CryptoCompare

2017-10-31 13:46:28.7035289 +0100 CET m=+0.443470500

1.0 BTC -> 4714.31 GBP
1.0 BTC -> 6208.79 USD
1.0 BTC -> 5354.35 EUR

```

## Building

You need to inject version and compiled at link time:

```
go build -ldflags "-X github.com/riussi/cryptovalues/cmd.compiled=20171031-141614 -X github.com/riussi/cryptovalues/cmd.version=1.0.0"
```

## License

[Apache 2.0](LICENSE)