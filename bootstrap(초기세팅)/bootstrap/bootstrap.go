package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"  // 요청에 대해서 로깅을 해줌 (메서드, 경로, 상태코드, 처리 시간...)
	"github.com/gofiber/fiber/v2/middleware/monitor" // 앱의 성능, 상태를 모니터링할 수 있는 대쉬보드 제공
	"github.com/gofiber/fiber/v2/middleware/recover" // panic시 서버가 중단되지 않도록 함.
	"github.com/gofiber/template/html/v2"            // html template engine
	"github.com/kooroshh/fiber-boostrap/pkg/database"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
	"github.com/kooroshh/fiber-boostrap/pkg/router"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}
