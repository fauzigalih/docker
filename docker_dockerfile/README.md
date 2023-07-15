## Dockerfile Documentation
Dockerfile is file script to create docker image.

### General Syntax Dockerfile
```dockerfile
# FROM Instruction, to use docker image with tag
FROM alpine:3.18

# RUN Instruction, running script on build docker image progress.
RUN apk nano
RUN mkdir app
```
