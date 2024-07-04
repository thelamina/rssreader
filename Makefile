all:
	@echo "***************************************************"
	@echo "**                rss reader tool                **"
	@echo "***************************************************"

run :
	sqlc generate && $(MAKE) build && ./server

build:
	go build -v -o server
