package ossinspector

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type PolicyStub struct {
	Policy Policy `yaml:"policy"`
}
type Owner struct {
	Age           string `yaml:"age,omitempty"`
	Repos         string `yaml:"repos,omitempty"`
	Followers     string `yaml:"followers,omitempty"`
	Contributions string `yaml:"contributions,omitempty"`
}
type Repo struct {
	Age          string `yaml:"age,omitempty"`
	Stars        string `yaml:"stars,omitempty"`
	Forks        string `yaml:"forks,omitempty"`
	Watchers     string `yaml:"watchers,omitempty"`
	Contributors string `yaml:"contributors,omitempty"`
}
type Commit struct {
	LastCommitAge string `yaml:"last_commit_age,omitempty"`
	Commits       string `yaml:"commits,omitempty"`
}
type Release struct {
	LastRelease string `yaml:"last_release,omitempty"`
}
type Policy struct {
	Owner   Owner   `yaml:"owner,omitempty"`
	Repo    Repo    `yaml:"repo"`
	Commit  Commit  `yaml:"commit"`
	Release Release `yaml:"release"`
}

func (p *Policy) ToString() string {
	bytes, _ := yaml.Marshal(p)
	return string(bytes)
}

func (p *Policy) isValid() (bool, error) {
	if !isAge(p.Owner.Age) {
		return false, errors.New("owner.age is not valid")
	}

	if !isAge(p.Repo.Age) {
		return false, errors.New("repo.age is not valid")
	}

	if !isAge(p.Commit.LastCommitAge) {
		return false, errors.New("commit.age is not valid")
	}

	if !isAge(p.Release.LastRelease) {
		return false, errors.New("release.last_release is not valid")
	}

	return true, nil
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
	policy := &pstub.Policy
	ok, err := policy.isValid()
	if !ok{
		return nil, err
	}

	return policy, nil
}
