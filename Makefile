# This will kick off the gqlgen codegen
gen:
	- go generate ./...

build:
	- CGO_ENABLED=0 GOOS=linux go build -a -gcflags='-N -l' -installsuffix cgo -o main .

start:
	- go run main.go

dev:
	- chmod +x tmp/main
	- air

lint:
	- go vet
