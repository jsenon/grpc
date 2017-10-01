# gRPC example

This section explains how to use Traefik and linkerd as reverse proxy for gRPC application 

# Traefik

With self-signed certificates.

Server backend.local listen on port :50051

Traefik frontend.local listen on port :4443 and rediret to grpc server

## gRPC Server certificat

Autosigned certificate has been generated for frontend.local and backend.local

## Launch traefik and grpc server

```
docker compose up -d
```

Launch Server

```
./server&
```

## Use it 

Launch Client

```
./client 
```

## Note/Issues

Google grpc server sample only serve on tcp 
Traefik redirect to an https server

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







