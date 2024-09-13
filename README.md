# safehouse-tech-back
Backend microservice for tech submodule

# Run locally
go run src/cmd/main.go

# Docker

## start container

docker build -t safehouse-back .

docker run --name safehouse-back-container -p 4000:4000 -d safehouse-back

## stop container

docker stop safehouse-back-container

docker rm safehouse-back-container


## see logs

docker logs safehouse-back-container -f

## shell the container

docker exec -it safehouse-back-container /bin/bash
