package controllers

import (
  "github.com/gofiber/fiber/v2"
  "log"
  )

func RenderHome(c *fiber.Ctx) error {
  log.Println("Higshsss!") 
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	},"layouts/htm")
}      

func RenderAbout(c *fiber.Ctx) error {
  log.Println("Higshsss!") 
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	},"layouts/htm")
}       

func RenderContact(c *fiber.Ctx) error {
  log.Println("High")  
	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	},"layouts/htm")
}     
 