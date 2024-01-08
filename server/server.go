package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	train "github.com/thejas01/train-ticketing/proto"
	"google.golang.org/grpc"
)

type trainServer struct {
	mu       sync.Mutex
	sectionA map[string]*train.User
	sectionB map[string]*train.User
	receipts map[string]*train.Receipt
}

func (s *trainServer) PurchaseTicket(ctx context.Context, req *train.TicketRequest) (*train.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pricePaid := 20.0

	receipt := &train.Receipt{
		From:      req.From,
		To:        req.To,
		User:      req.User,
		PricePaid: pricePaid,
	}

	s.receipts[req.User.Email] = receipt

	return receipt, nil
}

func (s *trainServer) AllocateSeat(ctx context.Context, req *train.SeatAllocationRequest) (*train.SeatAllocationResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var seat string
	switch req.Section {
	case "A":
		seat = allocateSeat(s.sectionA)
	case "B":
		seat = allocateSeat(s.sectionB)
	default:
		return nil, fmt.Errorf("unknown section: %s", req.Section)
	}

	return &train.SeatAllocationResponse{Seat: seat}, nil
}

func allocateSeat(sectionSeats map[string]*train.User) string {

	for i := 1; i <= 50; i++ {
		seat := fmt.Sprintf("%d", i)
		if _, occupied := sectionSeats[seat]; !occupied {
			return seat
		}
	}

	return ""
}

func (s *trainServer) GetReceiptDetails(ctx context.Context, req *train.ReceiptRequest) (*train.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userID := req.UserId
	receipt, found := s.receipts[userID]
	if !found {
		return nil, fmt.Errorf("receipt not found for user: %s", userID)
	}

	return receipt, nil
}

func (s *trainServer) ViewUsersBySection(req *train.ViewUsersRequest, stream train.TrainService_ViewUsersBySectionServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var sectionUsers map[string]*train.User
	switch req.Section {
	case "A":
		sectionUsers = s.sectionA
	case "B":
		sectionUsers = s.sectionB
	default:
		return fmt.Errorf("unknown section: %s", req.Section)
	}

	for seat, user := range sectionUsers {
		if err := stream.Send(&train.SeatDetails{
			Seat: seat,
			User: user,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *trainServer) RemoveUser(ctx context.Context, req *train.RemoveUserRequest) (*train.RemoveUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userID := req.UserId

	if _, found := s.sectionA[userID]; found {
		delete(s.sectionA, userID)
		log.Printf("User %s removed from Section A", userID)
		return &train.RemoveUserResponse{Success: true}, nil
	}

	if _, found := s.sectionB[userID]; found {
		delete(s.sectionB, userID)
		log.Printf("User %s removed from Section B", userID)
		return &train.RemoveUserResponse{Success: true}, nil
	}

	return &train.RemoveUserResponse{Success: false}, fmt.Errorf("user not found: %s", userID)
}

func (s *trainServer) ModifySeat(ctx context.Context, req *train.ModifySeatRequest) (*train.ModifySeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userID := req.UserId
	newSeat := req.NewSeat

	if user, found := s.sectionA[userID]; found {
		oldSeat := fmt.Sprintf("A_%d", user)
		s.sectionA[userID] = &train.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
		log.Printf("User %s moved from %s to %s", userID, oldSeat, newSeat)
		return &train.ModifySeatResponse{Success: true}, nil
	}

	if user, found := s.sectionB[userID]; found {
		oldSeat := fmt.Sprintf("B_%d", user)
		s.sectionB[userID] = &train.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
		log.Printf("User %s moved from %s to %s", userID, oldSeat, newSeat)
		return &train.ModifySeatResponse{Success: true}, nil
	}

	return &train.ModifySeatResponse{Success: false}, fmt.Errorf("user not found: %s", userID)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server started")
	server := grpc.NewServer()
	Service := &trainServer{
		sectionA: make(map[string]*train.User),
		sectionB: make(map[string]*train.User),
		receipts: make(map[string]*train.Receipt),
	}

	train.RegisterTrainServiceServer(server, Service)

	log.Println("Server listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
