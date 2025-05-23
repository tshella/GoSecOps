# 🛡️ GoSecOps – Penetration Testing & Cloud Security Toolkit

**GoSecOps** is a modular, containerized security tool built in **GoLang**. It offers both **CLI** and **REST API** interfaces for red-team attack simulations (e.g., port scanning, spoofed emails) and blue-team validation (e.g., SPF/DKIM/DMARC checks, IAM audits, DNS misconfigs, S3 exposure).

> **Cloud Security Architect:** Manaka Anthony Raphasha

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
├── api/
│ ├── handlers/ # REST handlers
│ └── routes.go # API endpoint mapping
├── cmd/ # CLI commands
│ ├── root.go
│ ├── scan.go
│ ├── email.go
│ ├── iam.go
│ ├── s3.go
│ └── dns.go
├── internal/ # Core modules
│ ├── scanner/ # Port scanner logic
│ ├── email/ # Email spoof + analysis
│ ├── cloud/ # IAM, S3, DNS misconfig tools
│ ├── utils/ # Common helpers
│ └── validator/ # (Planned) header checks
├── web/ # Optional SvelteKit frontend
├── Dockerfile
├── docker-compose.yml
├── main.go
├── go.mod
└── README.md


---

## 💻 Getting Started

### Prerequisites

- Go 1.20+
- Docker + Docker Compose (optional but recommended)
- Node.js (for optional web frontend)

---

## 🔧 Build and Run

### 🧪 CLI Mode

```bash
# Port scan
go run main.go scan --target 192.168.0.1

# Spoofed email
go run main.go email attack --from spoof@microsoft.com --to victim@yourdomain.test

# SPF/DKIM/DMARC analysis
go run main.go email analyze --domain yourdomain.test

# IAM audit
go run main.go iam check --profile default

# S3 bucket audit
go run main.go s3 audit --bucket my-bucket --profile default

# DNS misconfiguration check
go run main.go dns cloud --domain example.com --subdomains www,api,staging

🌐 API Mode

# Start API server on port 8181
go run main.go

Then access endpoints at: http://localhost:8181/api
🐳 Docker

docker compose up --build

📬 API Endpoints
Method	Endpoint	Description
POST	/api/scan/port	Run a TCP port scan
POST	/api/email/attack	Send spoofed test email
POST	/api/email/analyze	Analyze SPF/DKIM/DMARC config
POST	/api/cloud/iam	Analyze AWS IAM policies
POST	/api/cloud/s3	Audit S3 bucket for exposure
POST	/api/cloud/dns	Scan for DNS CNAME misconfigs
🔐 Security & Ethics Notice

⚠️ GoSecOps is for educational and authorized testing only.
Never scan or spoof domains you do not own or have written permission to test.

    Spoofed emails are sandboxed via Mailhog/Maildev

    Designed for red/blue team simulations in test environments

    Logs and future features will include audit trails

🧪 Testing Environment

Use this in docker-compose.yml to test spoofed email output:

mailtest:
  image: maildev/maildev
  ports:
    - "1080:1080" # Mail Web UI
    - "1025:1025" # SMTP test port

📬 View captured emails at: http://localhost:1080
🛠️ Libraries Used
Purpose	Library
CLI Framework	github.com/spf13/cobra
HTTP API Server	github.com/gin-gonic/gin
Email Spoofing	net/smtp
AWS Cloud SDK	github.com/aws/aws-sdk-go-v2
DNS Lookups	net.LookupTXT / net.LookupIP
Port Scanning	net.DialTimeout
📦 Roadmap

    ✅ Cloud IAM, S3, DNS misconfig modules

    ✅ SvelteKit frontend dashboard

    ⏳ Security header analyzer

    ⏳ CSV/JSON report exports

    ⏳ WebSocket log streams

    ⏳ Role-based API auth (JWT/OAuth)

👨‍💻 Dev Commands

# Run a CLI audit
go run main.go s3 audit --bucket my-bucket --profile default

# Start API
go run main.go

# Docker dev environment
docker compose up --build

🤝 Contributing

Pull requests, feedback, and security reviews are welcome.
Please include test coverage for new modules/features.
📜 License

MIT License – see LICENSE file.