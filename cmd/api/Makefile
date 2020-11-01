GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix
CONTAINER_NAME=goat_dev
DEPLOY_FILE=app.yml

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./mod
run:
	$(GORUN) main.go wire_gen.go
test:
	$(GORTEST) -v ./...
sql:
	$(DOCKER) exec -it $(CONTAINER_NAME) mysql -u root -p
deploy:
	gcloud app deploy $(DEPLOY_FILE)
