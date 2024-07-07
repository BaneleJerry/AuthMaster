build:
	go build -o out/build cmd/authmaster/main.go

run:build
	./out/build 