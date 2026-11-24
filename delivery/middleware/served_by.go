package mv

import "github.com/labstack/echo/v5"

const HeaderServedBy = "X-Served-By"

func ServedBy(InstanceID string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Response().Header().Set(HeaderServedBy, InstanceID)
			return next(c)
		}
	}
}
