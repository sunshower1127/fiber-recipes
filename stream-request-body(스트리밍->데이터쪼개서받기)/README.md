---
title: Stream Request Body
keywords: [stream, request body]
description: Streaming request bodies.
---

# Stream Request Body

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/stream-request-body) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/stream-request-body)

This project demonstrates how to handle streaming request bodies in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/gofiber/recipes.git
   cd recipes/stream-request-body
   ```

2. Install dependencies:
   ```sh
   go get
   ```

## Running the Application

1. Start the application:
   ```sh
   go run main.go
   ```

## Example

Here is an example of how to handle a streaming request body in Go using Fiber:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "io"
    "log"
)

func main() {
    app := fiber.New()

    app.Post("/upload", func(c *fiber.Ctx) error {
        q
        // Open a file to write the streamed data
        file, err := os.Create("uploaded_file")
        if err != nil {
            return err
        }
        defer file.Close()

        // Stream the request body to the file
        _, err = io.Copy(file, c.BodyStream())
        if err != nil {
            return err
        }

        return c.SendString("File uploaded successfully")
    })

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Go io Package Documentation](https://pkg.go.dev/io)
