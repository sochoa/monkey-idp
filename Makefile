IMAGE_PREFIX := idp
IMAGE_VERSION ?= latest
UI_IMAGE := $(IMAGE_PREFIX)_ui:$(IMAGE_VERSION)
API_IMAGE := $(IMAGE_PREFIX)_api:$(IMAGE_VERSION)

api-image:
	docker build -t $(API_IMAGE) -f ./docker/api.Dockerfile .

ui-image:
	docker build -t $(UI_IMAGE) -f ./docker/ui.Dockerfile .

api-exec:
	docker run -it -v $(PWD)/api:/src -w /src $(API_IMAGE) bash

ui-exec:
	docker run -it -v $(PWD)/ui:/src -w /src $(UI_IMAGE) bash

build: api-image ui-image
