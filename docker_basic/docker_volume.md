### Docker Mount & Volume Documentation

#### Docker Mount
Docker mount type bind:<br>
syntax: `docker container create --name containername --mount "type=bind,source=folderHost,destination=folderEnv,readonly" image:tag`<br>
```
docker container create --name portainer --mount "type=bind,source=/home/admins/docker/volume/portainer,destination=/var/run/docker.sock --publish 9000:9000 portainer/portainer:latest
```

#### Docker Volume [Recomended]
Create docker volume:
syntax: `docker volume create volumename`
```
docker volume create mysql-data
```

Mount volume in docker container:
syntax: `docker container create --name containername --mount "type=volume,source=folderhost,destination=folderEnv" image:tag`
```
docker container --name mysql --mount "type=volume,source=mysql-data,destination=/var/lib/mysql" mysql:8.0
```

Simple syntax docker volume:
syntax: `docker container create --name containername --volume sourcevolume:destinationvolume image:tag`
```
docker container --name mysql --volume mysql-data:/var/lib/mysql" mysql:8.0
```

See docker volume:
```
docker volume ls
```

Delete docker volume:
syntax: `docker volume rm volumename`
```
docker volume rm mysql-data
```

