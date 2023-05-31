GO_VER = 1.20

tidy:
	@go mod tidy -go=${GO_VER}

gotest: tidy
	@go test -coverprofile=cover.out ./...
	@go tool cover -html=cover.out -o cover.html
