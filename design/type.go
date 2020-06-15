package design

import . "goa.design/goa/v3/dsl"

// var Resources = CollectionOf(Resource, func() {
// 	View("extended")
// })

// Resource defines return type for all apis
var Resource = ResultType("application/vnd.hub.resource", func() {
	TypeName("Resource")
	Attributes(func() {

		Attribute("id", UInt, "ID is the unique id of the resource", func() {
			Example("id", 1)
		})
		Attribute("name", String, "Name of the resource", func() {
			Example("name", "golang")
		})

		Attribute("type", String, "Type of resource", func() {
			Example("type", "task")
		})

		Attribute("rating", UInt, "Rating of resource", func() {
			Example("rating", 5)
		})

	})
	View("extended", func() {
		Attribute("id")
		Attribute("name")
	})

	View("desc", func() {
		Attribute("id")
		Attribute("name")
		Attribute("rating")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("type")
		Attribute("rating")
	})
	Required("id", "name", "type", "rating")

})

var User = ResultType("application/vnd.goa-crud.stored-user", func() {
	Description("A User describes a user retrieved by the storage service.")
	TypeName("User")
	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the user")
		Attribute("name", String, "Name of user")
		Attribute("age", UInt, "Age of user")
		Attribute("class", String, "Class of user")
		Required("name", "age", "class")
	})

})

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
