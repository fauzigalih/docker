## Docker Ignore 
ignore file or folder on copying to folder image destination.

#### Requirement image
```
docker pull alpine:3.18
```

Running build image:
```
docker build -t fauzigalih/ignore:1.0 docker_dockerfile/ignore
```

Running container and see file destination:
```
docker run -d --name copy fauzigalih/ignore:1.0
```
```
docker container logs ignore
```