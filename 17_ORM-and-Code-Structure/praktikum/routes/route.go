package routes

import (
	"github.com/labstack/echo/v4"

	controller "praktikum/controllers"
)

func New() *echo.Echo {
	e := echo.New()


	// Users Route Group
	users := e.Group("/users")
	users.GET("", controller.GetUsers)
	users.POST("", controller.CreateUser)

	users.GET("/:id", controller.GetUser)
	users.PUT("/:id", controller.UpdateUser)
	users.DELETE("/:id", controller.DeleteUser)

	// Books Route Group
	books := e.Group("/books")
	books.GET("", controller.GetBooks)
	books.POST("", controller.CreateBook)

	books.GET("/:id", controller.GetBook)
	books.PUT("/:id", controller.UpdateBook)
	books.DELETE("/:id", controller.DeleteBook)

	// Blogs Route Group
	blogs := e.Group("/blogs")
	blogs.GET("", controller.GetBlogs)
	blogs.POST("", controller.CreateBlog)

	blogs.GET("/:id", controller.GetBlog)
	blogs.PUT("/:id", controller.UpdateBlog)
	blogs.DELETE("/:id", controller.DeleteBlog)

	return e
}