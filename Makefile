.PHONY: ond

ond:
	go build -o ./bin/ond .

docker:
	docker build -t ytgy:latest .
