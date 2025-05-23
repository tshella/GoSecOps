# 🛡️ GoSecOps – Penetration Testing & Cloud Security Toolkit

**GoSecOps** is a modular, containerized security toolkit built with **GoLang**. It supports both **CLI** and **REST API** interfaces for red-team attack simulation (e.g., port scanning, spoofed emails) and blue-team validation (e.g., SPF/DKIM/DMARC analysis, IAM audits, DNS misconfigs, and S3 exposure).

> **Cloud Security Architect:** Manaka Anthony Raphasha

---

## 📚 Table of Contents

- [🚀 Features](#-features)
- [🧱 Project Structure](#-project-structure)
- [💻 Getting Started](#-getting-started)
- [🔧 Build & Run](#-build--run)
- [📬 API Endpoints](#-api-endpoints)
- [📚 Swagger UI (Interactive API Docs)](#-swagger-ui-interactive-api-docs)
- [📬 MailTest (MailDev) – Safe Email Testing](#-mailtest-maildev--safe-email-testing)
- [🔐 Security & Ethics Notice](#-security--ethics-notice)
- [🛠️ Libraries Used](#️-libraries-used)
- [📦 Roadmap](#-roadmap)
- [👨‍💻 Dev Commands](#-dev-commands)
- [🤝 Contributing](#-contributing)
- [📜 License](#-license)

---

## 🚀 Features

| Feature              | CLI Command             | REST API Endpoint         | Status   |
|----------------------|-------------------------|----------------------------|----------|
| Port Scanning        | `scan`                  | `/api/scan/port`           | ✅ Done  |
| Email Spoofing       | `email attack`          | `/api/email/attack`        | ✅ Done  |
| Email Analysis       | `email analyze`         | `/api/email/analyze`       | ✅ Done  |
| IAM Policy Audit     | `iam check`             | `/api/cloud/iam`           | ✅ Done  |
| S3 Bucket Audit      | `s3 audit`              | `/api/cloud/s3`            | ✅ Done  |
| DNS Cloud Scanner    | `dns cloud`             | `/api/cloud/dns`           | ✅ Done  |
| Web UI (SvelteKit)   | ❌ In progress           | ✅ Backend Ready            | 🧪 Alpha |
| Logging              | JSON/stdout (planned)   | —                          | 🟡 Soon  |
| Report Export (JSON) | —                       | —                          | 🟡 Soon  |

---

## 🧱 Project Structure

gosecops/
├── api/ # REST API layer
│ ├── handlers/ # Endpoint logic
│ └── routes.go # API route definitions
├── cmd/ # CLI command handlers
│ ├── root.go
│ ├── scan.go
│ ├── email.go
│ ├── iam.go
│ ├── s3.go
│ └── dns.go
├── internal/ # Core security logic
│ ├── scanner/ # Port scanning logic
│ ├── email/ # Spoof + SPF/DKIM/DMARC logic
│ ├── cloud/ # IAM, S3, DNS audits
│ ├── utils/ # Shared utilities
│ └── validator/ # (Planned) security header analyzer
├── web/ # Optional SvelteKit frontend
├── Dockerfile # Multi-stage backend build
├── docker-compose.yml # For API + mail testing
├── prometheus.yml # Prometheus config
├── main.go # Entry point
├── go.mod # Go module metadata
└── README.md # This documentation


---

## 💻 Getting Started

### ✅ Prerequisites

- Go 1.20+
- Docker + Docker Compose
- Node.js (for optional SvelteKit frontend)

---

## 🔧 Build & Run

### 🧪 CLI Usage

```bash
# Port Scan
go run main.go scan --target 192.168.0.1

# Send spoofed email
go run main.go email attack --from spoof@microsoft.com --to victim@yourdomain.test

# Analyze SPF/DKIM/DMARC
go run main.go email analyze --domain yourdomain.test

# IAM Policy Audit
go run main.go iam check --profile default

# S3 Bucket Audit
go run main.go s3 audit --bucket my-bucket --profile default

# DNS Misconfig Scan
go run main.go dns cloud --domain example.com --subdomains www,api,staging

🌐 API Mode

go run main.go

Access:

    http://localhost:8181/api

    http://localhost:8181/swagger/index.html

🐳 Docker Mode

docker compose up --build

📬 API Endpoints
Method	Endpoint	Description
POST	/api/scan/port	Run a TCP port scan
POST	/api/email/attack	Send a spoofed test email
POST	/api/email/analyze	Analyze SPF/DKIM/DMARC
POST	/api/cloud/iam	Audit AWS IAM policies
POST	/api/cloud/s3	Detect public exposure of S3 buckets
POST	/api/cloud/dns	Scan for DNS CNAME misconfigurations
📚 Swagger UI (Interactive API Docs)

Once the API server is running, access the Swagger UI at:

➡️ http://localhost:8181/swagger/index.html

This provides:

    📖 Full documentation of each API route

    🧪 Built-in testing interface

    📂 Schema definitions for request and response bodies

    You can also export OpenAPI JSON/YAML from this UI.

📬 MailTest (MailDev) – Safe Email Testing

MailTest uses MailDev, a fake SMTP server with a web inbox. It captures spoofed emails locally, ensuring safe red-team simulations.
🔧 Docker Integration

mailtest:
  image: maildev/maildev
  ports:
    - "1080:1080"  # Web UI
    - "1025:1025"  # SMTP test port

✉️ Example Usage

go run main.go email attack --from admin@paypal.com --to victim@example.com

Then open http://localhost:1080 to view the email.
Feature	Benefit
SMTP inbox	No email is actually delivered
Web UI	Real-time message viewing
Safe sandbox	Great for testing phishing flows
🔐 Security & Ethics Notice

⚠️ GoSecOps is intended for educational or authorized security testing only.

    Never test domains or IPs you do not own or have written permission for.

    Spoofed emails are routed to sandboxed environments (MailDev).

    Future versions will log audit trails and offer RBAC.

🛠️ Libraries Used
Purpose	Library
CLI Framework	github.com/spf13/cobra
HTTP API Server	github.com/gin-gonic/gin
Email Spoofing	net/smtp
AWS SDK	github.com/aws/aws-sdk-go-v2
Swagger UI	github.com/swaggo/gin-swagger
DNS Lookups	net.LookupTXT, net.LookupIP
Port Scanning	net.DialTimeout
📦 Roadmap

    ✅ Cloud IAM, S3, DNS misconfig modules

    ✅ Swagger/OpenAPI docs

    ✅ SvelteKit frontend

    ⏳ Security header analyzer

    ⏳ JSON/CSV report export

    ⏳ WebSocket log streams

    ⏳ JWT/OAuth-based API authentication

👨‍💻 Dev Commands

# Run CLI audit
go run main.go s3 audit --bucket my-bucket --profile default

# Start API server
go run main.go

# Docker dev environment
docker compose up --build

🤝 Contributing

We welcome pull requests, issue reports, and security reviews.
Please include unit tests for new features or modules.
📜 License

MIT License – see LICENSE
