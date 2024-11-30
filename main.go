package main

import (
	"go-crud-1/config"
	"go-crud-1/controllers"
	"go-crud-1/models"
	component "go-crud-1/views/components"
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// templ renderer
func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func main() {
	app := fiber.New()
	config.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		users, err := controllers.FetchUsers()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return render(c, component.UserList(users))
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := controllers.FetchUsers()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return render(c, component.UserList(users))
	})

	app.Get("/post", func(c *fiber.Ctx) error {
		return render(c, component.PostUser())
	})

	app.Get("/edit/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User
		err := config.DB.QueryRow("SELECT id, name, email, gender FROM users WHERE id = ?", id).Scan(
			&user.Id, &user.Name, &user.Email, &user.Gender,
		)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return render(c, component.EditUser(user))
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		gender := c.FormValue("gender-post")
		log.Println(name)
		log.Println(email)
		log.Println(gender)
		if name == "" || email == "" || gender == "" {
			log.Println("fileds must not empty")
			return render(c, component.PostUser())
		}
		_, err := config.DB.Exec("INSERT INTO users (name, email, gender) VALUES (?,?,?)", name, email, gender)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		users, err := controllers.FetchUsers()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return render(c, component.UserList(users))
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		name := c.FormValue("name")
		email := c.FormValue("email")
		gender := c.FormValue("gender-edit")
		var user models.User
		err := config.DB.QueryRow("SELECT id, name, email, gender FROM users WHERE id = ?", id).Scan(
			&user.Id, &user.Name, &user.Email, &user.Gender,
		)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if name == "" || email == "" || gender == "" {
			log.Println("fileds must not empty")
			return render(c, component.EditUser(user))
		}
		result, err := config.DB.Exec("UPDATE users SET name = ?, email = ?, gender = ? WHERE id = ?", name, email, gender, id)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return c.Status(500).SendString("Failed to check affected rows")
		}
		if rowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{"msg": "User not found"})
		}

		// Get user list
		users, err := controllers.FetchUsers()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return render(c, component.UserList(users))
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := controllers.FetchUsers()
		if err != nil {
			c.Status(500).SendString(err.Error())
		}

		if _, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return nil
		// return render(c, component.Users(result))
	})

	app.Static("/css", "./views/css")
	app.Static("/static", "./static")
	log.Println("Server running on port 3001")
	log.Fatal(app.Listen(":3001"))
}
