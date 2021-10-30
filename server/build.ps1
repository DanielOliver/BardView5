$ErrorActionPreference = "Stop"
docker-compose -f docker-compose-local.yml up -d
echo "sleeping 8"
Sleep 8
$env:BARDVIEW5_CONNECTION="postgresql://postgres:mysecretpassword@localhost/bardview5?sslmode=disable"
go run . migrate
$Env:BARDVIEW5_CONNECTION=""
echo "migrated"
docker-compose -f docker-compose-local.yml exec db /bin/bash -c 'pg_dump -U postgres -s bardview5 > /sql_dump/snapshot.sql'
echo "pg_dumped"
docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate
echo "sqlc generated"
# go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -o api/bardview5.go -package api -generate types,skip-prune bardview5.yaml
go generate ./...
echo "other generated"
docker-compose -f docker-compose-local.yml down
