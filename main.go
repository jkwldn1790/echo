package main

import (
	"html/template"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path"
)

type Metadata struct {
	Hostname string
	LocalIp  net.IP
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	slog.Info("got / request\n")

	fp := path.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name, err := os.Hostname()

	address := GetOutboundIP()

	metadata := Metadata{name, address}

	if err := tmpl.Execute(w, metadata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	slog.Info("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	logger.Info("Echo service Online")
	http.ListenAndServe(":3333", nil)
}
