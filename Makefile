run:
	go run cmd/main.go
test:
	go test ./...  
bench:
	go test -bench=.  ./...  