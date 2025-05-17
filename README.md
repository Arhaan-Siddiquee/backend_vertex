# 🔐 Vertex – Unique Password Generator API

A simple yet **unique password generator API** written in Go. Each generated password is **never repeated**, stored in a local SQLite database, and includes creative entropy-based mutations for real uniqueness.

---

## 🚀 Features

- 100% Unique password every time
- Pure Go (✅ No CGO / C compiler needed)
- SQLite database to track generated passwords
- Creative, never-before-seen combinations
- JSON API ready for frontend use

---

## 📦 Tech Stack

- Language: Go (Golang)
- Database: SQLite (using `modernc.org/sqlite`)
- Driver: Pure Go SQLite (No CGO required)
- Server: Go's native HTTP package

---

## 🧪 Example Passwords

```json
{
  "password": "Ph@nt0mCrypTdr1ft53!",
  "timestamp": "2025-05-17T20:49:00Z",
  "entropyBits": 128
}
```

## 🛠️ Installation
1. Clone the repository
```bash
git clone https://github.com/yourusername/trueuniquepass.git
cd trueuniquepass
```
2. Initialize Go modules
```bash
go mod init trueuniquepass
```
3. Install the dependencies
```bash
go get modernc.org/sqlite
```
4. Run the server
```bash
go run main.go
```
5. Open Request
```bash
http://localhost:8080/generate
```
