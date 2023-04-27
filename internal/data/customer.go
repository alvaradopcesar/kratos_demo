package data

import (
	"context"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	"errors"
)

type customerRepo struct {
	data  *CustomerData
	log   *log.Helper
	cache Cache
}

// NewCustomerRepo .
func NewCustomerRepo(logger log.Logger) biz.CustomerRepo {
	return &customerRepo{
		//data: data,
		log:   log.NewHelper(logger),
		cache: NewCacheCustomer(),
	}
}

func (r *customerRepo) Save(ctx context.Context, g *biz.Customer) (*biz.Customer, error) {
	r.log.WithContext(ctx).Infof("customerRepo Save %v", g)
	exist, err := r.cache.Exists(g.ID)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("id already exists")
	}

	err = r.cache.Set(CustomerData{ID: g.ID, Name: g.Name})
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *customerRepo) Update(ctx context.Context, g *biz.Customer) (*biz.Customer, error) {
	return g, nil
}

func (r *customerRepo) FindByID(ctx context.Context, id string) (*biz.Customer, error) {
	r.log.WithContext(ctx).Infof("customerRepo FindByID %v", id)
	cus, err := r.cache.Get(id)
	if err != nil {
		return nil, err
	}
	return &biz.Customer{
		ID:   cus.ID,
		Name: cus.Name,
	}, nil
}

func (r *customerRepo) ListAll(ctx context.Context) ([]*biz.Customer, error) {
	r.log.WithContext(ctx).Infof("customerRepo ListAll")
	customerData, err := r.cache.GetAllKeys()
	if err != nil {
		return nil, err
	}
	var cus []*biz.Customer
	for _, data := range customerData {
		cus = append(cus, &biz.Customer{
			ID:   data.ID,
			Name: data.Name,
		})
	}
	return cus, nil
}
