# 🛡️ GoSecOps – Penetration Testing & Cloud Security Toolkit

**GoSecOps** is a modular, containerized security toolkit built with **GoLang**. It supports both **CLI** and **REST API** interfaces for red-team attack simulation (e.g., port scanning, spoofed emails) and blue-team validation (e.g., SPF/DKIM/DMARC analysis, IAM audits, DNS misconfigs, and S3 exposure).

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
├── main.go # Entry point
├── go.mod # Go module metadata
└── README.md # Project documentation


---

## 💻 Getting Started

### ✅ Prerequisites

- Go 1.20+
- Docker + Docker Compose (recommended)
- Node.js (for optional frontend)

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

# DNS CNAME Misconfiguration Scan
go run main.go dns cloud --domain example.com --subdomains www,api,staging

🌐 API Server

# Start API server on http://localhost:8181
go run main.go

🐳 Docker Mode

docker compose up --build

Then access the API at:

http://localhost:8181/api
http://localhost:8181/swagger/index.html

✅ Swagger UI Access (📚 API Documentation)

## 📚 Swagger UI (Interactive API Docs)

Once the API server is running, access the Swagger UI at:

➡️ [http://localhost:8181/swagger/index.html](http://localhost:8181/swagger/index.html)

This provides:

- 📖 **Full documentation** of each API route
- 🧪 **Built-in testing interface** (send requests directly from the browser)
- 📂 **Schema definitions** for request and response bodies

> You can also export your API definition as OpenAPI JSON/YAML from the UI

✅ Where to Insert It

This section should come right after 📬 API Endpoints and before 🔐 Security & Ethics Notice.

So now the sequence in your README would be:

📬 API Endpoints
📚 Swagger UI (Interactive API Docs)
🔐 Security & Ethics Notice

✅ Resulting Snippet Example

## 📬 API Endpoints

| Method | Endpoint             | Description                             |
|--------|----------------------|-----------------------------------------|
| POST   | `/api/scan/port`     | Run a TCP port scan                     |
| POST   | `/api/email/attack`  | Send a spoofed test email               |
| POST   | `/api/email/analyze` | Analyze SPF/DKIM/DMARC                  |
| POST   | `/api/cloud/iam`     | Audit AWS IAM policies                  |
| POST   | `/api/cloud/s3`      | Detect public exposure of S3 buckets    |
| POST   | `/api/cloud/dns`     | Scan for DNS CNAME misconfigurations    |

---

## 📚 Swagger UI (Interactive API Docs)

Once the API server is running, access the Swagger UI at:

➡️ [http://localhost:8181/swagger/index.html](http://localhost:8181/swagger/index.html)

This provides:

- 📖 **Full documentation** of each API route
- 🧪 **Built-in testing interface** (send requests directly from the browser)
- 📂 **Schema definitions** for request and response bodies

> You can also export your API definition as OpenAPI JSON/YAML from the UI


⚠️ GoSecOps is for educational and authorized testing only.

    Never scan or spoof any system or domain you do not own or have explicit permission to test.

    Spoofed emails are routed to Mailhog/Maildev in isolated testing environments.

    Logs and audit trails are in planning to ensure safe usage and traceability.

🧪 Testing Environment (Email Spoofing)

Include this in your docker-compose.yml:

mailtest:
  image: maildev/maildev
  ports:
    - "1080:1080"  # Mail Web UI
    - "1025:1025"  # SMTP test port

📬 Access captured emails at: http://localhost:1080
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

    ✅ Swagger API documentation

    ✅ SvelteKit frontend (in progress)

    ⏳ Security header analyzer

    ⏳ JSON/CSV report export

    ⏳ WebSocket log streams

    ⏳ Role-based API auth (JWT/OAuth)

👨‍💻 Dev Commands

# Run a scan
go run main.go s3 audit --bucket my-bucket --profile default

# Start the API server
go run main.go

# Docker development environment
docker compose up --build


🤝 Contributing

Contributions are welcome!
Please include unit tests for any new modules or features.
We appreciate pull requests, feedback, and security reviews.
📜 License

MIT License – see LICENSE