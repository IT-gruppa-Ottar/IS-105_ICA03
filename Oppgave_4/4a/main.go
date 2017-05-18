package main

import (
	"./unicode"
	"fmt"
)

func main() {

	//oversetter på japansk
	japanese := unicode.Translate("nord og sør", "jp")
	fmt.Printf("%q\n", japanese)

	//oversetter nord og sør på islandsk
	icelandic := unicode.Translate("nord og sør", "is")
	fmt.Printf("%q\n", icelandic)

	//prøver å oversette nord og sør på engelsk
	english := unicode.Translate("nord og sør", "eng")
	fmt.Printf("%q\n", english)
}
