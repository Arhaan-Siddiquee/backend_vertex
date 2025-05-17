package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "./passwords.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	http.HandleFunc("/generate", generateHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS passwords (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			password TEXT UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	var password string
	var isUnique bool

	for attempts := 0; attempts < 10; attempts++ {
		password = generatePassword()
		isUnique = checkAndStorePassword(password)
		if isUnique {
			break
		}
	}

	if !isUnique {
		http.Error(w, "Failed to generate a unique password", http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"password":    password,
		"timestamp":   time.Now().UTC().Format(time.RFC3339),
		"entropyBits": 128,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func checkAndStorePassword(pwd string) bool {
	stmt, err := db.Prepare("INSERT INTO passwords(password) VALUES(?)")
	if err != nil {
		log.Println("Prepare failed:", err)
		return false
	}
	_, err = stmt.Exec(pwd)
	return err == nil
}

var wordPool = []string{"solar", "phantom", "drift", "flame", "nova", "crypt", "aether", "quake", "venom"}
var symbols = []rune("!@#$%^&*()_+~∆¥")

func generatePassword() string {
	word1 := mutateWord(randomWord())
	word2 := mutateWord(randomWord())
	word3 := mutateWord(randomWord())

	num := randomNumber(10, 99)
	symbol := string(symbols[randomInt(len(symbols))])

	return fmt.Sprintf("%s%s%s%d%s", word1, strings.Title(word2), strings.Title(word3), num, symbol)
}

func randomWord() string {
	return wordPool[randomInt(len(wordPool))]
}

func randomNumber(min, max int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return int(nBig.Int64()) + min
}

func randomInt(limit int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	return int(nBig.Int64())
}

func mutateWord(word string) string {
	var mutated []rune
	for _, ch := range word {
		if flip := randomInt(2); flip == 1 {
			ch = flipCase(ch)
		}
		if sub := mutateChar(ch); sub != 0 {
			ch = sub
		}
		mutated = append(mutated, ch)
	}
	return string(mutated)
}

func flipCase(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	} else if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func mutateChar(ch rune) rune {
	switch ch {
	case 'a', 'A':
		return '@'
	case 'i', 'I':
		return '1'
	case 'o', 'O':
		return '0'
	case 'e', 'E':
		return '3'
	default:
		return 0
	}
}
