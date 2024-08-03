package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olezhek28/microservices_course/week_6/jwt/internal/model"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	descAccess "https://github.com/MentalMentos/jwt_tls_tokens/tree/136c057d131d05ec473b16427178bbc4ec82e69a/tokens/jwt/internal/model"
)

var accessToken = flag.String("a", "", "access token")

const servicePort = 50051

func main() {
	flag.Parse()

	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: model.ExamplePath,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Access granted")
}
