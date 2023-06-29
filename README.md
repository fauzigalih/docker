# Docker Documentation

## Install Docker
Note: tutorial install docker for Ubuntu Server 20.04

First, update your existing list of packages:
```
$ sudo apt-get update
```
install a few prerequisite packages which let apt use packages over HTTPS:
```
$ sudo apt install apt-transport-https ca-certificates curl software-properties-common
```
Then add the GPG key for the official Docker repository to your system:
```
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
Add the Docker repository to APT sources:
```
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
```
Finally, install Docker:
```
$ sudo apt install docker-ce
```
Check status docker:
```
$ sudo systemctl status docker
```

## Execute docker comand without sudo
Adjust `username` with you have:
```
$ sudo usermod -aG docker username
```
```
$ su - username
```
Running docker without sudo:
```
$ docker
```
