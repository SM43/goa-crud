package blogapi

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	user "github.com/sm43/goa-crud/gen/user"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(db *gorm.DB, logger *log.Logger) user.Service {
	return &usersrvc{db, logger}
}

// Add a new blog
func (s *usersrvc) Create(ctx context.Context, p *user.CreatePayload) (err error) {
	err = VerifyJWT(p.Auth)
	if err != nil {
		s.logger.Println("Invalid user", err.Error())
		return user.MakeInvalidToken(fmt.Errorf(err.Error()))
	}
	userObj := &User{Name: p.User.Name, Age: p.User.Age, Class: p.User.Class}
	if err = s.db.Create(userObj).Error; err != nil {
		s.logger.Println("Db error ", err.Error())
		return user.MakeDbError(fmt.Errorf(err.Error()))
	}
	return nil
}

// List all the users
func (s *usersrvc) List(ctx context.Context) (res []*user.User, err error) {

	var all []User
	if err = s.db.Find(&all).Error; err != nil {
		return nil, user.MakeDbError(fmt.Errorf(err.Error()))
	}
	for _, r := range all {
		id := r.ID
		res = append(res, &user.User{
			ID:    &id,
			Name:  r.Name,
			Age:   r.Age,
			Class: r.Class,
		})
	}
	return res, nil
}
