# Cactus Core service

Core service for the Cactus app.

```sh
docker build --no-cache -t cactus-core .
docker run -it -d -v $(pwd):/go/src/app -p 9040:9040 --name cactus-core cactus-core
```

## License
MIT © [Jan Guzman](https://github.com/Krystian19)
