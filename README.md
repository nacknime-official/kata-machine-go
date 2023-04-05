# Kata Machine (Go) 
[ThePrimeagen's kata machine](https://github.com/ThePrimeagen/kata-machine) for **CHAD** gophers.

## How it works
Make sure you have [Go](https://go.dev/doc/install) installed.

clone the repo and install the dependencies

```bash
go get ./...
```

edit the `config.yaml` file
```yaml
DSA: [
    QuickSort,
]
```

create a day of katas, this will use the list in the `config.yaml`.
```bash
go generate
# OR go run main.go generate
```
OR if you want to generate katas not from config but from arguments instead:
```bash
go run main.go generate QuickSort Queue
```
OR if you want to pass custom config:
```bash
go run main.go generate --configPath some_other_config.yaml
```

this will progressively create folders named

```
src/day1
src/day2
...
```

that will contain files with empty kata implementation and files with tests for it.

## Testing
Just run `go test ./...` within a day folder if you want to run all the tests for a day.

Or `go test ./QuickSort` within a day folder if you want to run only *QuickSort* tests.

Check `go help test` for more details.

Soon I will make `go run main.go test` command soon that will automatically runs all the tests for the current day or that you have specified in the arguments.

## TODO
- [ ] Add other DSA
- [ ] "test" command
- [ ] Stats
