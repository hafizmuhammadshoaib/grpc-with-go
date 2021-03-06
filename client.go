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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	userProto "example.com/grpc_with_go/userproto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userProto.NewUserProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &userProto.UserCreateRequest{Name: "Hammad", Email: "hammad@gmail.com", Phone: "03367887046"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.GetMessage())

	reply, errInGet := c.GetUserById(ctx, &userProto.UserGetRequest{Id: 1})
	if errInGet != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(reply.GetEmail())
}
