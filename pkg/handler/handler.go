package handler

import (
	"github.com/gaponovalexey/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// start gin
	router.Use(cors.New(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})) // safe cors

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) //return created user id
		auth.POST("/sign-in", h.signIn) //return created user token_hash
	}
	api := router.Group("/api", h.userIdentity) // token JWT
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList) // add title and description at the todo_lists
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := router.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT(":item_id", h.updateItem)
				items.DELETE(":item_id", h.deleteItem)
			}
		}
	}
	return router
}
