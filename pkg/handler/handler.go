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
		AllowedOrigins:   []string{"https://*", "http://*"},
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
			lists.POST("/", h.createList)      // add title and description at the todo_lists
			lists.GET("/", h.getAllList)       // get all lists
			lists.GET("/:id", h.getListById)   // getById list
			lists.PUT("/:id", h.updateList)    //update lists
			lists.DELETE("/:id", h.deleteList) // delete lists

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem) // add items
				items.GET("/", h.getAllItems) // get all items
			}
		}
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)   // getById items
			items.PUT("/:id", h.updateItem)    // update items
			items.DELETE("/:id", h.deleteItem) // delete items

		}
	}
	return router
}
