package ossinspector

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type PolicyStub struct {
	Policy Policy `yaml:"policy"`
}
type Author struct {
	Age           string `yaml:"age"`
	Repos         string `yaml:"repos"`
	Followers     string `yaml:"followers"`
	Contributions string `yaml:"contributions"`
}
type Repo struct {
	Age          string `yaml:"age"`
	Stars        string `yaml:"stars"`
	Forks        string `yaml:"forks"`
	Watchers     string `yaml:"watchers"`
	Contributors string `yaml:"contributors"`
	Commits      string `yaml:"commits"`
	LastRelease  string `yaml:"last_release"`
}
type Commit struct {
	LastCommitAge string `yaml:"last_commit_age"`
}
type Policy struct {
	Author Author `yaml:"author"`
	Repo   Repo   `yaml:"repo"`
	Commit Commit `yaml:"commit"`
}

func NewPolicy(path string) (*Policy, error) {
	var pstub PolicyStub
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, &pstub)
	if err != nil {
		return nil, err
	}

	return &pstub.Policy, nil
}
