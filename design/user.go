package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Description("The user service gives user details.")

	Error("db_error", ErrorResult, "Unable to process db request")
	Error("invalid-token", ErrorResult, "User token not valid")

	//Method to post new user
	Method("create", func() {
		Description("Add a new blog")
		Payload(func() {
			Attribute("user", User, "Adding a new user")
			Attribute("auth", String, "Access github token")
			Required("auth", "user")
		})
		HTTP(func() {
			POST("/user")
			Header("auth:Authorization") // JWT token passed in "X-Authorization" header
			Response(StatusCreated)
			Response("db_error", StatusInternalServerError)
			Response("invalid-token", StatusUnauthorized)
		})
	})

	//Method to get all existing users
	Method("list", func() {
		Description("List all the users")
		Result(ArrayOf(User))
		HTTP(func() {
			GET("/users")
			Response(StatusOK)
			Response("db_error", StatusInternalServerError)
		})
	})
})
