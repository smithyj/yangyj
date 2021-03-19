build-debug:
	go build -gcflags "all=-N -l" -o ./tmp/serve ./cmd/serve/main.go

run-dev:
	lsof -i:2345 | grep dlv | awk '{print $2}' | xargs kill -9
	ENV=dev dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./tmp/serve