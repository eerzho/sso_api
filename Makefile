.PHONY: build up down logs restart swagger

build:
	docker compose build

up:
ifdef name
	docker compose up -d $(name)
else 
	docker compose up -d
endif
	$(MAKE) swagger

down:
ifdef name
	docker compose down $(name)
else
	docker compose down
endif

logs:
ifdef name
	docker compose logs -f $(name)
else 
	docker compose logs -f
endif

restart:
ifdef name
	$(MAKE) down name=$(name)
	$(MAKE) up name=$(name)
else 
	$(MAKE) down
	$(MAKE) up
endif

swagger:
	docker compose exec http swag init -g ./internal/handler/handler.go