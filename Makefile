dev-dockerfile = -f docker-compose.yml -f docker-compose.dev.yml
prod-dockerfile = -f docker-compose.yml

.PHONY: build-dev
build-dev:
	docker-compose $(dev-dockerfile) build
	$(MAKE) install-go-dependencies

.PHONY: build-prod
build-prod:
	docker-compose $(prod-dockerfile) build
	$(MAKE) install-go-dependencies

.PHONY: install-go-dependencies
install-go-dependencies:
	docker-compose run --rm --no-deps backend dep ensure

.PHONY: dev
dev:
	docker-compose $(dev-dockerfile) up --remove-orphans

.PHONY: prod
prod:
	docker-compose $(prod-dockerfile) up --remove-orphans
