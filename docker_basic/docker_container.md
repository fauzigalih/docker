### Docker Container Documentation

Show all docker container:
```
docker container ls -a
```
Show docker container only running:
```
docker container ls
```
Create docker container:

syntax: `docker container create --name containername imagename:tag`
```
docker container create --name api-service alpine:3.18
```
Create docker container with port forwarding:

syntax: `docker container create --name containername --publish porthost:portcontainer imagename:tag`
```
docker container create --name web-server --publish 8080:80 nginx:latest
```
Running docker container:

syntax: `docker container start containerID/containername`
```
docker container start api-service
```
Stop docker container:

syntax: `docker container stop containerID/containername`
```
docker container stop api-service
```
Delete docker container:

syntax: `docker container rm containerID/containername`
```
docker container rm api-service
```
See detail log docker container:

syntax: `docker container logs containerID/containername`
```
docker container logs api-service
```
See detail realtime log docker container:

syntax: `docker container logs -f containerID/containername`
```
docker container logs -f api-service
```
Execute / Login to docker container:

syntax: `docker container exec -i -t containerID/containername /bin/bash`
```
docker container exec -i -t redis /bin/bash
```
Add environment variable on docker container:

syntax: `docker container create --name containername --env KEY="value" imagename:tag`
```
docker container create --name mongo-db --env MONGO_INITDB_ROOT_USERNAME=root --env MONGO_INITDB_ROOT_PASSWORD=password mongo:latest
```
Monitoring docker container:
```
docker container stats
```