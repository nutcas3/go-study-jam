# Concurrent Prime Number Finder

This project is part of Week 4 of the Go course. It finds prime numbers concurrently using goroutines and channels.

## How to run

Run the prime finder from the `cmd` directory with a start and end range:

```bash
# Syntax:
go run cmd/main.go <start> <end>

# Example:
go run cmd/main.go 10 50
# Finds primes between 10 and 50
```

## Tasks
- Accept a range of numbers from the user.
- Use goroutines to check for primes concurrently.
- Output the list of prime numbers found.
