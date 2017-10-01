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

Server backend.local listen on port :50052

Linkerd frontend.local listen on port :50053 and rediret to grpc server

Insecure Mode supported here

## Launch Linkerd and grpc server

```
 docker-compose -f docker-compose-linkerd.yml up
 ```

Linkerd listen on port 50053 for incoming grpc connection

```
./greeter_server&
```

## Use it 

Launch Client

```
./greeter_client 
```


