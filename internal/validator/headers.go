package validator

import (
	"net/mail"
	"strings"
)

type HeaderAnalysis struct {
	From         string
	ReceivedBy   []string
	ReplyTo      string
	IsSuspicious bool
}

func AnalyzeHeaders(raw string) HeaderAnalysis {
	msg, _ := mail.ReadMessage(strings.NewReader(raw))
	headers := msg.Header

	received := headers["Received"]
	replyTo := headers.Get("Reply-To")
	from := headers.Get("From")

	suspicious := strings.Contains(strings.ToLower(from), "support") && !strings.Contains(replyTo, "company.com")

	return HeaderAnalysis{
		From:         from,
		ReplyTo:      replyTo,
		ReceivedBy:   received,
		IsSuspicious: suspicious,
	}
}
