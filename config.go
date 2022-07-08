package ossinspector

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TrustRules PackageTrustRule `yaml:"trust_rules"`
}

type PackageTrustRule struct {
	AuthorRules AuthorRule `yaml:"author_rule"`
	RepoRules   RepoRule   `yaml:"repo_rule"`
	CommitRules CommitRule `yaml:"commit_rules"`
}
type AuthorRule struct {
	Age           string `yaml:"age"`
	Repos         string `yaml:"repos"`
	Followers     string `yaml:"followers"`
	Contributions string `yaml:"contributions"`
}

// TODO: Add repo age attribute
type RepoRule struct {
	Stars string `yaml:"stars"`
	// Age          string `yaml:age`
	Forks        string `yaml:"forks"`
	Watchers     string `yaml:"watchers"`
	Contributors string `yaml:"contributors"`
	Commits      string `yaml:"commits"`
	LastRelease  string `yaml:"last_release"`
}
type CommitRule struct {
	LastCommitAge string `yaml:"last_commit_age"`
}

func NewConfig(path string) (*Config, error) {
	var config Config
	bytes, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
