package ascii

import "fmt"


const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

//noinspection GoUnusedExportedFunction
func ReturnAscii()string{
	return ascii
}

//noinspection GoUnusedExportedFunction
func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 1		
	fmt.Println("Hexa/symbol/binært:")
	for i := 0; i < len(sl); i ++ {
		fmt.Printf ("%X %c %b\n", sl[i], sl[i], sl[i])
	}
}

//noinspection GoUnusedExportedFunction
func GreetingASCII() string{
	//"Hello :-)" i hexa
	fmt.Println()
	greeting := []byte {34, 72, 101, 108, 108, 111, 32, 58, 45, 41, 34}
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%c", greeting[i])
	}

	fmt.Println("")
	//String converter
	str := string(greeting)
	return str
}