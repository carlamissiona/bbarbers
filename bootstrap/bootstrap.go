package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"   
    "carlamissiona/golang-barbers/pkg/router"
    "log"  
	_ "github.com/lib/pq"
  // "net/http"
)   

 
func NewApplication() *fiber.App {
  
  engine := html.New("./templates", ".html")
  
 //  fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static", fs)

   
  app := fiber.New(fiber.Config{Views: engine})
  app.Static("/", "./assets")
  log.Println("assets") 
  app.Use(recover.New()) 
  app.Use(logger.New())
  
  app.Get("/dashboard", monitor.New())
  var r router.Router = nil
  r = router.NewHttpRouter()
  r.InstallRouter(app)
	return app
  
  
}

// chat - click button type - send - send to chatroom 
// display chatroom
// usermgmt