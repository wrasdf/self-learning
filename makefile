.PHONY: ut apiut shui shapi api ui e2e

ut:
	docker-compose -f docker-compose-test.yaml build ut
	docker-compose -f docker-compose-test.yaml run --rm ut

apiut:
	docker-compose -f docker-compose-test.yaml build apiut
	docker-compose -f docker-compose-test.yaml run --rm apiut

ui:
	docker-compose build ui
	docker-compose run --rm --service-ports ui

api:
	docker-compose build api
	docker-compose run --rm --service-ports api

shui:
	docker-compose build shui
	docker-compose run --rm shui

shapi:
	docker-compose build shapi
	docker-compose run --rm shapi

e2e:
	echo "TODO: This is E2E test."
