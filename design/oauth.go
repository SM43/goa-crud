package design

import . "goa.design/goa/v3/dsl"

var _ = Service("oauth", func() {
	Description("The oauth service authorise user to access other APIs")

	Error("internal_error", ErrorResult, "Unable to process request")

	Method("oauth", func() {
		Description("Github authentication to post a new blog")
		Payload(func() {
			Attribute("token", String, "Access github token")
		})
		Result(String)
		HTTP(func() {
			POST("/oauth/redirect")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
		})

	})
})
