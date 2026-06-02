# leetcode
leetcode problem solving




# Benchmarking 
```bash
# Run all benchmarks in the current directory (CPU execution metrics only)
go test -bench=.

# Run benchmarks and expose runtime memory tracking (B/op and allocs/op columns)
go test -bench=. -benchmem

# Filter execution to run ONLY benchmarks matching a specific string or regex pattern
go test -bench=ThreeSum -benchmem

# Skip standard TestXxx unit tests to jump straight into benchmark evaluations
go test -run=^$ -bench=. -benchmem

# Increase execution duration to gather a more stable, prolonged statistical sample
go test -bench=. -benchtime=5s -benchmem

# Force the loop to execute an exact block count instead of calculating across a time window
go test -bench=. -benchtime=5000x -benchmem

# Repeat the entire benchmarking sequence multiple times to monitor performance variance
go test -bench=. -count=5 -benchmem

# 
go test -run=^$ -bench=. -cpuprofile=cpu.pprof -memprofile=mem.pprof
```