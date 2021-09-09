package routes

import (
	"../controllers"
	"../middlewares"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	// маршрут для обновления данных авторизированного пользователя
	app.Put("/api/users/info", controllers.UpdateInfo)
	// маршрут для обновления пароля авторизированного пользователя
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermisions)

	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/employees", controllers.AllEmployees)
	app.Post("/api/employees", controllers.CreateEmployee)
	app.Get("/api/employees/:id", controllers.GetEmployee)
	app.Put("/api/employees/:id", controllers.UpdateEmployee)
	app.Delete("/api/employees/:id", controllers.DeleteEmployee)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")

	app.Get("/api/agents", controllers.AllAgents)
	app.Get("/api/agents-list", controllers.AllAgentsList)

	app.Post("/api/export-agents", controllers.ExportAgents)
	app.Post("/api/agents", controllers.CreateAgent)
	app.Get("/api/agents/:id", controllers.GetAgent)
	app.Put("/api/agents/:id", controllers.UpdateAgent)
	app.Delete("/api/agents/:id", controllers.DeleteAgent)

	app.Post("/api/upload/agent", controllers.AgentPhotoUpload)
	app.Static("/api/uploads/agents", "./uploads/agents")


	app.Get("/api/organizations", controllers.AllOrganizations)

	app.Get("/api/customers", controllers.AllCustomers)
	app.Get("/api/customers-list", controllers.AllCustomerList)

	app.Post("/api/export-customers", controllers.ExportCustomers)
	app.Post("/api/customers", controllers.CreateCustomer)
	app.Get("/api/customers/:id", controllers.GetCustomer)
	app.Put("/api/customers/:id", controllers.UpdateCustomer)
	app.Delete("/api/customers/:id", controllers.DeleteCustomer)

	app.Get("/api/contracts", controllers.AllContracts)
	app.Post("/api/contracts", controllers.CreateContract)
	app.Get("/api/contracts/:id", controllers.GetContract)
	app.Put("/api/contracts/:id", controllers.UpdateContract)
	app.Delete("/api/contracts/:id", controllers.DeleteContract)


	app.Get("/api/agreements", controllers.AllAgreements)
	app.Delete("/api/agreements/:id", controllers.DeleteAgreement)
	app.Post("/api/agreements", controllers.CreateAgreement)
	app.Get("/api/agreements/:id", controllers.GetAgreement)
	app.Put("/api/agreements/:id", controllers.UpdateAgreement)




}
