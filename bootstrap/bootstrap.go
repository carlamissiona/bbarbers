package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"  
  _ "carlamissiona/golang-barbers/pkg/database"
  "carlamissiona/golang-barbers/pkg/router"
  "log" 
) 
 
func NewApplication() *fiber.App {
  
  // _ = database.SetupDatabase()
  log.Println("db_instance")
	engine := html.New("./templates", ".html")
 
  
  app := fiber.New(fiber.Config{Views: engine})
  app.Static("/", "../assets")
  log.Println("assets") 
  app.Use(recover.New()) 
	app.Use(logger.New())
  
	app.Get("/dashboard", monitor.New())
  var r router.Router = nil
  r = router.NewHttpRouter()
  r.InstallRouter(app)
	return app
  
  
}

