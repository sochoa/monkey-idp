IMAGE_PREFIX := idp
IMAGE_VERSION ?= latest
UI_IMAGE := $(IMAGE_PREFIX)_ui:$(IMAGE_VERSION)
API_IMAGE := $(IMAGE_PREFIX)_api:$(IMAGE_VERSION)
POSTGRES_PASSWORD := $(shell awk '/POSTGRES_PASSWORD/{print $$NF}' docker-compose.yml | tail -n1 | sed 's/"//g')

api-image:
	docker build -t $(API_IMAGE) -f ./docker/api.Dockerfile ./api

ui-image:
	docker build -t $(UI_IMAGE) -f ./docker/ui.Dockerfile ./ui

api-pw:
	echo "$(POSTGRES_PASSWORD)"

api-run:
	docker run                                  \
		--rm                                      \
		-e POSTGRES_USER=idp_user                 \
		-e POSTGRES_DB=idp                        \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		--network=host                            \
		-it                                       \
		-v $(PWD)/api:/src                        \
		-w /src                                   \
		$(API_IMAGE)

api-exec:
	docker run --rm --network=host -it -v $(PWD)/api:/src -w /src $(API_IMAGE) bash

ui-exec:
	docker run --rm --network=host -it -v $(PWD)/ui:/src -w /src $(UI_IMAGE) bash

ui-run:
	docker run --rm --network=host -it -v $(PWD)/ui:/src -w /src $(UI_IMAGE)

build: api-image ui-image

run:
	docker-compose up
