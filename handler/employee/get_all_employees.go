package employee

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	employees, err := h.repo.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, employees)
}
