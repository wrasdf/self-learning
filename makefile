.PHONY: ut apiut shui shapi api ui e2e

ui:
	docker-compose -f docker-compose-dev.yaml build ui
	docker-compose -f docker-compose-dev.yaml run --rm --service-ports ui

api:
	docker-compose -f docker-compose-dev.yaml build api
	docker-compose -f docker-compose-dev.yaml run --rm --service-ports api

ut:
	docker-compose -f docker-compose-test.yaml build ut
	docker-compose -f docker-compose-test.yaml run --rm ut

apiut:
	docker-compose -f docker-compose-test.yaml build apiut
	docker-compose -f docker-compose-test.yaml run --rm apiut

shui:
	docker-compose -f docker-compose-dev.yaml build shui
	docker-compose -f docker-compose-dev.yaml run --rm shui

shapi:
	docker-compose -f docker-compose-dev.yaml build shapi
	docker-compose -f docker-compose-dev.yaml run --rm shapi

e2e:
	echo "TODO: This is E2E test."
