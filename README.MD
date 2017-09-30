# gRPC example

This section explains how to use Traefik as reverse proxy for gRPC application with self-signed certificates.

Server backend.local listen on port :50051
Traefik frontend.local listen on port :4443 and rediret to grpc server

## gRPC Server certificat

Autosigned certificate has been generated for frontend.local and backend.local

## Launch

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

