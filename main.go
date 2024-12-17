package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/sidra-gateway/go-pdk/server"
)

func main() {
	// --- STEP 1: Ambil username dan password dari Environment Variables ---
	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")

	// --- STEP 2: Validasi jika Env Var kosong ---
	if username == "" || password == "" {
		log.Fatal("Error: Environment variables BASIC_AUTH_USERNAME and BASIC_AUTH_PASSWORD must be set.")
	}

	// --- STEP 3: Buat server plugin ---
	server.NewServer("basic-auth", func(r server.Request) server.Response {
		// --- STEP 4: Ambil header "Authorization" ---
		authHeader, ok := r.Headers["Authorization"]
		if !ok { // Jika header "Authorization" tidak ditemukan
			return server.Response{
				StatusCode: 401,
				Body:       "Unauthorized - Missing Authorization Header",
			}
		}

		// --- STEP 5: Generate expected Authorization value ---
		expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

		// --- STEP 6: Bandingkan Authorization Header dengan nilai yang diharapkan ---
		if authHeader != expectedAuth {
			return server.Response{
				StatusCode: 401,
				Body:       "Unauthorized - Invalid Credentials",
			}
		}

		// --- STEP 7: Jika valid, kirim respons sukses ---
		return server.Response{
			StatusCode: 200,
			Body:       "You are authenticated",
		}
	}).Start() // Mulai server Sidra Gateway
}