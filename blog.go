package blogapi

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	blog "github.com/sm43/goa-crud/gen/blog"
)

// blog service example implementation.
// The example methods log the requests and return zero values.
type blogsrvc struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewBlog returns the blog service implementation.
func NewBlog(db *gorm.DB, logger *log.Logger) blog.Service {
	return &blogsrvc{db, logger}
}

// Add new blog and return its ID.
func (s *blogsrvc) Create(ctx context.Context, p *blog.CreatePayload) (err error) {

	err = VerifyJWT(p.Auth)
	if err != nil {
		s.logger.Println("Invalid user", err.Error())
		return blog.MakeInvalidToken(fmt.Errorf(err.Error()))
	}

	blog := &Blog{Name: p.Blog.Name}
	err = s.db.Create(blog).Error
	for _, comment := range p.Blog.Comments {
		if err := s.db.Model(&blog).Association("Comments").Append(&Comment{Text: comment.Comment}).Error; err != nil {
			return err
		}
	}
	return err
}

// List all entries
func (s *blogsrvc) List(ctx context.Context) (res []*blog.StoredBlog, err error) {

	var all []Blog
	if err = s.db.Preload("Comments").Find(&all).Error; err != nil {
		return nil, blog.MakeDbError(fmt.Errorf(err.Error()))
	}

	for _, r := range all {
		comments := []*blog.StoredComment{}
		for _, c := range r.Comments {
			comments = append(comments, &blog.StoredComment{
				ID:      c.ID,
				Comment: c.Text,
			})
		}
		res = append(res, &blog.StoredBlog{
			ID:       r.ID,
			Name:     r.Name,
			Comments: comments,
		})
	}

	return res, nil
}

// Show blog based on the id given
func (s *blogsrvc) Show(ctx context.Context, p *blog.ShowPayload) (res *blog.StoredBlog, err error) {

	b := Blog{}
	if err = s.db.Preload("Comments").First(&b, p.ID).Error; err != nil {
		return nil, blog.MakeDbError(fmt.Errorf(err.Error()))
	}

	comments := []*blog.StoredComment{}
	for _, c := range b.Comments {
		comments = append(comments, &blog.StoredComment{
			ID:      c.ID,
			Comment: c.Text,
		})
	}
	res = &blog.StoredBlog{
		ID:       b.ID,
		Name:     b.Name,
		Comments: comments,
	}
	return res, nil
}

// Delete a blog
func (s *blogsrvc) Remove(ctx context.Context, p *blog.RemovePayload) (err error) {

	if err = s.db.Unscoped().Where("id = ?", p.ID).Delete(Blog{}).Error; err != nil {
		return blog.MakeDbError(fmt.Errorf(err.Error()))
	}
	return
}

// Add a new comment for a blog
func (s *blogsrvc) Add(ctx context.Context, p *blog.AddPayload) (err error) {

	b := Blog{}
	if err = s.db.Where("id = ?", p.ID).First(&b).Error; err != nil {
		return blog.MakeDbError(fmt.Errorf(err.Error()))
	}
	if err = s.db.Model(&b).Association("Comments").Append(&Comment{Text: p.Comments.Comment}).Error; err != nil {
		return blog.MakeDbError(fmt.Errorf(err.Error()))
	}

	return
}
