package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("blog", func() {
	Title("Blog Service")
	Description("Service to perform CRUD operations using goa")
	Meta("swagger:example", "false")
	Server("blog", func() {
		Services("oauth", "blog", "swagger")
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
	cors.Origin("/.*localhost/", func() {
		cors.Headers("X-Shared-Secret")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})
})
