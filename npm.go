package ossinspector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetNodeRepos(path string) ([]string, error) {
	// return repo of all deps

	var packageJson map[string]map[string]string

	var output []string
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &packageJson)

	if deps, ok := packageJson["dependencies"]; ok {
		for dep := range deps {
			link, err := getRepoLink(dep)
			if err == nil {
				repo, err := getRepo(link)
				if err == nil {
					logger.Printf("%s -->  %v\n", dep, link)
					output = append(output, repo)
				}
			}
		}
	}

	if deps, ok := packageJson["peerDependencies"]; ok {
		for dep := range deps {
			link, err := getRepoLink(dep)
			if err == nil {
				repo, err := getRepo(link)
				if err == nil {
					logger.Printf("%s -->  %v\n", dep, link)
					output = append(output, repo)
				}
			}
		}
	}

	if deps, ok := packageJson["devDependencies"]; ok {
		for dep := range deps {
			link, err := getRepoLink(dep)
			if err == nil {
				repo, err := getRepo(link)
				if err == nil {
					logger.Printf("%s -->  %v\n", dep, link)
					output = append(output, repo)
				}
			}
		}
	}

	return output, nil
}

type registry struct {
	Repository repository `json:"repository"`
}

type repository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

func getRepo(link string) (string, error) {
	if link == "" {
		return "", errors.New("empty link")
	}
	parsed, err := url.Parse(link)
	if err != nil {
		logger.Printf("[!] error parsing %v", link)
		return "", err
	}

	return strings.Split(parsed.Path, ".git")[0][1:], nil
}
func getRepoLink(dep string) (string, error) {
	registryUrl := fmt.Sprintf("https://registry.npmjs.com/%s", dep)
	client := http.DefaultClient

	resp, err := client.Get(registryUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var reg registry
	err = json.Unmarshal(bytes, &reg)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(reg.Repository.Url, "git+") {
		return strings.Split(reg.Repository.Url, "+")[1], nil
	}
	return reg.Repository.Url, nil

}
