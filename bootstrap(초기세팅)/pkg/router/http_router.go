package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/kooroshh/fiber-boostrap/app/controllers"
)

type HttpRouter struct{}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New(), csrf.New())
	group.Get("/", controllers.RenderHello)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}

/* 여기서 http router는 cors하고 csrf를 챙김
-> 브라우저로부터 호출될거고, 쿠키를 쓸거임
*/
