package router
 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"carlamissiona/golang-barbers/app/controllers"
)
type Router interface {
	InstallRouter(app *fiber.App)
}

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	web := app.Group("", cors.New(), csrf.New())
  web.Get("/", controllers.RenderHome)
  web.Get("/contact-us", controllers.RenderContact)
  
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
