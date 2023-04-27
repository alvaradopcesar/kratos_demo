package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	uc *biz.CustomerUsecase
}

func NewCustomerService(uc *biz.CustomerUsecase) *CustomerService {
	return &CustomerService{uc: uc}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerReply, error) {
	_, err := s.uc.CreateCustomer(ctx, &biz.Customer{
		ID:   req.CreateCustomerBody.Id,
		Name: req.CreateCustomerBody.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCustomerReply{}, nil
}
func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerReply, error) {
	customer, err := s.uc.GetCustomerById(ctx, &biz.Customer{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &pb.GetCustomerReply{
		Id:   customer.ID,
		Name: customer.Name,
	}, nil
}

func (s *CustomerService) ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (*pb.ListCustomerReply, error) {
	cus, err := s.uc.GetCustomerAll(ctx)
	if err != nil {
		return nil, err
	}
	var getCustomerReply []*pb.GetCustomerReply
	for _, data := range cus {
		getCustomerReply = append(getCustomerReply, &pb.GetCustomerReply{
			Id:   data.ID,
			Name: data.Name,
		})
	}
	return &pb.ListCustomerReply{
		GetCustomers: getCustomerReply,
	}, nil
}
