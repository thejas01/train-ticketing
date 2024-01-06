package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	train "github.com/thejas01/train-ticketing/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := train.NewTrainServiceClient(conn)

	receipt, err := client.PurchaseTicket(context.Background(), &train.TicketRequest{
		From: "London",
		To:   "France",
		User: &train.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
	})
	if err != nil {
		log.Fatalf("PurchaseTicket failed: %v", err)
	}

	fmt.Printf("Receipt: %+v\n", receipt)

	seatResponse, err := client.AllocateSeat(context.Background(), &train.SeatAllocationRequest{
		Section: "A",
		User: &train.User{
			FirstName: "Alice",
			LastName:  "Smith",
			Email:     "alice.smith@example.com",
		},
	})
	if err != nil {
		log.Fatalf("AllocateSeat failed: %v", err)
	}

	fmt.Printf("Allocated Seat: %s\n", seatResponse.Seat)

	viewStream, err := client.ViewUsersBySection(context.Background(), &train.ViewUsersRequest{
		Section: "A",
	})
	if err != nil {
		log.Fatalf("ViewUsersBySection failed: %v", err)
	}

	fmt.Println("Users in Section A:")
	for {
		userDetails, err := viewStream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Seat: %s, User: %+v\n", userDetails.Seat, userDetails.User)
	}

	removeUserResponse, err := client.RemoveUser(context.Background(), &train.RemoveUserRequest{
		UserId: "john.doe@example.com",
	})
	if err != nil {
		log.Fatalf("RemoveUser failed: %v", err)
	}

	if removeUserResponse.Success {
		fmt.Println("User removed successfully.")
	} else {
		fmt.Println("User not found.")
	}

	modifySeatResponse, err := client.ModifySeat(context.Background(), &train.ModifySeatRequest{
		UserId:  "alice.smith@example.com",
		NewSeat: "A_10",
	})
	if err != nil {
		log.Fatalf("ModifySeat failed: %v", err)
	}

	if modifySeatResponse.Success {
		fmt.Println("User's seat modified successfully.")
	} else {
		fmt.Println("User not found or seat modification failed.")
	}

	receiptDetails, err := client.GetReceiptDetails(context.Background(), &train.ReceiptRequest{
		UserId: "alice.smith@example.com",
	})
	if err != nil {
		log.Fatalf("GetReceiptDetails failed: %v", err)
	}

	fmt.Printf("Receipt Details: %+v\n", receiptDetails)
}
