package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc/codes"

	"github.com/grpc-demo/chat_app/chatpb"

	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

type Connection struct {
	stream chatpb.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type Server struct {
	Connections map[int64][]*Connection
}

func (s *Server) CreateStream(pcon *chatpb.Connect, stream chatpb.Broadcast_CreateStreamServer) error {
	grpcLog.Infoln("CreateStream called")
	conn := &Connection{
		stream: stream,
		id:     pcon.User.Id,
		active: true,
		error:  make(chan error),
	}
	chatroom := pcon.User.GetChatRoomNo()

	if pcon.WantsNewRoom {
		if _, ok := s.Connections[chatroom]; !ok {
			newConnection := []*Connection{conn}
			s.Connections[chatroom] = newConnection
			grpcLog.Infoln("New chatroom created")
		} else {
			grpcLog.Infoln("Chatroom already exists")
			return status.Errorf(
				codes.AlreadyExists,
				fmt.Sprintf("The Chatroom already exists"),
			)
		}
	} else {
		if _, ok := s.Connections[chatroom]; ok {
			s.Connections[chatroom] = append(s.Connections[chatroom], conn)
			grpcLog.Infoln("New user added to chatroom - ", chatroom)
		} else {
			grpcLog.Infoln("Chatroom does not exists")
			return status.Errorf(
				codes.NotFound,
				fmt.Sprintf("The Chatroom does not exist"),
			)
		}
	}

	// if _, ok := s.Connections[chatroom]; ok {
	// 	s.Connections[chatroom] = append(s.Connections[chatroom], conn)
	// } else {
	// 	newConnection := []*Connection{conn}
	// 	s.Connections[chatroom] = newConnection
	// }

	return <-conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *chatpb.Message) (*chatpb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)
	chatroomToJoin := msg.GetChatRoom()
	for _, conn := range s.Connections[chatroomToJoin] {
		wait.Add(1)

		go func(msg *chatpb.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending msg to stream:", conn.stream)

				if err != nil {
					grpcLog.Errorf("Error with stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &chatpb.Close{}, nil
}

func main() {
	fmt.Println("Server has started")
	Connections := make(map[int64][]*Connection)

	server := &Server{Connections}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Error creating the server:%v \n", err)
	}

	chatpb.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)
}
