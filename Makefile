.PHONY: run test bench profile-cpu clean

# Default directory target to run all if 'dir' is not provided
dir ?= ...

# Run the application logic directly. Usage: make run dir=1-two-sum
run:
	@if [ "$(dir)" = "..." ]; then echo "Error: Please specify a directory, e.g., make run dir=1-two-sum"; exit 1; fi
	go run ./problems/$(dir)/main.go

# Run tests. Usage: 'make test' (all) OR 'make test dir=1-two-sum'
test:
ifeq ($(dir),...)
	go test -v ./problems/...
else
	go test -v ./problems/$(dir)/...
endif

# Run benchmarks. Usage: 'make bench' (all) OR 'make bench dir=1-two-sum'
bench:
ifeq ($(dir),...)
	go test -bench=. -benchmem ./problems/... | tee -a benchmark_results.txt
else
	go test -bench=. -benchmem ./problems/$(dir) | tee -a benchmark_results.txt
endif

# Profile CPU. Usage: 'make profile-cpu dir=1-two-sum'
profile-cpu:
	@if [ "$(dir)" = "..." ]; then echo "Error: Please specify a directory name under problems/, e.g., make profile-cpu dir=1-two-sum"; exit 1; fi
	go test -bench=. -cpuprofile=problems/$(dir)/cpu.pprof ./problems/$(dir)
	go tool pprof -pdf ./problems/$(dir)/$(dir).test problems/$(dir)/cpu.pprof > ./problems/$(dir)/cpu_profile.pdf

clean:
	find . -name "*.pprof" -type f -delete
	find . -name "*.test" -type f -delete