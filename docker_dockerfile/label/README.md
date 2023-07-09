## LABEL Instruction
for Metadata Information, like company, author, website, etc

#### Requirement pull image alpine:3.18
```
docker pull alpine:3.18
```

Running build image:
```
docker build -t fauzigalih/label:1.0 docker_dockerfile/label
```

You can see label on inpect image:
```
docker image inspect fauzigalih/label:1.0
```
