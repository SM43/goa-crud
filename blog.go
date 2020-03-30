package blogapi

import (
	"context"
	blog "crud/gen/blog"
	"log"
)

// blog service example implementation.
// The example methods log the requests and return zero values.
type blogsrvc struct {
	logger *log.Logger
}

type Blog struct{
	ID uint32
	Name string
	comment []string
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

// NewBlog returns the blog service implementation.
func NewBlog(logger *log.Logger) blog.Service {
	return &blogsrvc{logger}
}

// Add new blog and return its ID.
func (s *blogsrvc) Create(ctx context.Context, p *blog.Blog) (res *blog.Blog, err error) {
	res = &blog.Blog{}
	s.logger.Print("blog.create")

	var obj Blog
	obj.ID = *p.ID
	obj.Name = *p.Name
	obj.comment = p.Comments
	blogs = append(blogs, obj)

	newBlog := (&blog.Blog{ID: p.ID, Name: p.Name, Comments: p.Comments})
	return newBlog, nil
}

// List all entries
func (s *blogsrvc) List(ctx context.Context) (res []*blog.Storedblog, err error) {
	s.logger.Print("blog.list")

	result := []*blog.Storedblog{}
	for _,all_blogs := range blogs{
		item := blog.Storedblog{all_blogs.ID, all_blogs.Name, all_blogs.comment}
		result= append(result, &item)
	}
	return result, nil
}

// Remove blog from storage
func (s *blogsrvc) Remove(ctx context.Context, p *blog.RemovePayload) (err error) {
	s.logger.Print("blog.remove")

	for i, singleBlog := range blogs {
		if singleBlog.ID == p.ID {
			blogs = append(blogs[:i], blogs[i+1:]...)
			log.Print("The event with ID has been deleted successfully", singleBlog.ID)
		}
	}
	return
}

// Updating the existing blog
func (s *blogsrvc) Update(ctx context.Context, p *blog.UpdatePayload) (err error) {
	s.logger.Print("blog.update")

	for i, singleBlog := range blogs {
		if singleBlog.ID == *p.ID {
			singleBlog.Name = p.Name
			singleBlog.comment = p.Comments
			blogs = append(blogs[:i], singleBlog)
		}
	}
	return
}
