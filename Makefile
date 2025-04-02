# get
build-%:
	cd src/cmd/lambda/get/ && \
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go

start-api: build-get
	sam local start-api --env-vars local-env.json
