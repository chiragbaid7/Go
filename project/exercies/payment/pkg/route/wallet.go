package route

import "github.com/labstack/echo"

func RegisterRoutes() {
	e := echo.New()
	e.GET("/health-check", func(c echo.Context) error {
		return c.String(200, "Working fine")
	})

	// context represent represents the context of the current HTTP request req and res objects and methods
	e.POST("/setup", func(c echo.Context) error {
		data := c.Request().Body
	})
	e.Logger.Fatal(e.Start(":3000"))
}
