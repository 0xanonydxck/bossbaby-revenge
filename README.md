# Problem: Boss Baby's Revenge

## Setup & Execute
1. Download project used dependencies.
```bash
go mod download
```

2. Run `Go` command.
```bash
go run main.go
```
or (optional)
```bash
make run
```

## Test & Benchmark
1. **Time Complexity**: `O(n)`
2. **Memory Complexity**: `O(n)`

### Test Result
```bash
=== RUN   TestBossBabyRevenge
=== RUN   TestBossBabyRevenge/[SRSSRRR]GoodBoy
=== RUN   TestBossBabyRevenge/[SSRSRRR]GoodBoy
=== RUN   TestBossBabyRevenge/[RSSRR]GoodBoy
=== RUN   TestBossBabyRevenge/[SSSRRRRS]BadBoy
=== RUN   TestBossBabyRevenge/[SRRSSR]BadBoy
=== RUN   TestBossBabyRevenge/[SRRSSR]BadBoy#01
=== RUN   TestBossBabyRevenge/[]Panic
=== RUN   TestBossBabyRevenge/[ABC]Panic
--- PASS: TestBossBabyRevenge (0.00s)
    --- PASS: TestBossBabyRevenge/[SRSSRRR]GoodBoy (0.00s)
    --- PASS: TestBossBabyRevenge/[SSRSRRR]GoodBoy (0.00s)
    --- PASS: TestBossBabyRevenge/[RSSRR]GoodBoy (0.00s)
    --- PASS: TestBossBabyRevenge/[SSSRRRRS]BadBoy (0.00s)
    --- PASS: TestBossBabyRevenge/[SRRSSR]BadBoy (0.00s)
    --- PASS: TestBossBabyRevenge/[SRRSSR]BadBoy#01 (0.00s)
    --- PASS: TestBossBabyRevenge/[]Panic (0.00s)
    --- PASS: TestBossBabyRevenge/[ABC]Panic (0.00s)
PASS
coverage: 0.0% of statements
ok      command-line-arguments  0.381s  coverage: 0.0% of statements
```

### Benchmark Result
```bash
goos: darwin
goarch: arm64
pkg: bossbaby/module
BenchmarkBossBabyRevenge-14       198735              6219 ns/op
BenchmarkBossBabyRevenge-14       185779              6180 ns/op
BenchmarkBossBabyRevenge-14       184828              6113 ns/op
BenchmarkBossBabyRevenge-14       190645              6116 ns/op
BenchmarkBossBabyRevenge-14       189220              6125 ns/op
PASS
ok      bossbaby/module 7.495s
```