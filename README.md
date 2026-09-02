# ⚡ HyperPDF — Next-Gen All-in-One Online PDF Suite

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Fiber-v2-000000?style=for-the-badge&logo=gofiber&logoColor=white" alt="Fiber" />
  <img src="https://img.shields.io/badge/Vue-3-4FC08D?style=for-the-badge&logo=vuedotjs&logoColor=white" alt="Vue 3" />
  <img src="https://img.shields.io/badge/TypeScript-5.0+-3178C6?style=for-the-badge&logo=typescript&logoColor=white" alt="TypeScript" />
  <img src="https://img.shields.io/badge/TailwindCSS-3.4+-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white" alt="Tailwind" />
  <img src="https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>

A modern, blazing-fast, and secure full-stack **PDF Suite** (Compress, Merge, Split, Auth, and Subscriptions) built with **Go (Fiber v2, GORM, PostgreSQL, Ghostscript/QPDF)** and **Vue 3 (TypeScript, Vite, Tailwind CSS)**.

---

## 📑 Table of Contents

1. [Features](#-features)
2. [Compression Levels & Freemium Model](#-compression-levels--freemium-model)
3. [Tech Stack](#-tech-stack)
4. [Architecture Overview](#-architecture-overview)
5. [Prerequisites](#-prerequisites)
6. [Installation & Setup](#-installation--setup)
   - [Ghostscript & QPDF Installation](#ghostscript--qpdf-installation)
   - [Environment Configuration](#environment-configuration)
   - [Database Setup & Migrations](#database-setup--migrations)
7. [Running Locally](#-running-locally)
   - [Backend](#backend)
   - [Frontend](#frontend)
8. [Running with Docker Compose](#-running-with-docker-compose)
9. [REST API Documentation & Swagger](#-rest-api-documentation--swagger)
10. [Testing](#-testing)
11. [Security & Privacy](#-security--privacy)
12. [Troubleshooting](#-troubleshooting)

---

## ✨ Features

### ⚡ 1. PDF Compressor (5 Smart Levels)
- **Ultra Extreme (~50 DPI)**: Massive reduction (**~80–95%**) for strict portal file limits (e.g. CPNS, job portals).
- **Extreme (~72 DPI)**: High compression (**~70–85%**) for quick email and web sharing.
- **Balanced / Recommended (~150 DPI)**: Best balance (**~50–70%**) for daily office documents, ebooks, and coursework.
- **High-Fidelity (~220 DPI) [PRO ⭐]**: Crisp vector typography and charts for business presentations.
- **Studio Master Lossless (300+ DPI) [PRO ⭐]**: Prepress print quality and long-term digital legal archives.

### 📑 2. Merge PDF
- Combine multiple PDF documents in one sequential file.
- Interactive **visual reordering** (*Move Up / Move Down*) and individual file removal before merging.
- Instant single-click download of merged document.

### ✂️ 3. Split PDF
- **Extract Specific Page Ranges**: Extract pages (e.g. `1-3`, `2`) into a single output PDF.
- **Split Every Page to ZIP Archive**: Separates every page into distinct PDF files and packages them into a clean `.zip` archive.

### 🔐 4. Authentication & User Accounts
- Modern **Sign In** and **Sign Up / Create Account** dialog with email/password validation.
- Session persistence via `localStorage` remembering user credentials and plan status across browser refreshes.
- Profile dropdown with user name, email, subscription badge (**FREE** or **⭐ PRO**), and logout options.

### 💎 5. Subscription & Freemium Paywall (`PricingModal`)
- **Starter Free**: 50 MB max file size, 1.000 compressions/day quota with live decrement tracker.
- **HyperPDF PRO (Rp 49.000 / bln)**: 500 MB upload limit, 300+ DPI Studio Master quality, unlimited batch operations, dedicated turbo engine queue.
- Interactive **1-click checkout simulation** that activates PRO immediately.

### 📱 6. Modern Responsive UI (Glassmorphism)
- Designed with *Plus Jakarta Sans* & *Inter* typography.
- **Mobile Drawer**: Responsive sliding hamburger drawer for phones and tablets.
- **Desktop Dropdown**: Sleek "All PDF Tools" menu and header navigation pills.

---

## 📊 Compression Levels & Freemium Model

| Level | Tier | Target DPI | File Reduction | Best For |
| :--- | :---: | :---: | :---: | :--- |
| **Ultra Extreme** | 🟢 **FREE** | ~50 DPI | **~80% – 95%** | CPNS & government portals with strict KB limits |
| **Extreme** | 🟢 **FREE** | ~72 DPI | **~70% – 85%** | Email attachments & rapid messaging |
| **Balanced [REC]** | 🟢 **FREE** | ~150 DPI | **~50% – 70%** | Daily workplace reports, ebooks & coursework |
| **High-Fidelity** | 👑 **PRO ⭐** | ~220 DPI | **~25% – 45%** | Client presentations, portfolios & marketing decks |
| **Studio Master** | 👑 **PRO ⭐** | 300+ DPI | **~10% – 25%** | Prepress publication, blueprints & legal archives |

> **Daily Free Limit**: Free accounts enjoy **1.000 compressions per day** with a live real-time countdown counter (`Free Quota: 998 / 1.000 sisa hari ini`).

---

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.22+
- **Framework**: [Fiber v2](https://gofiber.io/) (High-performance Express-like web framework)
- **ORM & DB**: [GORM](https://gorm.io/) with PostgreSQL Driver
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
- **Logging**: [Logrus](https://github.com/sirupsen/logrus) (Structured JSON / text logging)
- **PDF Engines**: Ghostscript (`gs` / `gswin64c`) & QPDF (`qpdf`)
- **API Docs**: Swagger / [swag](https://github.com/swaggo/swag)

### Frontend
- **Framework**: Vue 3 (Composition API `<script setup>`)
- **Language**: TypeScript
- **Tooling & Bundler**: Vite
- **Styling**: Tailwind CSS 3.4+
- **HTTP Client**: Axios

---

## 🏛️ Architecture Overview

```
                          ┌───────────────────────────┐
                          │   Vue 3 + Tailwind UI     │
                          │   (HyperPDF Frontend)     │
                          └─────────────┬─────────────┘
                                        │ (REST API / JSON)
                                        ▼
                          ┌───────────────────────────┐
                          │     Fiber HTTP Router     │
                          │    & Middleware Stack     │
                          └─────────────┬─────────────┘
                                        │
                         ┌──────────────┴──────────────┐
                         ▼                             ▼
            ┌───────────────────────────┐ ┌───────────────────────────┐
            │   PostgreSQL Repository   │ │     PDF Service Engine    │
            │   (Metadata, UUIDs, Logs) │ │   (Upload, Validate, Exec)│
            └───────────────────────────┘ └─────────────┬─────────────┘
                                                        │
                                        ┌───────────────┴───────────────┐
                                        ▼                               ▼
                           ┌──────────────────────────┐   ┌──────────────────────────┐
                           │   Ghostscript Engine     │   │       QPDF Engine        │
                           │   (Raster & Downsampling)│   │   (Stream & Vector Opt)  │
                           └──────────────────────────┘   └──────────────────────────┘
```

### Folder Structure

```
hyperpdf/
├── backend/
│   ├── cmd/server/main.go            # Application entrypoint & routes
│   ├── docs/                         # Swagger OpenAPI 2.0 specifications
│   ├── internal/
│   │   ├── compressor/               # PDFCompressor interface (Ghostscript, QPDF, Mock)
│   │   ├── config/                   # Viper configuration manager
│   │   ├── database/                 # GORM initialization & migrations
│   │   ├── dto/                      # Request & Response structs with validators
│   │   ├── handler/                  # Fiber HTTP endpoints (Compress, Merge, Split)
│   │   ├── middleware/               # Logger, CORS, RateLimiter, Recover
│   │   ├── model/                    # GORM database models (PDFJob, Levels)
│   │   ├── repository/               # PostgreSQL job repository
│   │   ├── service/                  # PDF service orchestrator
│   │   └── utils/                    # Magic bytes, MIME validation, sanitize, zip helper
│   ├── migrations/                   # SQL migration scripts
│   ├── storage/                      # Storage directories (uploads/ and outputs/)
│   ├── test/                         # Unit tests & integration test suite
│   ├── Dockerfile
│   ├── go.mod
│   └── .env.example
├── frontend/
│   ├── src/
│   │   ├── api/                      # Axios API service
│   │   ├── components/               # DropZone, LevelSelector, MergeTool, SplitTool, AuthModal, PricingModal, Navbar, Footer
│   │   ├── composables/              # useAuth & usePdfCompressor state composables
│   │   ├── types/                    # TypeScript type definitions
│   │   ├── App.vue                   # Root Vue component with tool switcher
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
- **PostgreSQL**: 14+ (or run automatically with Docker Compose)
- **Ghostscript** or **QPDF** (preinstalled in Docker image)

---

## 🚀 Installation & Setup

### Ghostscript & QPDF Installation

#### Windows:
```powershell
# Install Ghostscript
winget install ArtifexSoftware.GhostScript

# Install QPDF
winget install QPDF.QPDF
```

#### Linux (Debian / Ubuntu):
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

Key environment variables:

| Variable | Default | Description |
| :--- | :--- | :--- |
| `APP_PORT` | `8080` | Port for the backend API |
| `APP_ENV` | `development` | Environment (`development` / `production`) |
| `DATABASE_HOST` | `localhost` | PostgreSQL host |
| `DATABASE_PORT` | `5432` | PostgreSQL port |
| `DATABASE_USER` | `postgres` | Database username |
| `DATABASE_PASSWORD` | `postgres` | Database password |
| `DATABASE_NAME` | `pdf_tools` | Database name |
| `DATABASE_SSLMODE` | `disable` | SSL mode |
| `STORAGE_UPLOAD_DIR` | `./storage/uploads` | Temporary upload folder |
| `STORAGE_OUTPUT_DIR` | `./storage/outputs` | Processed files storage folder |
| `MAX_FILE_SIZE_MB` | `50` | Maximum upload size in Megabytes |
| `COMPRESSOR_ENGINE` | `ghostscript` | Engine mode: `ghostscript`, `qpdf`, `mock` |
| `COMPRESS_TIMEOUT_SECONDS` | `120` | Execution timeout per job |
| `RATE_LIMIT_MAX` | `100` | Rate limiter maximum requests |
| `RATE_LIMIT_DURATION_SECONDS`| `60` | Rate limiter time window |

---

### Database Setup & Migrations

Create the PostgreSQL database:
```sql
CREATE DATABASE pdf_tools;
```

Migrations run automatically upon backend startup via `golang-migrate`. To run them manually:
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
The frontend will start at `http://localhost:3000` with hot-module replacement (HMR).

---

## 🐳 Running with Docker Compose

To start the entire stack (PostgreSQL + Backend with Ghostscript + Frontend with Nginx):

```bash
docker-compose up --build
```

Access services:
- **Web UI**: `http://localhost:3000`
- **Backend API**: `http://localhost:8080`
- **Swagger Documentation**: `http://localhost:8080/swagger/index.html`
- **Health Check**: `http://localhost:8080/health`

---

## 📖 REST API Documentation & Swagger

Interactive Swagger UI:
```
http://localhost:8080/swagger/index.html
```

### Endpoints

#### 1. Compress PDF
- **Endpoint**: `POST /api/v1/pdf/compress`
- **Content-Type**: `multipart/form-data`
- **Form Fields**:
  - `file`: (File, required) PDF file.
  - `compression_level`: (String, optional) `ULTRA_EXTREME`, `EXTREME`, `RECOMMENDED` (default), `HIGH_FIDELITY`, `STUDIO_MASTER`.
- **Response**:
```json
{
  "success": true,
  "message": "PDF compressed successfully",
  "data": {
    "id": "c8beb569-c1c5-45b9-9bda-949cd25c7f35",
    "original_filename": "annual_report.pdf",
    "original_size": 5242880,
    "compressed_size": 1835008,
    "saved_bytes": 3407872,
    "compression_percentage": 65.0,
    "compression_level": "RECOMMENDED",
    "status": "COMPLETED",
    "created_at": "2026-09-02T15:58:35Z"
  }
}
```

#### 2. Merge PDF Files
- **Endpoint**: `POST /api/v1/pdf/merge`
- **Content-Type**: `multipart/form-data`
- **Form Fields**:
  - `files`: (Multiple Files, required) 2 or more PDF files in sequential order.
- **Response**:
```json
{
  "success": true,
  "message": "PDF files merged successfully",
  "data": {
    "id": "aa0dcf7f-8cdf-4be1-8881-07ef7f223efc",
    "merged_filename": "merged_document.pdf",
    "file_count": 3,
    "total_size": 4194304,
    "download_url": "/api/v1/pdf/jobs/aa0dcf7f-8cdf-4be1-8881-07ef7f223efc/download"
  }
}
```

#### 3. Split PDF File
- **Endpoint**: `POST /api/v1/pdf/split`
- **Content-Type**: `multipart/form-data`
- **Form Fields**:
  - `file`: (File, required) Single PDF file.
  - `split_mode`: `range` or `all`.
  - `page_ranges`: (Optional, string) e.g. `1-3` or `2`.
- **Response**:
```json
{
  "success": true,
  "message": "PDF file split successfully",
  "data": {
    "id": "6564bc0b-3c73-4677-9403-2f1af0af05e9",
    "original_filename": "contract.pdf",
    "split_mode": "all",
    "generated_count": 5,
    "is_zip_archive": true,
    "download_filename": "split_pages_contract.zip",
    "download_url": "/api/v1/pdf/jobs/6564bc0b-3c73-4677-9403-2f1af0af05e9/download"
  }
}
```

#### 4. Download Processed File
- **Endpoint**: `GET /api/v1/pdf/jobs/:id/download`
- **Response**: Binary stream (`application/pdf` or `application/zip`) with `Content-Disposition: attachment; filename="..."`.

#### 5. Delete Job & Files
- **Endpoint**: `DELETE /api/v1/pdf/jobs/:id`

#### 6. System Health Check
- **Endpoint**: `GET /health`

---

## 🧪 Testing

### Backend Test Suite
```bash
cd backend
go test -v ./test/...
```

Test coverage includes:
- `TestSanitizeFilename`: Path traversal prevention (`../../evil.pdf`), illegal characters, empty names.
- `TestCalculateSavings`: Accurate percentage and byte reduction calculations.
- `TestValidatePDFFile`: Magic byte validation (`%PDF-`), MIME check, size checks.
- `TestPDFService_CompressPDF_Success`: Service compression workflow with mock engine.
- `TestPDFService_MergePDF`: Multi-file merging unit test.
- `TestPDFService_SplitPDF`: Page range extraction and zip packaging unit test.
- `TestAPI_CompressPDF_Endpoint`: Fiber HTTP upload & compression test.
- `TestAPI_MergePDF_Endpoint`: Fiber HTTP multi-file merge test.
- `TestAPI_SplitPDF_Endpoint`: Fiber HTTP page split test.
- `TestAPI_Download_Endpoint`: File download streaming test.
- `TestAPI_Delete_Endpoint`: Safe job deletion and file cleanup test.

### Frontend Typecheck & Build
```bash
cd frontend
npm run build
```

---

## 🛡️ Security & Privacy

1. **Magic Byte Signature Verification**: Analyzes the first 512 bytes of all uploads for the authentic `%PDF-` header signature, rejecting disguised executables.
2. **Path Traversal Sanitization**: User-supplied filenames are stripped of `../` and directory prefixes. File storage uses isolated UUID filenames.
3. **Execution Sandbox & Timeout**: Shell executions are bound to strict context timeouts (`exec.CommandContext`) preventing resource exhaustion attacks.
4. **Immediate Disk Cleanup**: Uploaded temporary files are wiped from disk immediately after processing.
5. **Rate Limiting**: Configurable Fiber middleware restricts excessive requests per IP.

---

## ❓ Troubleshooting

### Ghostscript binary not found
- **Symptom**: `PDF engine binary not found in PATH` in server logs.
- **Solution**:
  1. Install Ghostscript via `winget install ArtifexSoftware.GhostScript` (Windows) or `sudo apt install ghostscript` (Linux).
  2. Set explicit path in `.env`: `GHOSTSCRIPT_BINARY=C:\Program Files\gs\gs10.03.0\bin\gswin64c.exe`.
  3. Or run using `docker-compose up`, which has Ghostscript pre-installed.

### Local Development Simulator Mode
If running on a development machine without Ghostscript, set:
```env
COMPRESSOR_ENGINE=mock
```
HyperPDF will simulate realistic compression ratios for all 5 levels without external dependencies.

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
