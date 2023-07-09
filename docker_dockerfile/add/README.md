## ADD Instruction
for Adding file host to docker image, if file compressed like zip, rar, tar.gz its will uncompressed / extract to docker image destination. 

#### Requirement image
```
docker pull alpine:3.18
```

Running build image:
```
docker build -t fauzigalih/add:1.0 docker_dockerfile/add
```
