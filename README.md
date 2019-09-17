
```sh
docker build --no-cache -t cactus-core .
docker run -it -d -v $(pwd):/go/src/app -p 9040:9040 --name cactus-core cactus-core
```