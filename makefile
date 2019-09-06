.PHONY: ut apiut shui shapi api ui e2e stop clean

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

clean:
	docker stop $(shell docker ps -aq) && docker rm $(shell docker ps -aq)
	docker volume rm $(docker volume ls -f driver=local | awk '{print $2}' | tail -n+2)

migrate:
	docker-compose build flyway
	docker-compose run --rm flyway

e2e:
	echo "TODO: This is E2E test."
