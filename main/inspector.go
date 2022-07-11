package main

import (
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
		log.Fatalf("Unable to parse policy file\n%s\n", err)
	}
	log.Println("policy: ", policy)
	info, _ := ossinspector.FetchRepoInfo("step-security", "harden-runner")

	ok := ossinspector.Validate(policy, info)
	log.Printf("policy followed: %v\n", ok)
}
