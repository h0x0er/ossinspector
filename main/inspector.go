package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/h0x0er/ossinspector"
)

type Options struct {
	policy   string
	verbose  bool
	respType string
	repo     string
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("oss-inspector: ")
}

func main() {
	options := new(Options)
	flag.StringVar(&options.policy, "policy", "policy.yml", "location of policy.yml file")
	flag.StringVar(&options.policy, "p", "policy.yml", "location of policy.yml file")
	flag.BoolVar(&options.verbose, "verbose", false, "enable verbose response")
	flag.BoolVar(&options.verbose, "v", false, "enable verbose response")
	flag.StringVar(&options.repo, "repo", "facebook/react", "repository to inspect")
	flag.StringVar(&options.repo, "r", "facebook/react", "repository to inspect")

	flag.Parse()
	policy, err := ossinspector.NewPolicy(options.policy)
	if err != nil {
		log.Fatalf("Unable to parse policy file\n%s\n", err)
	}
	log.Println("policy: ", policy)
	repo := strings.Split(options.repo, "/")
	info, _ := ossinspector.FetchRepoInfo(repo[0], repo[1])

	ok := ossinspector.Validate(policy, info)
	fmt.Printf("policy followed: %v\n", ok)
}
