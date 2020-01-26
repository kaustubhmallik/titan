package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		res := map[string]string{"status": "OK"}
		resB, _ := json.Marshal(res)
		_, _ = fmt.Fprint(w, string(resB))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
