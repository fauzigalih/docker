## Command Instruction
CMD Instruction is not running on processing build image, but running on container start.

#### Requirement pull image alpine:3.18
```
docker pull alpine:3.18
```

Running build image:
```
docker build -t fauzigalih/cmd:1.0 docker_dockerfile/command
```

Create docker container:
```
docker container create --name command fauzigalih/cmd:1.0
```
```
docker container start command
```
```
docker container logs command
```