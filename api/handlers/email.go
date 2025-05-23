package handlers

import (
	"gosecops/internal/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailAttackInput struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailAnalyzeInput struct {
	Domain string `json:"domain"`
}

// POST /api/email/attack
func HandleEmailAttack(c *gin.Context) {
	var req EmailAttackInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := email.SendSpoofedEmail(email.EmailAttackRequest{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Spoofed email sent", "to": req.To})
}

// POST /api/email/analyze
func HandleEmailAnalyze(c *gin.Context) {
	var req EmailAnalyzeInput
	if err := c.ShouldBindJSON(&req); err != nil || req.Domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid domain"})
		return
	}

	result := email.AnalyzeEmailSecurity(req.Domain)
	c.JSON(http.StatusOK, result)
}
