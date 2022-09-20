package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/pemarsao/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cound not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Pedro",
		Email: "pedro.marsao@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Cound not make gRPC request: %v", err)
	}

	fmt.Println(res)

}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Pedro",
		Email: "pedro.marsao@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Cound not make gRPC request: %v", err)
	}

	for {

		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cound not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status, "-", stream.User)

	}

}

func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{
		&pb.User{
			Id:    "p1",
			Name:  "Pedro",
			Email: "pedro@pedro.com",
		},
		&pb.User{
			Id:    "p2",
			Name:  "Pedro 2",
			Email: "pedro2@pedro.com",
		},
		&pb.User{
			Id:    "p3",
			Name:  "Pedro 3",
			Email: "pedro3@pedro.com",
		},
		&pb.User{
			Id:    "p4",
			Name:  "Pedro 4",
			Email: "pedro4@pedro.com",
		},
		&pb.User{
			Id:    "p5",
			Name:  "Pedro 5",
			Email: "pedro5@pedro.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "p1",
			Name:  "Pedro",
			Email: "pedro@pedro.com",
		},
		&pb.User{
			Id:    "p2",
			Name:  "Pedro 2",
			Email: "pedro2@pedro.com",
		},
		&pb.User{
			Id:    "p3",
			Name:  "Pedro 3",
			Email: "pedro3@pedro.com",
		},
		&pb.User{
			Id:    "p4",
			Name:  "Pedro 4",
			Email: "pedro4@pedro.com",
		},
		&pb.User{
			Id:    "p5",
			Name:  "Pedro 5",
			Email: "pedro5@pedro.com",
		},
	}

	go func() {

		for _, req := range reqs {
			fmt.Println("Send user:", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	wait := make(chan int)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Cound not receive the message: %v", err)
				break
			}
			fmt.Printf("Recebendo user %v, com status: %v\n", res.User.GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
