package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to cloud native</h1>"))

	//1.写入Header
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	//2. Version
	os.Setenv("VERSION", "0.0.5")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	//3.get client ip
	clientIP := getClientIP(r)
	httpCode := http.StatusOK
	log.Printf("clientIP: %s Response code: %d \n", clientIP, httpCode)

}

// get client real ip address
func getClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "APP Working")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("start server failed, %s \n", err.Error())
	}
}
