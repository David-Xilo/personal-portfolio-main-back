# safehouse-tech-back
Backend microservice for tech submodule

# Run locally
go run src/cmd/main.go

# Docker

## start container

docker build -t safehouse-main-back-container .

docker run --name safehouse-main-back-container -p 4000:4000 -d safehouse-main-back-container

## stop container

docker stop safehouse-main-back-container

docker rm safehouse-main-back-container


## see logs

docker logs safehouse-main-back-container -f

## shell the container

docker exec -it safehouse-back-container /bin/bash
