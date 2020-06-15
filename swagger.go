package blogapi

import (
	"context"
	"log"

	swagger "github.com/sm43/goa-crud/gen/swagger"
)

// swagger service example implementation.
// The example methods log the requests and return zero values.
type swaggersrvc struct {
	logger *log.Logger
}

// NewSwagger returns the swagger service implementation.
func NewSwagger(logger *log.Logger) swagger.Service {
	return &swaggersrvc{logger}
}

// Add a new blog
func (s *swaggersrvc) Sm1(ctx context.Context) (res *swagger.Resource, err error) {

	res = &swagger.Resource{
		ID:   1,
		Name: "shivam",
	}
	s.logger.Print("swagger.sm1")
	return res, nil
}

// Add a new blog
func (s *swaggersrvc) Sm2(ctx context.Context) (res *swagger.Resource, err error) {
	s.logger.Print("shivam....")

	res = &swagger.Resource{
		ID:   12,
		Name: "ad",
		Type: "adsa",

		Rating: 3,
	}

	s.logger.Print("swagger.sm2")
	return res, nil
}

// Add a new blog
func (s *swaggersrvc) Sm3(ctx context.Context) (res *swagger.Resource, err error) {
	res = &swagger.Resource{
		Name: "sm",
		ID:   23,
		Rating: 3,
	}
	s.logger.Print("swagger.sm3")
	return
}
