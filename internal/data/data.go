package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"helloworld/internal/conf"
)

// ProviderSet is data providers.
//var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCustomerRepo)
var ProviderSet = wire.NewSet(NewData, NewCacheCustomer, NewGreeterRepo, NewCustomerRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

type CustomerData struct {
	ID   string
	Name string
}

type Cache interface {
	Exists(id string) (bool, error)
	Set(CustomerData) error
	Get(key string) (CustomerData, error)
	GetAllKeys() ([]CustomerData, error)
}

type universalClient struct {
	DataInMemory map[string]CustomerData
}

func NewCacheCustomer() Cache {
	return &universalClient{
		DataInMemory: make(map[string]CustomerData),
	}
}

func (r *universalClient) Exists(key string) (bool, error) {
	c := r.DataInMemory[key]
	if c.ID == "" {
		return false, nil
	}
	return true, nil
}

func (r *universalClient) Set(customer CustomerData) error {
	r.DataInMemory[customer.ID] = customer
	return nil
}

func (r *universalClient) Get(key string) (CustomerData, error) {
	return r.DataInMemory[key], nil
}

func (r *universalClient) GetAllKeys() ([]CustomerData, error) {
	var customerList []CustomerData
	for _, data := range r.DataInMemory {
		customerList = append(customerList, CustomerData{
			ID:   data.ID,
			Name: data.Name,
		})
	}
	return customerList, nil
}
