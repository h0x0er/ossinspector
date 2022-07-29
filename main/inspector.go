package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/h0x0er/ossinspector"
)

type Options struct {
	policy      string
	verbose     bool
	respType    string
	repo        string
	projectType string
}

var logger *ossinspector.Logger

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("oss-inspector: ")
}

func main() {
	options := new(Options)

	flag.StringVar(&options.policy, "policy", "policy.yml", "location of policy.yml file (default is current directory)")
	flag.StringVar(&options.policy, "p", "policy.yml", "location of policy.yml file (default is current directory)")
	flag.BoolVar(&options.verbose, "verbose", false, "enable verbose response")
	flag.BoolVar(&options.verbose, "v", false, "enable verbose response")
	flag.StringVar(&options.repo, "repo", "facebook/react", "repository to inspect")
	flag.StringVar(&options.repo, "r", "facebook/react", "repository to inspect")
	flag.StringVar(&options.respType, "resp", "", "format of response\nsupported formats are json, yaml, score ; by default boolean value is returned.")
	flag.StringVar(&options.respType, "rt", "", "format of response")
	flag.StringVar(&options.projectType, "t", "", "type of project (node, python, go etc..)")

	flag.Parse()

	wg := new(sync.WaitGroup)
	logger = ossinspector.NewLogger(options.verbose)
	ossinspector.MakeClient()

	policy, err := ossinspector.NewPolicy(options.policy)
	if err != nil {
		logger.Printf("Unable to parse policy file\n%s\n", err)
		fmt.Println("policy_file:", err)
		os.Exit(0)
	}
	logger.Println("Policy: ", policy)

	switch options.projectType {
	case "node":
		logger.Println("[!] Running in project mode")
		logger.Printf("[!] Project Type:  %s", options.projectType)
		repos, err := ossinspector.GetNodeRepos("./package.json")
		if err != nil {
			logger.Fatalf("%s", err)
			fmt.Printf("ossinspector: unable to perform action \n%v\n", err)
			os.Exit(-1)
		}
		logger.Printf("repos: %v\n", repos)
		for _, repo := range repos {
			wg.Add(1)
			go func(repo string, wg *sync.WaitGroup) {
				defer wg.Done()
				temp := strings.Split(repo, "/")
				logger.Printf("[!] Policy Check [%s/%s]", temp[0], temp[1])
				info, err := ossinspector.FetchRepoInfo(temp[0], temp[1])
				if err != nil {
					logger.Printf("unable to fetch info for %s", repo)
				} else {
					ok, resp := ossinspector.Validate(policy, info)
					fmt.Printf("[%s]\n", repo)
					handleReponse(options.respType, ok, resp)

				}
			}(repo, wg)
			wg.Wait()

		}

	default:
		logger.Println("[!] Running in repo mode...")
		repo := strings.Split(options.repo, "/")
		logger.Printf("[*] Policy Check [%s/%s]", repo[0], repo[1])
		info, err := ossinspector.FetchRepoInfo(repo[0], repo[1])
		if err != nil {
			logger.Printf("%v", err)
			fmt.Printf("err: %v\n", err)
			os.Exit(0)
		}

		ok, resp := ossinspector.Validate(policy, info)
		handleReponse(options.respType, ok, resp)

	}

}

func handleReponse(respType string, respOk bool, resp *ossinspector.Response) {
	switch respType {
	case "json":
		fmt.Println(resp.ToJson())

	case "yaml":
		fmt.Println(resp.ToYaml())

	case "yml":
		fmt.Println(resp.ToYaml())

	case "score":
		fmt.Println(resp.GetScores())

	default:
		fmt.Println(respOk)

	}
}
