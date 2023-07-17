package router

import (
	//user defined package
	"echo/authentication"
	"echo/handler"

	//third party package

	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()
    // Router
	e.POST("/signup", handler.Signup)//common                                                  
	e.POST("/login", handler.Login)//common                                                     
	e.POST("/jobposting", handler.Jobposting, authentication.AuthMiddleware)//admin
	e.GET("/getjobposts", handler.GetJobPostingDetails,authentication.AuthMiddleware)//common                                
	e.GET("getjobpostbyid/:id", handler.GetJobPostingByID,authentication.AuthMiddleware)//common                              
	e.GET("getjobpostbycompanyname/:companyname", handler.GetJobPostingByCompany,authentication.AuthMiddleware)//common
	e.PUT("updatejobpostbyid/:id", handler.UpdateJob, authentication.AuthMiddleware)//admin
	e.DELETE("deletejobpostbyid/:id", handler.DeleteJob, authentication.AuthMiddleware)//admin
	e.POST("/user/comments", handler.UserComments,authentication.AuthMiddleware)//user                               
	e.GET("/getallcomments", handler.GetUserComments,authentication.AuthMiddleware)//common                                  
	e.GET("/getcommentsbyid/:id", handler.GetCommentByID,authentication.AuthMiddleware)//common
	e.PUT("updatecommentbyid/:id", handler.UpdateComment,authentication.AuthMiddleware)//user
	e.DELETE("deletecommentbyid/:id", handler.DeleteCommentById,authentication.AuthMiddleware)//user
	e.Logger.Fatal(e.Start(":8000"))
}
