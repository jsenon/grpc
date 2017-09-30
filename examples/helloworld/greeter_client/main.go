/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

const (
	address     = "frontend.local:4443"
	defaultName = "world"
)

func main() {

	// Read cert file
	FrontendCert, err := ioutil.ReadFile("./frontend.cert")
	fmt.Println("Cert Read:", FrontendCert)

	if err != nil {
		log.Fatalf("cert read error: %v", err)
		fmt.Println("cert read error:", err)

	}

	// Create CertPool
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(FrontendCert)

	// Create credentials
	credsClient := credentials.NewClientTLSFromCert(roots, "")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credsClient))
	fmt.Println("Connection:", conn)

	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Println("Connection error:", err)

	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	fmt.Println("GreeterClient:", c)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	fmt.Println("SayHello:", r)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
