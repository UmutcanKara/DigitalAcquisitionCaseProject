build:
	go mod tidy
	docker build -f Dockerfile.auth -t umutcan-case-auth:latest
	docker build -f Dockerfile.weather -t umutcan-case-weather:latest

run:
	docker run -d umutcan-case-auth:latest
	docker run -d umutcan-case-weather:latest

sudo-build:
	go mod tidy
	sudo docker build -f Dockerfile.auth -t umutcan-case-auth:latest .
	sudo docker build -f Dockerfile.weather -t umutcan-case-weather:latest .

sudo-run-auth:
	sudo docker run -p 8080:8080 umutcan-case-auth:latest
sudo-run-weather:
	sudo docker run -p 8081:8081 umutcan-case-weather:latest