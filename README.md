# Legacy (do not follow)

#Steps

1. checkout ev-user-service and `docker-compose build`
2. in this repo, `docker-compose build`
3. `docker ps` will see the following

```
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS          PORTS                               NAMES
27e71b3afcbe   ev-user-service:latest      "./main"                 16 seconds ago   Up 14 seconds   0.0.0.0:8080->8080/tcp              ev-user-service
bb8d654d997d   ev-provider-service:latest   "./main"                 2 minutes ago    Up 43 seconds   0.0.0.0:8081->8081/tcp              ev-provider-service
5b04e4865a57   adminer                     "entrypoint.sh docke…"   2 minutes ago    Up 43 seconds   0.0.0.0:4000->8080/tcp              ev-adminer
c3d006ed5890   mysql:latest                "docker-entrypoint.s…"   2 minutes ago    Up 43 seconds   0.0.0.0:3306->3306/tcp, 33060/tcp   ev-database
```

After you are done, you can launch a request
