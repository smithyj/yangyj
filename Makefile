build-backend-dev:
	rm -rf ./backend/tmp
	go build -o ./backend/tmp/srv/main ./backend/cmd/srv/main.go