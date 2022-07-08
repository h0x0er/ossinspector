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

	info, _ := ossinspector.FetchRepoInfo("step-security", "harden-runner")

	ok := ossinspector.Validate(conf, info)
	fmt.Printf("ok: %v\n", ok)
}
