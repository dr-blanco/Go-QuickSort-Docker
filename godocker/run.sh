sudo docker rmi -f go-app-image
sudo docker build -t go-app-image .
sudo docker run go-app-image
