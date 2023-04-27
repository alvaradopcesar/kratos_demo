package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

// Customer is a Customer model.
type Customer struct {
	ID   string
	Name string
}

// CustomerRepo is a Greater repo.
type CustomerRepo interface {
	Save(context.Context, *Customer) (*Customer, error)
	Update(context.Context, *Customer) (*Customer, error)
	FindByID(context.Context, string) (*Customer, error)
	ListAll(context.Context) ([]*Customer, error)
}

// CustomerUsecase is a Customer usecase.
type CustomerUsecase struct {
	repo CustomerRepo
	log  *log.Helper
}

// NewCustomerUsecase new a Customer usecase.
func NewCustomerUsecase(repo CustomerRepo, logger log.Logger) *CustomerUsecase {
	return &CustomerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCustomer creates a Customer, and returns the new Customer.
func (uc *CustomerUsecase) CreateCustomer(ctx context.Context, g *Customer) (*Customer, error) {
	uc.log.WithContext(ctx).Infof("CreateCustomer: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

func (uc *CustomerUsecase) GetCustomerById(ctx context.Context, g *Customer) (*Customer, error) {
	uc.log.WithContext(ctx).Infof("GetCustomerById %v", g.ID)
	return uc.repo.FindByID(ctx, g.ID)
}

func (uc *CustomerUsecase) GetCustomerAll(ctx context.Context) ([]*Customer, error) {
	uc.log.WithContext(ctx).Infof("GetCustomerAll ")
	return uc.repo.ListAll(ctx)
}
