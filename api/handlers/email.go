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

// HandleEmailAttack godoc
// @Summary Send spoofed test email
// @Description Sends an email using fake "From" field (Maildev/Mailhog only)
// @Tags Email
// @Accept json
// @Produce json
// @Param input body EmailAttackInput true "Spoofed Email Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /email/attack [post]
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

// HandleEmailAnalyze godoc
// @Summary Analyze SPF/DKIM/DMARC records
// @Description Checks email security configs for a domain
// @Tags Email
// @Accept json
// @Produce json
// @Param input body EmailAnalyzeInput true "Domain to analyze"
// @Success 200 {object} email.EmailAnalysis
// @Failure 400 {object} map[string]string
// @Router /email/analyze [post]
func HandleEmailAnalyze(c *gin.Context) {
	var req EmailAnalyzeInput
	if err := c.ShouldBindJSON(&req); err != nil || req.Domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid domain"})
		return
	}

	result := email.AnalyzeEmailSecurity(req.Domain)
	c.JSON(http.StatusOK, result)
}
