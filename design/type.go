package design

import . "goa.design/goa/v3/dsl"

var StoredBlogs = ResultType("application/vnd.cellar.stored-blog", func() {
	Description("A Storedblog describes a blog retrieved by the storage service.")
	Reference(Blog)
	TypeName("Storedblog")

	Attributes(func() {
		Attribute("id", UInt32, "ID is the unique id of the blog.", func() {
			Meta("rpc:tag", "8")
		})
	Attribute("name", String, "Name of person", func() {
		MaxLength(100)
		Meta("rpc:tag", "1")
	})
	Attribute("comments", ArrayOf(String), "Comments", func() {
		MaxLength(100)
		Meta("rpc:tag", "4")
	})

	})

	// View("default", func() {
	// 	Attribute("id")
	// 	Attribute("name")

	// })

	// View("tiny", func() {
	// 	Attribute("id")
	// 	Attribute("name")

	// })

	Required("id", "name")
})

var Blog = Type("Blog", func() {
	Description("Blog with id and name of a person")

	Attribute("id", UInt32, "ID of a person", func() {
		Meta("rpc:tag", "2")
	})
	Attribute("name", String, "Name of person", func() {
		MaxLength(100)
		Meta("rpc:tag", "1")
	})
	Attribute("comments", ArrayOf(String), "Comments", func() {
		MaxLength(100)
		Meta("rpc:tag", "4")
	})

})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a blog that does not exist.")
	Field(2, "id", String, "ID of missing blog")
	Required("id")
})

