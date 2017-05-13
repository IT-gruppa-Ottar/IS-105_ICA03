package iso

import "fmt"

//noinspection GoUnusedExportedFunction
func AddExtendedAscii() string{
	var ExtendedAscii []byte
	for i := 0x80; i <= 0xFF; i ++ {
		ExtendedAscii = append(ExtendedAscii, byte(i))
	}

	str := string(ExtendedAscii)
	//fmt.Printf("%q", str)
	return str
}

//noinspection GoUnusedExportedFunction
func IterateOverExtendedASCIIStringLiteral(sl string) {
	// Kode for Oppgave 2a
	fmt.Println("Extended Ascii")
	fmt.Println("Heksa/symbol/binært:")
	for i := 0; i < len(sl); i ++ {
		fmt.Printf("%X %c %b \n", sl[i], sl[i], sl[i])
	}
	fmt.Println()
}

//noinspection GoUnusedExportedFunction
func GreetingExtendedASCII() string{
	// Kode for Oppgave 2b
	greeting := []byte{34, 83, 97, 108, 117, 116, 44, 32, 231, 97, 32, 118, 97, 32, 176, 45, 41, 32, 128, 53, 48, 34}
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%c", greeting[i])
	}
    fmt.Println("")
	//String converter
	str := string(greeting)
	return str

}