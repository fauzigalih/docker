### Docker Volume Backup Documentation

#### Backup Volume
syntax: `docker container run --rm --name containerName --mount "type=bind,source=sourceHost,destination=destinationBind" --mount "type=volume,source=targetVolumeBackup,destination=destinationVolume" image:tag tar cvf destinationBind/backup.tar.gz destinationVolume`
```
docker container run --rm --name backup-mongodb --mount "type=bind,source=/home/admins/docker/backup,destination=/backup" --mount "type=volume,source=mongodb-volume,destination=/data" ubuntu:latest tar cvf /backup/backup.tar.gz /data
```
note: backup file in `sourceHost` and image ubuntu recomended for backup volume

#### Restore Volume
syntax: `docker container run --rm --name containerName --mount "type=bind,source=sourceHost,destination=destinationBind" --mount "type=volume,source=targetVolumeRestore,destination=destinationVolume" ubuntu:latest bash -c "cd destinationVolume && tar xvf destinationBind/backup.tar.gz --strip 1`
```
docker container run --rm --name restore-mongodb --mount "type=bind,source=/home/admins/docker/backup,destination=/backup" --mount "type=volume,source=mongodb-restore,destination=/data" ubuntu:latest bash -c "cd /data && tar xvf /home/admins/docker/backup/backup.tar.gz --strip 1"
```


