package controllers

import (
  "github.com/gofiber/fiber/v2"
  "log"
  )

func RenderHello(c *fiber.Ctx) error {
  log.Println("High") 
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}  
