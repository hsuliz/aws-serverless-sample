# find get
build-%:
	cd src/cmd/lambda/$* && \
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go

start-api: build-find build-get build-post build-patch build-delete
	sam local start-api --env-vars local-env.json
