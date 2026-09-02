# 📄 PDF Compressor (iLovePDF Style)

A modern, high-performance, and secure full-stack **PDF Compressor** web application built with **Go (Fiber v2, GORM, PostgreSQL, Ghostscript/qpdf engine)** and **Vue 3 (TypeScript, Vite, Tailwind CSS)**.

---

## 📑 Table of Contents
1. [Features](#-features)
2. [Tech Stack](#-tech-stack)
3. [Architecture Overview](#-architecture-overview)
4. [Prerequisites](#-prerequisites)
5. [Installation & Setup](#-installation--setup)
   - [Ghostscript / QPDF Installation](#ghostscript--qpdf-installation)
   - [Environment Configuration](#environment-configuration)
   - [Database Setup & Migrations](#database-setup--migrations)
6. [Running Locally](#-running-locally)
   - [Backend](#backend)
   - [Frontend](#frontend)
7. [Running with Docker Compose](#-running-with-docker-compose)
8. [REST API Documentation & Swagger](#-rest-api-documentation--swagger)
9. [Testing](#-testing)
10. [Security & Validation](#-security--validation)
11. [Troubleshooting](#-troubleshooting)

---

## ✨ Features

- **Drag & Drop Upload**: Modern, interactive upload zone with instant file validation.
- **Smart Compression Levels**:
  - **LOW**: High quality preservation (~300 DPI, `-dPDFSETTINGS=/printer`).
  - **MEDIUM (Recommended)**: Optimal balance of visual sharpness and reduction (~150 DPI, `-dPDFSETTINGS=/ebook`).
  - **HIGH**: Maximum compression footprint reduction (~72 DPI, `-dPDFSETTINGS=/screen`).
- **Real-Time Progress & Status**: Smooth multi-stage animation tracking upload, compression, and result finalization.
- **Detailed Comparison Metrics**: Visual before vs. after file size comparison, bytes saved, and reduction percentage.
- **Instant Secure Download**: Direct download with sanitized original file attachment headers.
- **Secure File Lifecycle**:
  - Magic byte verification (`%PDF-` signature check).
  - Strict MIME validation (`application/pdf`).
  - Filename sanitization & path traversal mitigation (`../../`).
  - Automatic temporary input file wiping immediately after processing.
  - Job deletion endpoint (`DELETE /api/v1/pdf/jobs/:id`) to clean files on demand.
- **Swagger / OpenAPI 2.0**: Interactive API playground at `/swagger/index.html`.

---

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.22+
- **Framework**: [Fiber v2](https://gofiber.io/)
- **ORM & DB**: [GORM](https://gorm.io/) with PostgreSQL Driver
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
- **Logging**: [Logrus](https://github.com/sirupsen/logrus) (Structured logging)
- **PDF Engine**: Ghostscript (`gs` / `gswin64c`) & QPDF (`qpdf`)
- **API Docs**: Swagger / [swag](https://github.com/swaggo/swag)

### Frontend
- **Framework**: Vue 3 (Composition API `<script setup>`)
- **Language**: TypeScript
- **Tooling & Bundler**: Vite
- **Styling**: Tailwind CSS
- **HTTP Client**: Axios

---

## 🏛️ Architecture Overview

```
User Browser (Vue 3 + Tailwind CSS)
            │
            ▼ (HTTP / REST API)
    Fiber HTTP Handler
            │
            ▼ (DTOs & Validation)
     PDF Service Layer (Orchestrator)
      ├── PostgreSQL Repository (Job Metadata)
      └── PDF Compressor Interface
            ├── Ghostscript Engine (exec.CommandContext)
            └── QPDF Engine (Fallback)
```

### Folder Structure

```
cobagolang/
├── backend/
│   ├── cmd/server/main.go            # Application entrypoint
│   ├── docs/                         # Swagger generated API docs
│   ├── internal/
│   │   ├── compressor/               # PDFCompressor interface, Ghostscript & QPDF
│   │   ├── config/                   # Viper configuration loader
│   │   ├── database/                 # GORM init & golang-migrate runner
│   │   ├── dto/                      # Request & Response structs
│   │   ├── handler/                  # Fiber HTTP route handlers
│   │   ├── middleware/               # Logger, CORS, Recover, Rate Limiter
│   │   ├── model/                    # GORM database models
│   │   ├── repository/               # PostgreSQL job repository
│   │   ├── service/                  # PDF business logic orchestrator
│   │   └── utils/                    # Magic bytes, MIME validation, sanitize, math
│   ├── migrations/                   # SQL migration scripts
│   ├── storage/                      # Storage directories for uploads and outputs
│   ├── test/                         # Unit tests & Fiber integration tests
│   ├── Dockerfile
│   ├── go.mod
│   └── .env.example
├── frontend/
│   ├── src/
│   │   ├── api/                      # Axios API service
│   │   ├── components/               # DropZone, LevelSelector, ProcessingCard, ResultCard, Navbar, Footer
│   │   ├── composables/              # usePdfCompressor state composable
│   │   ├── types/                    # TypeScript type definitions
│   │   ├── App.vue                   # Root Vue component
│   │   └── main.ts                   # Frontend bootstrap
│   ├── Dockerfile
│   ├── nginx.conf
│   ├── package.json
│   ├── vite.config.ts
│   └── tailwind.config.js
├── docker-compose.yml
├── .gitignore
└── README.md
```

---

## 📋 Prerequisites

- **Go**: 1.22 or newer
- **Node.js**: 18+ and npm
- **PostgreSQL**: 14+ (or run via Docker)
- **Ghostscript** or **QPDF** installed locally (or run via Docker where they are preinstalled)

---

## 🚀 Installation & Setup

### Ghostscript / QPDF Installation

#### Windows:
Using Winget:
```powershell
winget install ArtifexSoftware.GhostScript
# Or QPDF:
winget install qpdf
```

Using Chocolatey:
```powershell
choco install ghostscript
# Or QPDF:
choco install qpdf
```

#### Linux (Debian/Ubuntu):
```bash
sudo apt update
sudo apt install -y ghostscript qpdf
```

#### macOS:
```bash
brew install ghostscript qpdf
```

---

### Environment Configuration

In `backend/`, copy `.env.example` to `.env`:

```bash
cd backend
cp .env.example .env
```

Default configuration variables:

| Variable | Default | Description |
| :--- | :--- | :--- |
| `APP_PORT` | `8080` | Port for the Go backend server |
| `APP_ENV` | `development` | Environment (`development` / `production`) |
| `DATABASE_HOST` | `localhost` | PostgreSQL host |
| `DATABASE_PORT` | `5432` | PostgreSQL port |
| `DATABASE_USER` | `postgres` | Database user |
| `DATABASE_PASSWORD` | `postgres` | Database password |
| `DATABASE_NAME` | `pdf_tools` | Database name |
| `DATABASE_SSLMODE` | `disable` | SSL mode (`disable`, `require`, etc.) |
| `STORAGE_UPLOAD_DIR` | `./storage/uploads` | Temporary upload folder |
| `STORAGE_OUTPUT_DIR` | `./storage/outputs` | Compressed files storage folder |
| `MAX_FILE_SIZE_MB` | `50` | Maximum upload size in Megabytes |
| `COMPRESSOR_ENGINE` | `ghostscript` | Engine to use: `ghostscript`, `qpdf`, `auto` |
| `COMPRESS_TIMEOUT_SECONDS` | `120` | Compression timeout in seconds |
| `RATE_LIMIT_MAX` | `100` | Max requests per rate limit duration |
| `RATE_LIMIT_DURATION_SECONDS`| `60` | Rate limiter window in seconds |

---

### Database Setup & Migrations

Create the PostgreSQL database:
```sql
CREATE DATABASE pdf_tools;
```

Migrations run automatically upon backend startup via `golang-migrate`. If you want to run migrations manually using the `golang-migrate` CLI:

```bash
migrate -path backend/migrations -database "postgres://postgres:postgres@localhost:5432/pdf_tools?sslmode=disable" up
```

---

## 🏃 Running Locally

### Backend

```bash
cd backend
go run cmd/server/main.go
```
The server will start at `http://localhost:8080`.

### Frontend

```bash
cd frontend
npm install
npm run dev
```
The frontend will start at `http://localhost:3000` and automatically proxy API calls to backend port `8080`.

---

## 🐳 Running with Docker Compose

The easiest way to run the entire stack (PostgreSQL + Backend with Ghostscript + Frontend with Nginx):

```bash
docker-compose up --build
```

Access the services:
- **Frontend Web UI**: `http://localhost:3000`
- **Backend API**: `http://localhost:8080`
- **Swagger Documentation**: `http://localhost:8080/swagger/index.html`
- **Health Check**: `http://localhost:8080/health`

To stop all containers:
```bash
docker-compose down
```

---

## 📖 REST API Documentation & Swagger

Interactive Swagger UI is accessible at:
```
http://localhost:8080/swagger/index.html
```

### Endpoints Overview

#### 1. Compress PDF
- **Endpoint**: `POST /api/v1/pdf/compress`
- **Content-Type**: `multipart/form-data`
- **Form Fields**:
  - `file`: (File, required) The PDF file to compress.
  - `compression_level`: (String, optional) `LOW`, `MEDIUM` (default), or `HIGH`.
- **Response**:
```json
{
  "success": true,
  "message": "PDF compressed successfully",
  "data": {
    "id": "787c807b-c9a9-467f-8561-9f2010839c0f",
    "original_filename": "annual_report.pdf",
    "original_size": 5242880,
    "compressed_size": 1835008,
    "saved_bytes": 3407872,
    "compression_percentage": 65.0,
    "compression_level": "MEDIUM",
    "status": "COMPLETED",
    "created_at": "2026-09-02T14:30:00Z",
    "completed_at": "2026-09-02T14:30:04Z"
  }
}
```

#### 2. Get Job Details
- **Endpoint**: `GET /api/v1/pdf/jobs/:id`
- **Response**:
```json
{
  "success": true,
  "message": "Job retrieved successfully",
  "data": {
    "id": "787c807b-c9a9-467f-8561-9f2010839c0f",
    "original_filename": "annual_report.pdf",
    "original_size": 5242880,
    "compressed_size": 1835008,
    "saved_bytes": 3407872,
    "compression_percentage": 65.0,
    "compression_level": "MEDIUM",
    "status": "COMPLETED"
  }
}
```

#### 3. Download Compressed PDF
- **Endpoint**: `GET /api/v1/pdf/jobs/:id/download`
- **Response**: Returns binary stream `application/pdf` with `Content-Disposition: attachment; filename="compressed_annual_report.pdf"`.

#### 4. Delete Job & Files
- **Endpoint**: `DELETE /api/v1/pdf/jobs/:id`
- **Response**:
```json
{
  "success": true,
  "message": "Job and associated files deleted successfully",
  "data": null
}
```

#### 5. Health Check
- **Endpoint**: `GET /health`
- **Response**:
```json
{
  "success": true,
  "message": "System is healthy",
  "data": {
    "database": "connected",
    "engine_status": "available",
    "pdf_engine": "ghostscript",
    "status": "ok",
    "uptime": "1h24m10s"
  }
}
```

---

## 🧪 Testing

### Backend Tests
Execute unit tests and Fiber integration tests:

```bash
cd backend
go test -v ./test/...
```

Tests include:
- `TestSanitizeFilename`: Path traversal prevention (`../../evil.pdf`), illegal characters, empty names.
- `TestCalculateSavings`: File reduction and percentage math accuracy.
- `TestFormatBytes`: Human-readable size converter tests.
- `TestValidatePDFFile`: Magic bytes validation (`%PDF-`), MIME check, size check.
- `TestPDFService_CompressPDF_Success`: Service workflow with mock compressor.
- `TestPDFService_CompressPDF_EngineFailure`: Error handling & failure state persistence.
- `TestAPI_CompressPDF_Endpoint`: Fiber HTTP upload & compression test.
- `TestAPI_GetJob_Endpoint`: Fiber HTTP job detail test.
- `TestAPI_Download_Endpoint`: Fiber HTTP file download stream test.
- `TestAPI_Delete_Endpoint`: Fiber HTTP job deletion and cleanup test.

### Frontend TypeScript Check & Build
```bash
cd frontend
npm run build
```

---

## 🛡️ Security & Validation

1. **Magic Bytes Header Verification**: Reads first 512 bytes to ensure `%PDF-` signature is present. Prevents malicious executable files disguised as `.pdf`.
2. **Path Traversal Mitigation**: User filenames are stripped of `../` and directory components. Internal storage uses UUIDs (`<uuid>_compressed.pdf`).
3. **Command Injection Prevention**: Parameters are strictly mapped to static Ghostscript preset flags (`-dPDFSETTINGS=...`) and executed safely with `exec.CommandContext` without raw shell interpretation.
4. **Execution Timeouts**: Ghostscript/QPDF execution is bound to context timeouts to prevent resource exhaustion from corrupted or recursive PDFs.
5. **Rate Limiting & Body Limits**: Fiber middleware enforces maximum file size and request rate limits.
6. **Automatic Cleanup**: Input files are deleted immediately after compression finishes.

---

## ❓ Troubleshooting

### Ghostscript binary not found
- **Symptom**: `PDF engine binary not found in PATH` in server logs.
- **Solution**:
  1. Install Ghostscript via `winget install ArtifexSoftware.GhostScript` (Windows) or `sudo apt install ghostscript` (Linux).
  2. Or set explicit path in `.env`: `GHOSTSCRIPT_BINARY=C:\Program Files\gs\gs10.03.0\bin\gswin64c.exe`.
  3. Or run the app using `docker-compose up`, which has Ghostscript pre-installed.

### Database connection refused
- **Symptom**: `PostgreSQL connection failed`.
- **Solution**: Ensure PostgreSQL is running on port 5432 or use Docker Compose (`docker-compose up -d postgres`).
