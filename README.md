# 🛡️ GoSecOps – Penetration Testing & Email Security Toolkit

GoSecOps is a modular, containerized security tool built with **GoLang**. It offers both **CLI** and **REST API** interfaces for red-team attack simulations (port scanning, spoofed emails) and blue-team validation (SPF/DKIM/DMARC checks, DNS recon, etc.).

Cloud Security Architect: Manaka Anthony Raphasha
---

## 🚀 Features

| Feature         | CLI             | API            | Status  |
|----------------|------------------|----------------|---------|
| Port Scanner    | ✅ `scan ports`  | ✅ `/scan/port` | ✅ Done |
| Email Spoofing  | ✅ `email attack`| ✅ `/email/attack` | ✅ Done |
| Email Analysis  | ✅ `email analyze` | ✅ `/email/analyze` | ✅ Done |
| Web UI (Svelte) | ❌ In progress   | ✅ Ready to connect | 🔧 |
| DNS Recon       | Coming soon     | Coming soon    | ⏳     |
| Logging         | JSON + stdout (planned) | | 🔧     |

---

## 🧱 Project Structure

gosecops/
├── api/
│ ├── handlers/ # REST handlers (port scan, email, etc.)
│ └── routes.go # API route mappings
├── cmd/ # CLI commands using Cobra
│ ├── root.go
│ ├── scan.go
│ └── email.go
├── internal/ # Core modules
│ ├── scanner/ # TCP port scanning
│ ├── email/ # Email attack & SPF/DKIM analysis
│ ├── dns/ # Subdomain brute force (soon)
│ ├── utils/ # Shared helpers
│ └── validator/ # Header analysis (planned)
├── web/ # Optional SvelteKit UI
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
- Node.js (if using the web frontend)

---

## 🔧 Build and Run

### CLI Mode
```bash
# Port scan CLI
go run main.go scan --target 192.168.0.1

# Send spoofed email
go run main.go email attack --from spoof@microsoft.com --to victim@yourdomain.test

# Analyze SPF/DKIM/DMARC
go run main.go email analyze --domain yourdomain.test

API Mode

# Start the API server on port 8181
go run main.go

Docker

docker compose up --build

Access API via http://localhost:8181/api.
📬 API Endpoints
Method	Endpoint	Description
POST	/api/scan/port	Run a TCP port scan
POST	/api/email/attack	Send spoofed email (test only)
POST	/api/email/analyze	Analyze SPF/DKIM/DMARC
🔐 Security & Ethics Notice

    ⚠️ GoSecOps is for educational and authorized testing only. Never run this toolkit against production systems or networks you do not own or have explicit permission to test.

    Spoofed emails are routed to Mailhog/Maildev in Docker.

    All modules are sandboxed to avoid unsafe behavior.

🧪 Testing Environment

Use the following docker-compose.yml service to test emails:

mailtest:
  image: maildev/maildev
  ports:
    - "1080:1080" # Web UI
    - "1025:1025" # SMTP

Access test emails at: http://localhost:1080
🛠️ Libraries Used
Purpose	Library
CLI	github.com/spf13/cobra
API Server	github.com/gin-gonic/gin
Email	net/smtp
DNS Lookups	net.LookupTXT
Port Scanning	net.DialTimeout
📦 Coming Soon

    ✅ DNS Recon (dnsrecon)

    ✅ WebSocket logs

    ✅ Auth-protected Web UI (SvelteKit)

    ✅ Report export (JSON/CSV)

    ✅ Logger module

👨‍💻 Development Commands

# Run CLI with arguments
go run main.go email analyze --domain example.com

# Start API
go run main.go

# Docker up
docker compose up --build

🤝 Contributing

Pull requests, feedback, and security reviews are welcome. Please ensure tests are added for any new features.
📜 License

MIT License – see LICENSE
