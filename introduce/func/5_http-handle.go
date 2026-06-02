package ilovemyjobPkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Name == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := ResponseBody{Message: fmt.Sprintf("Hello %s", req.Name)}
	json.NewEncoder(w).Encode(resp)
}

func HttpHandle() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
