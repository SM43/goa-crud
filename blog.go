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

var blog_store = make([]*blog.Storedblog, 0)

// NewBlog returns the blog service implementation.
func NewBlog(logger *log.Logger) blog.Service {
	return &blogsrvc{logger}
}

// Add new blog and return its ID.
func (s *blogsrvc) Create(ctx context.Context, p *blog.Blog) (res *blog.Blog, err error) {
	res = &blog.Blog{}
	s.logger.Print("blog.create")

	item := blog.Storedblog{*p.ID, *p.Name, p.Comments}
	blog_store = append(blog_store, &item)

	res = (&blog.Blog{ID: p.ID, Name: p.Name, Comments: p.Comments})
	return
}

// List all entries
func (s *blogsrvc) List(ctx context.Context) (res []*blog.Storedblog, err error) {
	s.logger.Print("blog.list")

	return blog_store, nil
}

// Remove blog from storage
func (s *blogsrvc) Remove(ctx context.Context, p *blog.RemovePayload) (err error) {

	s.logger.Print("blog.remove")

	for i, singleBlog := range blog_store {
		if singleBlog.ID == p.ID {
			blog_store = append(blog_store[:i], blog_store[i+1:]...)
			log.Print("The event with ID has been deleted successfully", singleBlog.ID)
		}
	}
	return
}

// Updating the existing blog
func (s *blogsrvc) Update(ctx context.Context, p *blog.UpdatePayload) (err error) {
	s.logger.Print("blog.update")

	for i, singleBlog := range blog_store {
		if singleBlog.ID == *p.ID {
			singleBlog.Name = p.Name
			singleBlog.Comments = p.Comments
			blog_store = append(blog_store[:i], singleBlog)
		}
	}
	return
}

// Add new blog and return its ID.
func (s *blogsrvc) Add(ctx context.Context, p *blog.NewComment) (res *blog.NewComment, err error) {

	res = &blog.NewComment{}
	s.logger.Print("blog.add")

	for i := range blog_store {
		if blog_store[i].ID == *p.ID {
			blog_store[i].Comments = append(blog_store[i].Comments, p.Comments)
		}
	}

	return
}


// Show blog based on the id given
func (s *blogsrvc) Show(ctx context.Context, p *blog.Blog) (res *blog.Blog, err error) {
	res = &blog.Blog{}
	s.logger.Print("blog.show")

	for _,singleBlog := range blog_store {
		if singleBlog.ID == *p.ID {
			res.ID = &singleBlog.ID
			res.Name = &singleBlog.Name
			res.Comments = singleBlog.Comments
		}
	}
	return
}
