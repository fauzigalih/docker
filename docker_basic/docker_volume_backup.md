### Back up, restore, or migrate data volumes

#### Backup Volume
For example, create a new container named `dbstore` :
```
docker run -v /dbdata --name dbstore ubuntu /bin/bash
```
In the next command:

- Launch a new container and mount the volume from the `dbstore` container
- Mount a local host directory as `/backup`
- Pass a command that tars the contents of the `dbdata` volume to a `backup.tar` file inside our `/backup` directory
```
docker run --rm --volumes-from dbstore -v $(pwd):/backup ubuntu tar cvf /backup/backup.tar /dbdata
```
When the command completes and the container stops, it creates a backup of the `dbdata` volume.

note: backup file in `sourceHost` and image ubuntu recomended for backup volume

#### Restore Volume
With the backup just created, you can restore it to the same container, or to another container that you created elsewhere.

For example, create a new container named `dbstore2`:
```
docker run -v /dbdata --name dbstore2 ubuntu /bin/bash
```
Then, un-tar the backup file in the new container’s data volume:
```
docker run --rm --volumes-from dbstore2 -v $(pwd):/backup ubuntu bash -c "cd /dbdata && tar xvf /backup/backup.tar --strip 1"
```
You can use the techniques above to automate backup, migration, and restore testing using your preferred tools.
