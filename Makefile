GOAT_DEV=goat_dev

db:
	docker-compose up -d
rm:
	docker stop $(GOAT_DEV)
	docker rm $(GOAT_DEV)