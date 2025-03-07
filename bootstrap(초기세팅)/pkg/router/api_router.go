package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type ApiRouter struct{} // 빈 struct를 만들고 메서드를 만든다라..

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New()) // 단위 시간당 요청 limit 걸기
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
