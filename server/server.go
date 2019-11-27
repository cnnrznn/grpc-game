package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "math/rand"
    "net"

    pb "github.com/cnnrznn/grpc-game/proto"
)

type gameServer struct {
    players map[int64]*Coord
}

type Coord struct {
    x, y int
}

func (s *gameServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
    var r int64

    for {
        r = rand.Int63()
        if _, ok := s.players[r]; !ok {
            s.players[r] = &Coord{0,0}
            break
        }
    }

    return &pb.JoinResponse{PlayerId: r}, nil
}

func (s *gameServer) Leave(ctx context.Context, req *pb.LeaveRequest) (*pb.Nil, error) {
    if _, ok := s.players[req.PlayerId]; ok {
        delete(s.players, req.PlayerId)
    }

    return &pb.Nil{}, nil
}

func (s *gameServer) Move(ctx context.Context, req *pb.MoveRequest) (*pb.Nil, error) {
    if p, ok := s.players[req.PlayerId]; ok {
        switch req.Dir {
        case pb.MoveRequest_UP:
            p.y++
        case pb.MoveRequest_DOWN:
            p.y--
        case pb.MoveRequest_LEFT:
            p.x--
        case pb.MoveRequest_RIGHT:
            p.x++
        }
    }

    return &pb.Nil{}, nil
}


func main() {
    lis, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Fatal("Failed to open listening socket")
    }

    grpcServer := grpc.NewServer()
    pb.RegisterGameServer(grpcServer, &gameServer{})
    grpcServer.Serve(lis)
}
