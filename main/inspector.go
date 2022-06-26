package main

import (
	"fmt"
	"log"
	"os"

	"github.com/h0x0er/ossinspector"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("oss-inspector: ")
}

func main() {

	pth := os.Args[1]
	conf, err := ossinspector.NewConfig(pth)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("", conf)
	ossinspector.FetchRepoInfo("step-security", "secure-workflows")
}
