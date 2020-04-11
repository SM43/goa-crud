package design

import . "goa.design/goa/v3/dsl"

var StoredBlogs = ResultType("application/vnd.cellar.stored-blog", func() {
	Description("A Storedblog describes a blog retrieved by the storage service.")
	Reference(Blog)
	TypeName("Storedblog")

	Attributes(func() {
		Attribute("id", UInt32, "ID is the unique id of the blog.")

		Attribute("name", String, "Name of person", func() {
			MaxLength(100)
		})

		Attribute("comments", ArrayOf(comments), "Comments", func() {
			MaxLength(100)
		})

	})

	Required("id", "name")
})

var Blog = Type("Blog", func() {
	Description("Blog with id and name of a person")

	Attribute("id", UInt32, "ID of a person")

	Attribute("name", String, "Name of person", func() {
		MaxLength(100)
	})
	Attribute("comments", ArrayOf(comments), "Comments", func() {
		MaxLength(100)
	})

})

var new_comment = Type("new_comment", func() {
	Description("New comment to be added to an existing blog")

	Attribute("id", UInt32, "Id of blog")

	Attribute("comments", comments, "Comment added to an existing blog")

})

var comments = Type("Comments", func() {

	Description("Id and comments")

	Attribute("id", UInt32, "ID of a comment")

	Attribute("comments", String, "Comment for the blog")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a blog that does not exist.")
	Field(2, "id", UInt32, "ID of missing blog")
	Required("id")
})
