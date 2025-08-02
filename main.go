package main

import (
	"fmt"
	"os"
)

func main() {
	err := initGlobals()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) != 2 {
		fmt.Println(`Usage:

generate CA files: CA.crt, CA.key
    ./generate_pki CA

genarate cert files for server with CA files: Server.crt, Server.key
    ./generate_pki CertServer`)
	} else {
		switch os.Args[1] {
		case "CA":
			mainConfig.CA.Generate()
		case "CertServer":
			mainConfig.CertServer.Generate()
		default:
			fmt.Printf("unknown sub-command: %v\n", os.Args[1])
		}
	}
}
