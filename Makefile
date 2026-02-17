.PHONY: build run

build:
	@clear
	@date
	time go build cmd/main.go
	@date

run:
	@clear
	@date
	time go run cmd/main.go
	@date
