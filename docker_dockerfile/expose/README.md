## EXPOSE Instruction
Expose port of docker image.

#### Requirement image
```
docker pull golang:1.18-alpine
```

Running build image:
```
docker build -t fauzigalih/expose:1.0 docker_dockerfile/expose
```

Running container:
```
docker run -d --name expose -p 8080:8080 fauzigalih/expose:1.0
```

Open browser and acces `http://localhost:8080`