package builder

import (
	"github.com/labstack/echo/v4"
	"test/internal/pkg/http_server/handlers/time"
	"test/internal/pkg/http_server/middleware/admin_check"
)

type IServerBuilder interface {
	MakeRouter()
	MakeAPI()
	MakeServer()
}

type ServerBuilder struct {
	Product *Server
}

func (sb *ServerBuilder) MakeRouter() {
	sb.Product.Router = echo.New()
}

func (sb *ServerBuilder) MakeAPI() {
	sb.Product.Router.GET("/api/v0/time", time.New, admin_check.New)
}

func (sb *ServerBuilder) MakeServer() {
	sb.Product.Router.Logger.Fatal(sb.Product.Router.Start(":1323"))
}
