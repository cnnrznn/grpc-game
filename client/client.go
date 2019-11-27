package main

import (
    "context"
    "google.golang.org/grpc"
    "log"

    pb "github.com/cnnrznn/grpc-game/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
    if err != nil {
        log.Panic(err)
    }
    defer conn.Close()

    client := pb.NewGameClient(conn)

    // play the game
    resp, err := client.Join(context.Background(), &pb.JoinRequest{})
    if err != nil {
        log.Panic(err)
    }

    id := resp.PlayerId
    log.Println("Received PlayerId", id)

    for i:=0; i<100; i++ {
        client.Move(context.Background(), &pb.MoveRequest{PlayerId: id,
                                                          Dir: pb.MoveRequest_RIGHT})
    }

    log.Println("Leaving Game")
    client.Leave(context.Background(), &pb.LeaveRequest{PlayerId: id})
}
