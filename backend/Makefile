build-images:
	docker build -f Dockerfile.main -t umutcan-case-main:latest
	docker build -f Dockerfile.auth -t umutcan-case-auth:latest
	docker build -f Dockerfile.weather -t umutcan-case-weather:latest

run-images:
	docker run -d umutcan-case-main:latest
	docker run -d umutcan-case-auth:latest
	docker run -d umutcan-case-weather:latest

sudo-build-images:
	sudo docker build -f Dockerfile.main -t umutcan-case-main:latest
	sudo docker build -f Dockerfile.auth -t umutcan-case-auth:latest
	sudo docker build -f Dockerfile.weather -t umutcan-case-weather:latest

sudo-run-images:
	sudo docker run -d umutcan-case-main:latest
	sudo docker run -d umutcan-case-auth:latest
	sudo docker run -d umutcan-case-weather:latest