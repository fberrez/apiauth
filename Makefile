SERVICE_PORT?=8080
NETWORK?=apiauth_local

start-containers:
	docker-compose up -d

stop-containers:
	docker-compose down

watch:
	$(eval PACKAGE_NAME=$(shell head -n 1 go.mod | cut -d ' ' -f2))
	docker run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(shell pwd):/go/src/$(PACKAGE_NAME) -p $(SERVICE_PORT):$(SERVICE_PORT) --network $(NETWORK) cosmtrek/air
