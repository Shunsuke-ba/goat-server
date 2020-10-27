GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix
DOCKER-COMPOSE=docker-compose
DOCKER=docker
CONTAINER_NAME=goat_dev
DEPLOY_FILE=app.yml

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./mod
run:
	$(GORUN) mod/app/main.go mod/app/wire_gen.go mod/app/wire_set.go
test:
	$(GORTEST) -v ./...
db:
	$(DOCKER-COMPOSE) up -d
rm:
	$(DOCKER) stop $(CONTAINER_NAME)
	$(DOCKER) rm $(CONTAINER_NAME)
sql:
	$(DOCKER) exec -it $(CONTAINER_NAME) mysql -u root -p
deploy:
	gcloud app deploy $(DEPLOY_FILE)
