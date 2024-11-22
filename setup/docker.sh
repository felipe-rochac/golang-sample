sudo groupadd docker

sudo usermod -aG docker $USER

newgrp docker

curl -sSL https://get.docker.com/ | sh

docker buildx version

sudo apt-get install docker-buildx-plugin