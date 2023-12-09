package main

import (
	"fmt"
	"net/http"
	"os"
)

const correctPassword = "your-password-here"

func main() {
	http.HandleFunc("/", handleRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on %s\n", address)
	http.ListenAndServe(address, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		password := r.Form.Get("password")
		if password == correctPassword {
			fmt.Fprint(w, "認証成功")
		} else {
			http.Error(w, "認証失敗", http.StatusUnauthorized)
		}
		return
	}

	html := `
		<form method="POST" action="/">
			<label>
				パスワード:
				<input type="password" name="password" required>
			</label>
			<button type="submit">送信</button>
		</form>
	`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
