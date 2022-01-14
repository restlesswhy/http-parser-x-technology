run:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build

delpg:
	sudo rm -Rfv pgdata/