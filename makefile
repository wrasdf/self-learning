.PHONY: ut apiut shui shapi api ui e2e stop clean eks

DCR = docker-compose run --rm

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
	# docker volume rm $(docker volume ls -f driver=local | awk '{print $2}' | tail -n+2)

db-migrate:
	docker-compose build flyway
	docker-compose run --rm flyway

apply-%:
	$(DCR) ctpl validate -p cfns/envs/$(env).yaml -c $(*)
	$(DCR) ctpl apply -p cfns/envs/$(env).yaml -c $(*)

update-ui-%:
	$(DCR) aws s3 cp ./ui/ s3://employee.apollo-dev.platform.myobdev.com/ --recursive
	$(DCR) aws cloudfront create-invalidation --distribution-id EZYDXUK3CCPZH --paths /index.html /error.html

eks:
	$(DCR) stackup aws-eks-lab-stack up -t eks/template.yaml -p eks/lab.yaml

config:
	aws s3 cp s3://aws-eks-lab-stack-eksstack-1esvt-kubeconfigbucket-1602ul3nn0nob/.kube/config.enc ~/.kube/

e2e:
	echo "TODO: This is E2E test."
