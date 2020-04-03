package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("blog", func() {
	Title("Hello Service")
	Description("Service to print hello")
    Server("blog", func() {
        Host("localhost", func() {
            URI("http://localhost:8000")
            URI("grpc://localhost:8080")
        })
    })
})

var _ = Service("blog", func() {
	Description("The blog service gives blog details.")

	Method("create", func() {
		Description("Add new blog and return its ID.")
		Payload(Blog)
		Result(Blog)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})


	Method("list", func() {
		Description("List all entries")
		Result(ArrayOf(StoredBlogs))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})


	Method("remove", func() {
		Description("Remove blog from storage")
		Payload(func() {
			Field(1, "id", UInt32, "ID of blog to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Blog not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})

	Method("update", func() {
		Description("Updating the existing blog")
		Payload(func() {
			Field(1, "id", UInt32, "ID of blog to be updated")
			Field(2,"name", String, "Details of blog to be updated")
			Field(3, "comments", ArrayOf(comments), "Comments to be updated")
			Required("name", "comments")
		})
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusNoContent)
		})

	})

	Method("add", func() {
		Description("Add new blog and return its ID.")
		Payload(new_comment)
		Result(new_comment)
		HTTP(func() {
			POST("/{id}/comments/")
			Response(StatusCreated)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})