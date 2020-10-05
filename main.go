package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bs, _ := ioutil.ReadFile("index.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(bs)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		f, h, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		localF, err := os.Create(h.Filename)
		defer localF.Close()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		io.Copy(localF, f)
		w.Write([]byte("ok"))
	})
	ip := LocalIP()
	port := "9999"
	fmt.Println(ip + ":" + port)
	http.ListenAndServe(":"+port, nil)
}

// LocalIP 获取本地IP
func LocalIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, v := range addrs {
		ip := strings.Split(v.String(), "/")[0]
		if strings.HasPrefix(ip, "192.") || strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "172.") {
			return ip
		}
	}
	return ""
}
