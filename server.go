package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func anasayfa(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("<title>Anasayfa</title>\n"))
	w.Write([]byte("<center><h1>Anasayfa</h1></center><br>\n"))
	w.Write([]byte("<center><h3><a href='/listele'>Listele</a></h3></center><br>"))
}
func listeleme(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<title>20703043-Dizin Listeleme</title>\n"))
	w.Write([]byte("<h1>Dizin Listeleme</h1><br>\n"))
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if f.IsDir() {
			w.Write([]byte("dizin>_\t" + f.Name() + "<br>\n"))
		} else {
			w.Write([]byte("dosya>_\t" + f.Name() + "<br>\n"))
		}
	}
}
func main() {
	http.HandleFunc("/listele", listeleme)
	http.HandleFunc("/", anasayfa)
	server := &http.Server{
		Addr: "127.0.0.1:1337",
	}
	fmt.Println("Server http://localhost:1337/ adresinde başlatıldı.")
	log.Fatal(server.ListenAndServe())
}
