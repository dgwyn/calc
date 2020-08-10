verify:
	make run
	make fmt
run:
	go run main.go
fmt:
	go fmt
test:
	make fmt
	go test
