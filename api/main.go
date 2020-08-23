package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()
	app.Use(middleware.Logger(middleware.LoggerConfig{
    Format:     "${time} ${method} ${path}",
    TimeFormat: "15:04:05",
    TimeZone:   "Asia/Chongqing",
    Next: func (c *fiber.Ctx) bool {
        var isUserAdmin bool
        // Your logic here
        return isUserAdmin
    }
}))

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(6000)
}
