ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

default:
	echo "Choose a target"
build: go-api-codegen go-genny go-bindata
build-win: build
	go build -o build/bardview5.exe .
build-linux: build
	go build -o build/bardview5 .
sqlc-generate:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate
go-api-codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -o api/bardview5.go -package api -generate types,skip-prune bardview5.yaml
go-genny:
	go install github.com/cheekybits/genny@latest
	genny -in=models/gen-api-models.go -out=models/api-models.go =pkg=models gen "ApiModel=RPG,RPGList"
go-bindata:
	go install github.com/jteeuwen/go-bindata/...@latest
	go-bindata -pkg main migrations

