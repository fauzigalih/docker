### Install Portainer for Monitoring Docker

Download docker image portainer from `https://hub.docker.com` :<br>
```
docker image pull portainer/portainer:latest
```

Create docker container with port forwarding:
```
docker container create --name portainer --publish 9000:9000 --volume /var/run/docker.sock:/var/run/docker.sock portainer/portainer:latest
```

Running docker container portainer:
```
docker container start portainer
```

See docker container with status running:
```
docker container ls
```
<img width="1183" alt="image" src="https://github.com/fauzigalih/docker/assets/64176403/800b1ef0-a6a5-4873-bf33-1d5049df296c">
<br><br><br>

Lets open the browser with port 9000, example `localhost:9000`<br>
<img width="822" alt="image" src="https://github.com/fauzigalih/docker/assets/64176403/ed3d77ee-2349-4729-91ae-ab0c0983bc2d">
<img width="1436" alt="image" src="https://github.com/fauzigalih/docker/assets/64176403/b107e79c-2574-4001-9734-45f6d3a45a63">
<img width="1439" alt="image" src="https://github.com/fauzigalih/docker/assets/64176403/a1a14719-152a-4335-addc-2e6e91eaec57">
<img width="1439" alt="image" src="https://github.com/fauzigalih/docker/assets/64176403/6f2ec663-5363-4e16-9959-867e42d54aab">
<br><br>
Yeayy, Finish..
