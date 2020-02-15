package k8s

import "github.com/kataras/iris/v12"

func RegisterForK8s(c iris.Party) {
	c.Get("/users", allUsersHandler)
	c.Get("/user/{id:int}", singUserHandler)
	c.Post("/user", createUserHandler)
	c.Delete("/user/{id:int}", deleteUserHandler)
	c.Get("/topics", allTopicHandler)
	c.Post("/topic", createTopicHandler)
}
