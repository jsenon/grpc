# gRPC example

This section explains how to use Traefik and linkerd as reverse proxy for gRPC application 

# Traefik

Server hello listen on port :7777

Server world listen on port :7778

Traefik listen on port :4441 and rediret to http server

## Launch traefik and grpc server

```
docker-compose -f docker-compose-traefik-grpc.yml up
```

Server Hello Run on 7777

Server World Run on 7778

```
client -> Traefik -> Hello -> traefik -> World
```

## Use it 

Launch Client

```
./curl localhost:4140/toto
```

## Note/Issues

Traefik redirect only to an http server over http2
Traefik don't manage grpc protocol. We can't define grpc route

# Linkerd

Linkerd Demo GRPC 

## Launch Linkerd and grpc server

```
 docker-compose -f docker-compose-linkerd.yml up
 ```

Server Hello Run on 7777

Server World Run on 7778

```
client -> Linkerd -> Hello -> Linkerd -> World
```

## Use it 

Launch Client

```
go get -u github.com/linkerd/linkerd-examples/docker/helloworld/helloworld-client
```

```
helloworld-client 127.0.0.1:4140
```

# Tips

Greeter client came form google.golan.org/grpc/examples has been modified to send request to frontend.local:4443

```
./greeter_client
```

Greeter Server came form google.golan.org/grpc/examples has been modified to received request to :50051

```
./greeter_server&
```

### gRPC Server certificat

Autosigned certificate has been generated for frontend.local and backend.local





