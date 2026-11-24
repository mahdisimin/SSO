package mv

import "github.com/labstack/echo/v5"

func ServedBy(InstanceID string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, InstanceID)
			return next(c)
		}
	}
}
