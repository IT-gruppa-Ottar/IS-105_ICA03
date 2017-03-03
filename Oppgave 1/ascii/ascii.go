package ascii

import (
	"fmt"
)

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

//noinspection GoUnusedExportedFunction
func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 1
	for i := 0; i < len(sl); i ++ {
		fmt.Println("Hexa/symbol/binært:")
		fmt.Printf ("%X %s %b\n", sl, sl, sl[i])
	}
}

//noinspection GoUnusedExportedFunction
func GreetingASCII() {
	//"Hello :-)" i hexa
	s := "\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"
	fmt.Println()
	fmt.Println(s)
}