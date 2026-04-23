# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`cryptovalues` is a Go CLI that queries the public CryptoCompare API
(`https://min-api.cryptocompare.com`) for cryptocurrency metadata and prices.
Module path: `github.com/riussi/cryptovalues` (Go 1.24).

## Commands

```bash
# Build. The ldflags injection is required so `cryptovalues version` prints a
# real value instead of the zero-initialized vars in main.go.
go build -ldflags "-X github.com/riussi/cryptovalues/cmd.version=1.0.2-`date -u +%Y%m%d.%H%M%S`"

# Run without building
go run . <subcommand> [flags]      # e.g. go run . values -f ETH -t EUR,USD

go mod tidy                         # sync go.mod / go.sum
go vet ./...                        # static checks

# Cross-platform release archives (darwin/linux/windows amd64) — see .goreleaser.yml
goreleaser release --snapshot --clean
```

There is no test suite in this repo — no `*_test.go` files exist. Do not claim
tests were run; if you add tests, put them alongside the package they cover and
run `go test ./...`.

## Architecture

Three-layer layout; each layer depends only on the one below it.

- `main.go` — thin entrypoint. Holds four string vars (`version`, `commit`,
  `date`, `builtBy`) that goreleaser / `-ldflags` overwrite at link time, then
  copies them into the `cmd` package before calling `cmd.Execute()`.
- `cmd/` — Cobra subcommands. One file per command (`list`, `details`,
  `values`, `version`), plus `root.go` which owns the persistent `--config`
  flag and wires Viper up to `$HOME/.cryptovalues.yaml` + env vars.
- `api/coinlist.go` — the entire HTTP / JSON layer. Exposes `GetCoinlist()`,
  `GetCurrencyDetails(symbol)`, and `GetCurrencyValues(from, to, amount)` and
  defines the `CoinList` / `Datum` response structs. `APIBaseURL` is the only
  config knob.

### Conventions worth preserving

- **Command registration pattern:** every file in `cmd/` declares its own
  `var fooCmd = &cobra.Command{...}` and registers itself via
  `RootCmd.AddCommand(fooCmd)` inside that file's `init()`. When adding a
  subcommand, follow this pattern — do not touch `root.go`.
- **Flag ↔ config binding:** flags are bound to Viper keys namespaced by
  command (`values.from`, `details.symbol`, …) via `viper.BindPFlag`, and
  command bodies read through `viper.GetString(...)` rather than the Cobra
  flag directly. This is what lets `~/.cryptovalues.yaml` override defaults.
- **Error handling is intentionally fatal:** network / JSON failures in `api/`
  print to stdout and call `os.Exit(1)`. Matches the CLI style of the rest of
  the codebase; don't refactor to returned errors unless the task requires it.
- **Fiat symbols are hard-coded** in `cmd/values.go` as `AcceptedFiatCurrencies`.
  The TODO there is intentional — CryptoCompare doesn't expose a fiat list, so
  new symbols are added manually.
- **Version injection target:** ldflags must set
  `github.com/riussi/cryptovalues/cmd.version` (lowercase), which `main.go`
  then promotes to the exported `cmd.Version`. The other three fields
  (`commit`, `date`, `builtBy`) follow the same pattern and are populated
  automatically by goreleaser.
