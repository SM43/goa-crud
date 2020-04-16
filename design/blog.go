package design

import . "goa.design/goa/v3/dsl"

var _ = Service("blog", func() {
	Description("The blog service gives blog details.")

	Error("db_error", ErrorResult, "Unable to process db request")

	//Method to post new blog
	Method("create", func() {
		Description("Add a new blog")
		Payload(Blog)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
			Response("db_error", StatusInternalServerError)
		})
	})

	//Method to get all existing blogs
	Method("list", func() {
		Description("List all the blogs")
		Result(ArrayOf(StoredBlog))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Response("db_error", StatusInternalServerError)
		})
	})

	//Method to get a particular blog based on id
	Method("show", func() {
		Description("Show blog based on the id given")
		Payload(func() {
			Attribute("id", UInt, "ID of the blog to be fetched")
			Required("id")
		})
		Result(StoredBlog)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("db_error", StatusInternalServerError)
		})
	})

	//Method to remove a particular blog
	Method("remove", func() {
		Description("Delete a blog")
		Payload(func() {
			Attribute("id", UInt, "ID of blog to remove")
			Required("id")
		})
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
			Response("db_error", StatusInternalServerError)
		})
	})

	//Method to add a new comment to an existing blog
	Method("add", func() {
		Description("Add a new comment for a blog")
		Payload(func() {
			Attribute("comments", Comment, "Comment to be added for a blog")
			Attribute("id", UInt, "Id of the blog")
			Required("id", "comments")
		})
		HTTP(func() {
			PATCH("/{id}/comments/")
			Response(StatusOK)
			Response("db_error", StatusInternalServerError)
		})
	})

})
