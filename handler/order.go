package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/api/utils"
)

func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order ID not found",
		})
	}

	id64 := utils.ConvertToUint64(id)

	order, err := h.Repo.FindDelivery(uint(id64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed To Retrieve Order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"order":   order,
	})
	return

}