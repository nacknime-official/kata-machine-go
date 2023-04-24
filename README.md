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

BTW there is `src/DSA/helpers.go` file with some common functions/structs (e.g. `ZeroValue`), you can use them if you want to.

## Testing
Just run `go test ./...` within a day folder if you want to run all the tests for a day.

Or `go test ./QuickSort` within a day folder if you want to run only *QuickSort* tests.

Check `go help test` for more details.

There is also `test` command you can use from the root folder:
```bash
go run main.go test             # runs all the tests for current day
go run main.go test -day 1      # runs all the tests for day 1
go run main.go test Queue       # runs Queue test for current day
go run main.go test Queue Stack # runs Queue and Stack tests for current day
go run main.go test Queue -v    # runs Queue test for current day with verbose output
go run main.go test -- -v       # runs all the tests for current day with verbose output
```

## TODO
- [ ] Add other DSA
    - [x] LinearSearchList
    - [x] BubbleSort
    - [x] SinglyLinkedList
    - [x] DoublyLinkedList
    - [x] Queue
    - [x] Stack
    - [x] ArrayList
    - [x] QuickSort
    - [x] MergeSort
    - [x] MinHeap
    - [ ] DFSOnBST
    - [ ] LRU
    - [ ] BinarySearchList
    - [ ] TwoCrystalBalls
    - [ ] MazeSolver
    - [ ] BTPreOrder
    - [ ] BTInOrder
    - [ ] BTPostOrder
    - [ ] BTBFS
    - [ ] CompareBinaryTrees
    - [ ] DFSOnBST
    - [ ] DFSGraphList
    - [ ] Trie
    - [ ] BFSGraphMatrix
    - [ ] Map
- [x] "test" command
- [ ] "test" command (pattern matching)
- [ ] Stats
