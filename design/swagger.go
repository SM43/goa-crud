package design

import . "goa.design/goa/v3/dsl"

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")

	Method("sm1", func() {
		Description("Add a new blog")
		Result(Resource, func() { View("default") })
		HTTP(func() {
			POST("/sma")
			Response(StatusCreated)
		})
	})

	Method("sm2", func() {
		Description("Add a new blog")
		Result(Resource, func() { View("extended") })
		HTTP(func() {
			POST("/smb")
			Response(StatusCreated)
		})
	})

	Method("sm3", func() {
		Description("Add a new blog")
		Result(Resource, func() { View("desc") })
		HTTP(func() {
			POST("/smc")
			Response(StatusCreated)
		})								
	})
})
