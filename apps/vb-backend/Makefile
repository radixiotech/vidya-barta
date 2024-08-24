# Modules support
deps-reset:
	git checkout -- go.mod
	go mod tidy

tidy:
	go mod tidy

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

# Run Locally
build-vb-api:
	@go build -o bin/vb-api apps/vb-api/main.go

run-vb-api: build-vb-api
	@bin/vb-api