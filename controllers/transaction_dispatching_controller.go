package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okidwijaya/go_wms_a/models"
)

type DispatchingRequest struct {
	Header models.DispatchingHeader   `json:"header"`
	Detail []models.DispatchingDetail `json:"detail"`
}

func DispatchingController(c *gin.Context) {
	var request DispatchingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	request.Header.TrxOutDate = time.Now()

	trxID, err := models.TransactionDispatchingHeader(request.Header)
	if err != nil {
		log.Printf("Failed to add dispatching header: %v", err) // Log the specific error
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add receiver header"})
		return
	}

	for _, detail := range request.Detail {
		detail.TrxOutIDF = int(trxID)
		if err := models.TransactionDispatchingDetail(detail); err != nil {
			log.Printf("Failed to add product detail: %v", err) // Log the specific error
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add product detail"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dispatching transaction added successfully", "trxID": trxID})
}
