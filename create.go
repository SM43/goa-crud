package createapi

import (
	"context"
	create "crud/gen/create"
	"log"
)

// create service example implementation.
// The example methods log the requests and return zero values.
type createsrvc struct {
	logger *log.Logger
}

type Comment struct {
	comment string
}


type Blog struct{
	ID uint32
	Name string
	comment []string
}

// NewCreate returns the create service implementation.
func NewCreate(logger *log.Logger) create.Service {
	return &createsrvc{logger}
}

var comments = []string{
	"a",
	"b",
	"c", // comma added
}

type blog_obj []Blog

var blogs = blog_obj {
	{
		ID: 1,
		Name: "abc",
		comment: comments,
	},
}

// Add new blog and return its ID.
func (s *createsrvc) Create(ctx context.Context, p *create.Blog) (res *create.Blog, err error) {
	res = &create.Blog{}
	s.logger.Print("create.create")

	var obj Blog
	obj.ID = *p.ID
	obj.Name = *p.Name
	obj.comment = p.Comments
	blogs = append(blogs, obj)

	newBlog := (&create.Blog{ID: p.ID, Name: p.Name, Comments: p.Comments})
	return newBlog, nil
}


// List all entries
func (s *createsrvc) List(ctx context.Context) (res []*create.Storedblog, err error) {
	s.logger.Print("create.list")

	result := []*create.Storedblog{}
	for _,blog := range blogs{
		item := create.Storedblog{blog.ID, blog.Name, blog.comment}
		result= append(result, &item)
	}
	return result, nil
}

// Remove blog from storage
func (s *createsrvc) Remove(ctx context.Context, p *create.RemovePayload) (err error) {
	s.logger.Print("create.remove")

	for i, singleBlog := range blogs {
		if singleBlog.ID == p.ID {
			blogs = append(blogs[:i], blogs[i+1:]...)
			log.Print("The event with ID has been deleted successfully", singleBlog.ID)
		}
	}
	return
}

// Updating the existing blog
func (s *createsrvc) Update(ctx context.Context, p *create.UpdatePayload) (err error) {
	s.logger.Print("create.update")

	for i, singleBlog := range blogs {
		if singleBlog.ID == *p.ID {
			singleBlog.Name = p.Name
			singleBlog.comment = p.Comments
			blogs = append(blogs[:i], singleBlog)
		}
	}

	return
}
