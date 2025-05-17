# 🔐 Vertex (Go + SQLite)

A simple yet **powerful password generator backend** written in Go that ensures every password generated is **truly unique**. It combines rare word roots, strong symbols, numbers, and mutations to make structured, secure, and **non-repeating** passwords.

---

## 🚀 Features

- 🔄 Generates **never-seen-before** passwords
- 🧠 Uses **mutated rare words + numbers + symbols**
- 🔒 Secure randomness with `crypto/rand`
- 💾 Stores generated passwords in **SQLite**
- ⚡ Simple REST API (`/generate`)
- 🛡️ Prevents duplicate passwords

---

## 📦 Tech Stack

- **Language**: Go
- **Database**: SQLite (via `github.com/mattn/go-sqlite3`)
- **API**: net/http (standard library)

---
## 🧪 Sample Password Output

```json
{
  "password": "s0lArPhantOmQuakE92!",
  "timestamp": "2025-05-17T09:35:21Z",
  "entropyBits": 128
}
