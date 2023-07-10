## COPY Instruction
Copy file from host to docker image destination.

#### Requirement image
```
docker pull alpine:3.18
```

Running build image:
```
docker build -t fauzigalih/copy:1.0 docker_dockerfile/copy
```

Running container and see file destination:
```
docker run -d --name copy fauzigalih/copy:1.0
```
```
docker container logs copy
```