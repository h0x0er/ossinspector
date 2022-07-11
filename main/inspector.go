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
	policy, err := ossinspector.NewPolicy(pth)
	if err != nil {
		log.Fatalln("Unable to parse policy file\n", err)
	}

	info, _ := ossinspector.FetchRepoInfo("step-security", "harden-runner")

	ok := ossinspector.Validate(policy, info)
	fmt.Printf("policy followed: %v\n", ok)
}
