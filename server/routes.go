package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/QuizMaker/db"
)

// APIRoutes holds the routes of API
type APIRoutes struct {
	dbs    *db.DB
	engine *gin.Engine
}

// NewAPIRoutes is the constructor for APIRoutes
func NewAPIRoutes(dbs *db.DB) *APIRoutes {
	return &APIRoutes{
		dbs:    dbs,
		engine: gin.New(),
	}
}

// GetEngine returns the gin engine
func (routes *APIRoutes) GetEngine() *gin.Engine {
	return routes.engine
}

// SetRoutes set the endpoints of API
func (routes *APIRoutes) SetRoutes() {

	routes.engine.GET(
		"/",
		func() gin.HandlerFunc {
			return func(c *gin.Context) {
				c.JSON(200, map[string]interface{}{"message": "Hello World!!"})
			}
		}(),
	)

}
