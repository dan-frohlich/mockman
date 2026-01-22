package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	sourceParam      string
	destinationParam string
	packageParam     string
	filterParam      string
)

func main() {
	flag.StringVar(&sourceParam, "s", "", "Source file or directory")
	flag.StringVar(&destinationParam, "d", "", "Destination file or directory")
	flag.StringVar(&packageParam, "p", "", "destination package pattern")
	flag.StringVar(&filterParam, "f", "", "Optional filter pattern")
	flag.Parse()

	// fmt.Printf("     source: %s\n", sourceParam)
	// fmt.Printf("destination: %s\n", destinationParam)
	// fmt.Printf("    package: %s\n", packageParam)
	// fmt.Printf("     filter: %s\n", filterParam)

	z, err := findInterfaces(sourceParam, filterParam)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// fmt.Println("Found interfaces:")
	// for _, iface := range z {
	// 	fmt.Println(" -", iface.Describe())
	// }

	for _, iface := range z {
		var code string
		code, err = mockInterface(packageParam, iface)
		if err != nil {
			fmt.Printf("Error generating mock for interface %s: %v\n", iface.Name, err)
			continue
		}

		err := os.WriteFile(destinationParam, []byte(code), 0644)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}
