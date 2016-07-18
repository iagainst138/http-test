package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func ListAddrs() []string {
	r := regexp.MustCompile("^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$")
	localAddrs := make([]string, 0)

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.String() != "127.0.0.1" && ip.String() != "::1" {
				if r.MatchString(ip.String()) {
					localAddrs = append(localAddrs, ip.String())
				}
			}
		}
	}
	return localAddrs
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("accepted: %v %v", r.RemoteAddr, r.RequestURI)
	h, _ := os.Hostname()
	fmt.Fprintf(w, "request processed by %s [%s]\n", h, strings.Join(ListAddrs(), ","))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Println("listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
