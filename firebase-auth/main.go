// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	gofiberfirebaseauth "github.com/sacsand/gofiber-firebaseauth" // 누가 만들어놨네요 미들웨어
	"google.golang.org/api/option"                                // 구글 api 옵션? 여기다가 ACCOUNT 정보를 넣네요
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var fireApp *firebase.App

func main() {
	// Fiber instance
	app := fiber.New()

	// Get google service account credentials
	serviceAccount, fileExi := os.LookupEnv("GOOGLE_SERVICE_ACCOUNT")

	if !fileExi {
		log.Fatalf("Please provide valid firebbase auth credential json!")
	}

	// Initialize the firebase app.
	opt := option.WithCredentialsFile(serviceAccount)
	fireApp, _ = firebase.NewApp(context.Background(), nil, opt)

	// Unauthenticated routes
	app.Get("/salaanthe", salanthe)

	// Initialize the middleware with config. See https://github.com/sacsand/gofiber-firebaseauth for more configuration options.
	app.Use(gofiberfirebaseauth.New(gofiberfirebaseauth.Config{
		// Firebase Authentication App Object
		// Mandatory
		FirebaseApp: fireApp,
		// Ignore urls array.
		// Optional. These url will ignore by middleware
		IgnoreUrls: []string{"GET::/salut", "POST::/ciao"},
	}))

	// app.Get("/login/:uid", getToken)

	// Authenticaed Routes.
	app.Get("/hello", hello)
	app.Get("/salut", salut) // Ignore the auth by IgnoreUrls config
	app.Post("/ciao", ciao)  // Ignore the auth by IgnoreUrls config
	app.Get("/ayubowan", ayubowan)

	// Start server.
	log.Fatal(app.Listen(":3001"))
}

/**
*
* Controllers
*
 */

// 회원가입, 로그인 로직은 클라이언트 단에서 만들어야함 ㅇㅇ

// func getToken(c *fiber.Ctx) error {
// 	uid := c.Params("uid")
// 	// Firebase 앱에서 Auth 클라이언트 가져오기
// 	auth, err := fireApp.Auth(context.Background())
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Auth 클라이언트 생성 실패",
// 		})
// 	}

// 	// 커스텀 토큰 생성 (두 번째 인자는 uid)
// 	token, err := auth.CustomToken(context.Background(), uid)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "토큰 생성 실패",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"token": token,
// 	})
// }

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func salut(c *fiber.Ctx) error {
	return c.SendString("Salut, World 👋!")
}

func ciao(c *fiber.Ctx) error {
	return c.SendString("Ciao, World 👋! ")
}

func ayubowan(c *fiber.Ctx) error {
	// Get authenticated user from context.
	claims := c.Locals("user")
	fmt.Println(claims)
	return c.SendString("Ayubowan👋! ")
}

func salanthe(c *fiber.Ctx) error {
	return c.SendString("Salanthe👋! ")
}
