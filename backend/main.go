package main
import (
	"fmt"
	"net/http"
    "backend/config"
)

func main() {

	_, err := config.ConnectDB()
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}