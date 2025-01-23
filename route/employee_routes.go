package route

import (
	"github.com/gin-gonic/gin"
	"iHR/db"
	"iHR/db/repositories"
	. "iHR/handler/employee"
)

func RegisterEmployeeRoutes(r *gin.Engine) {
	employeeRoutes := r.Group("/employee")
	{
		employeeRepo := repositories.NewEmployeeRepo(db.DB)
		employeeHandler := NewEmployeeHandler(employeeRepo)
		employeeRoutes.POST("/", employeeHandler.CreateEmployee)
		employeeRoutes.GET("/", employeeHandler.GetAllEmployees)
		employeeRoutes.GET("/:id", employeeHandler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}
}
