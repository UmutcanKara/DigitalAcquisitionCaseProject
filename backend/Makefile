
build:
	go mod tidy
	sudo docker build -f Dockerfile.auth -t umutcan-case-auth:latest .
	sudo docker build -f Dockerfile.weather -t umutcan-case-weather:latest .

run-auth:
	sudo docker run -p 8080:8080 umutcan-case-auth:latest
run-weather:
	sudo docker run -p 8081:8081 umutcan-case-weather:latest