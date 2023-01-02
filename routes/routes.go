package routes

import (
	invoice "main/controlers/invoice"
	user "main/controlers/user"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(g *gin.RouterGroup) {

	g.GET("/user", user.GetAll())
	g.GET("/user/:id", user.GetUser())
	g.POST("/user", user.AddUser())
	g.PUT("/user/:id", user.EditUser())
	g.DELETE("/user/:id", user.DeleteUser())

	g.GET("/invoice", invoice.GetAll())
	g.GET("/paid", invoice.GetPaid())
	g.GET("/unpaid", invoice.GetUnpaid())
	g.GET("/invoice/:id", invoice.GetInvoice())
	g.GET("/download/:id", invoice.Download())
	g.POST("/add", invoice.Add())
	g.PUT("/edit/:id", invoice.Edit())
	g.DELETE("/delete/:id", invoice.Delete())
}
