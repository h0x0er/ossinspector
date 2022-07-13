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

var logger *ossinspector.Logger

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
	flag.StringVar(&options.respType, "resp", "", "format of response\nsupported formats are json, yaml; by default boolean value is returned.")
	flag.StringVar(&options.respType, "rt", "", "format of response")

	flag.Parse()

	logger = ossinspector.NewLogger(options.verbose)

	policy, err := ossinspector.NewPolicy(options.policy)
	if err != nil {
		logger.Printf("Unable to parse policy file\n%s\n", err)
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)

	}
	logger.Println("policy: ", policy)

	repo := strings.Split(options.repo, "/")
	logger.Printf("Performing policy check on %s/%s", repo[0], repo[1])
	info, err := ossinspector.FetchRepoInfo(repo[0], repo[1])
	if err != nil {
		logger.Printf("%v", err)
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)
	}

	ok, resp := ossinspector.Validate(policy, info)
	switch options.respType {
	case "json":
		fmt.Println(resp.ToJson())

	case "yaml":
		fmt.Println(resp.ToYaml())

	case "yml":
		fmt.Println(resp.ToYaml())

	default:
		fmt.Println(ok)

	}
}
