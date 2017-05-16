package treasure

import (
	"fmt"

)

//noinspection GoUnusedExportedFunction
func Temp(){
	fmt.Printf("%c", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	//Resultatet av %s er: 	"??=? ⌘"
	//Resultatet av %q er: 	"\xbd\xb2=\xbc ⌘"
	//Resultatet av %+q er: "\xbd\xb2=\xbc \u2318"
	//Resultatet av %c er: 	"[½ ² = ¼ â ]"


	//Forventer:		"½?=? ⌘"

}


// Kode for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen

//noinspection GoUnusedExportedFunction
func PrintTreasureUTF8(treasure_string string){
	//mangler æ, ø, å

	fmt.Printf("%s", treasure_string)


}