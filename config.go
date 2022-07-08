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
	Age           int `yaml:"age"`
	Repos         int `yaml:"repos"`
	Followers     int `yaml:"followers"`
	Contributions int `yaml:"contributions"`
}

// TODO: Add repo age attribute
type RepoRule struct {
	Stars int `yaml:"stars"`
	// Age          int `yaml:age`
	Forks        int `yaml:"forks"`
	Watchers     int `yaml:"watchers"`
	Contributors int `yaml:"contributors"`
	Commits      int `yaml:"commits"`
	LastRelease  int `yaml:"last_release"`
}
type CommitRule struct {
	LastCommitAge int `yaml:"last_commit_age"`
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
