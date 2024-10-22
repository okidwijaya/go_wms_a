package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okidwijaya/go_wms_a/models"
)

type ReceiverRequest struct {
	Header models.ReceiverHeader   `json:"header"`
	Detail []models.ReceiverDetail `json:"detail"`
}

func ReceiverController(c *gin.Context) {
	var request ReceiverRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	request.Header.TrxInDate = time.Now() // Set the current time

	trxID, err := models.TransactionReceiveHeader(request.Header)
	if err != nil {
		log.Printf("Failed to add receiver header: %v", err) // Log the specific error
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add receiver header"})
		return
	}

	for _, detail := range request.Detail {
		detail.TrxInIDF = int(trxID) // Associate the transaction ID
		if err := models.TransactionReceiveDetail(detail); err != nil {
			log.Printf("Failed to add product detail: %v", err) // Log the specific error
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add product detail"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Receiver transaction added successfully", "trxID": trxID})
}
