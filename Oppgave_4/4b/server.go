// Filen for å eksperimentere med webserver i ICA03 (IS-105)
// Starte server med go run server.go
// Du kan ha tilgang til server fra nettleser med http://localhost:3000
package main

import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	//getTime()
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

/**
Her kan man bl. a. skrive data som skal sendes til nettleser
Neste linje kan brukes for å endre koding av sekvensen som sendes til nettleser
Her skriver man data som er respons til brukeren som har skrevet
http://localhost:3000 i sin nettleser
 */
func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<font color=\"black\">\xE2\x8F\xB0\t</font><br/>"))

	klokke := getTime()
	w.Write([]byte(klokke))
}

func getTime()[]byte{
	//Henter tid
	t := time.Now()

	//Konverterer til byte
	test := fmt.Sprintf("%s", t)
	temp := []byte (test)
	
	return temp
}