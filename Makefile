# Usage:

release: clean build

gitsha = $(shell git rev-parse HEAD)

build: 
	sudo docker build -t ursho-builder .
	sudo docker run ursho-builder | sudo docker build -t godlzr/ursho:$(gitsha) -

# remove previous images and containers
clean:
	docker rm -f ursho-builder 2> /dev/null || true
	docker rmi -f ursho-builder || true
	docker rm -f ursho 2> /dev/null || true
	docker rmi -f ursho || true
