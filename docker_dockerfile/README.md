## Dockerfile Documentation
Dockerfile is file script to create docker image.

#### General Syntax Dockerfile
```dockerfile
# FROM Instruction, to use docker image with tag
FROM alpine:3.18

# LABEL Instruction, for meta data docker image
LABEL author="Fauzi Galih"

# RUN Instruction, running script on build docker image progress.
RUN apk update
RUN apk add nano
RUN mkdir /app
RUN mkdir /app/${APP_DATA}

# WORKDIR Instruction, targeting directory root for instruction RUM, CMD, ENTRYPOINT, COPY, and ADD
WORKDIR /app/${APP_DATA}

# COPY Instruction, copy files to docker image
COPY /var/www/html/. /app/${APP_DATA}
COPY /var/www/html/. .

# EXPOSE Instruction, expose port of docker image
EXPOSE ${APP_PORT}

# ENV Instruction, environment make variable can be customize on create docker container
ENV ${APP_PORT}=8080
ENV ${APP_DATA}="blog"

# CMD Instruction, this command run on socker container running
CMD npm run dev

```
