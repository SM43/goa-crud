package design

import . "goa.design/goa/v3/dsl"

var StoredBlog = ResultType("application/vnd.goa-crud.stored-blog", func() {
	Description("A Blog describes a blog retrieved by the storage service.")
	TypeName("StoredBlog")
	Reference(Blog)

	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the blog")
		Attribute("name")
		Attribute("comments", ArrayOf(StoredComment), "Blog with multiple comments")
		Required("id", "name", "comments")
	})

})

var Blog = Type("Blog", func() {
	Description("A Blog describes a blog retrieved by the storage service.")

	Attribute("name", String, "Name of person")
	Attribute("comments", ArrayOf(Comment), "Blog will have multiple comments")

	Required("name", "comments")
})

var StoredComment = Type("StoredComment", func() {
	Description("A blog will have multiple comments")
	Reference(Comment)

	Attribute("id")
	Attribute("comment")

	Required("id", "comment")
})

var Comment = Type("Comment", func() {
	Description("A blog will have multiple comments")

	Attribute("id", UInt, "ID of a comment")
	Attribute("comment", String, "Comment for the blog")

	Required("comment")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a blog that does not exist.")
	Field(2, "id", UInt32, "ID of missing blog")
	Required("id")
})
