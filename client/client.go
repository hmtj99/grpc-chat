package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/grpc-demo/chat_app/chatpb"
	"google.golang.org/grpc"
)

var client chatpb.BroadcastClient
var wait *sync.WaitGroup

func init() {
	wait = &sync.WaitGroup{}
}

func connect(user *chatpb.User) error {
	var streamError error
	stream, err := client.CreateStream(context.Background(), &chatpb.Connect{
		User:         user,
		Active:       true,
		WantsNewRoom: false,
	})

	if err != nil {
		return fmt.Errorf("Connection failed:%v", err)
	}

	wait.Add(1)
	go func(str chatpb.Broadcast_CreateStreamClient) {
		defer wait.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamError = fmt.Errorf("Error while reading the msg:%v", err)
				break
			} else {
				fmt.Println(msg.Id, ":", msg.Content)
			}
		}
	}(stream)

	return streamError
}

func main() {
	fmt.Println("Starting chat client")
	timeStamp := time.Now()
	done := make(chan int)

	//	newchatroom := flag.Int("chatroom", 2, "Chatroom to enter")
	name := flag.String("name", "Anon", "Enter your name")
	chatroom := flag.Int64("cr", 1, "Chatroom to connect to")

	flag.Parse()
	fmt.Println("Flags have been parsed")
	fmt.Println("Chartoom selected is", *chatroom)
	fmt.Println("Name is", *name)
	id := sha256.Sum256([]byte(timeStamp.String() + *name))

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}

	client = chatpb.NewBroadcastClient(conn)

	user := &chatpb.User{
		Id:         hex.EncodeToString(id[:]),
		Name:       *name,
		ChatRoomNo: *chatroom,
	}
	err_connect := connect(user)

	if err_connect != nil {
		fmt.Println("Error creating stream ", err)
		return
	}

	wait.Add(1)
	go func() {
		defer wait.Done()

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			msg := &chatpb.Message{
				Id:        user.Name,
				Content:   scanner.Text(),
				Timestamp: timeStamp.String(),
				ChatRoom:  user.GetChatRoomNo(),
			}

			_, err := client.BroadcastMessage(context.Background(), msg)

			if err != nil {
				fmt.Printf("Error sending the message: %v", err)
				break
			}
		}
	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}
