// Filen for å eksperimentere med webserver i ICA03 (IS-105)
// Starte server med
// 			go run server.go
// Du kan ha tilgang til server fra nettleser med
// 			http://localhost:3000
//
package main

import (
	"net/http"
	"time"

	"fmt"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
	Now()
}

func foo(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
	//w.Write([]byte("20060102150405"))
	// Her kan man bl. a. skrive data som skal sendes til nettleser
	// Neste linje kan brukes for å endre koding av sekvensen som sendes til nettleser
	// Standardinstilling er UTF-8
	//w.Header().Set("Content-Type", "text/html;charset=ISO-8859-1")

	// Her skriver man data som er respons til brukeren som har skrevet
	// http://localhost:3000 i sin nettleser
	w.Write([]byte("<font color=\"pink\">Hvordan g\xe5r det, <b>\u16a6</b> ?</font><br/>"))
	w.Write([]byte("\u16a6 - Thurs<br/>"))
	w.Write([]byte("\xE2\x8F\xB0"))


}

func Now() {
	//t := time.Now()

	//fmt.Println(t.Format("20060102150405"))
}





