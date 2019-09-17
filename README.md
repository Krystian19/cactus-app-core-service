
```sh
docker build --no-cache -t cactus-core .
docker run -it -d -v $(pwd):/go/src/app -p 8080:8080 --name cactus-core cactus-core
```