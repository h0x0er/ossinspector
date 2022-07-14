package ossinspector

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type Response struct {
	PolicyResp PolicyResp `yaml:"policy_response" json:"policy_response"`
}
type OwnerResp struct {
	Age       bool `yaml:"age" json:"age"`
	Repos     bool `yaml:"repos" json:"repos"`
	Followers bool `yaml:"followers" json:"followers"`
	// Contributions bool `yaml:"contributions" json:"contributions"`
}
type RepoResp struct {
	Age          bool `yaml:"age" json:"age"`
	Stars        bool `yaml:"stars" json:"stars"`
	Forks        bool `yaml:"forks" json:"forks"`
	Watchers     bool `yaml:"watchers" json:"watchers"`
	Contributors bool `yaml:"contributors" json:"contributors"`
}
type CommitResp struct {
	LastCommitAge bool `yaml:"last_commit_age" json:"last_commit_age"`
	// Commits       bool `yaml:"commits" json:"commits"`
}
type ReleaseResp struct {
	LastRelease bool `yaml:"last_release" json:"last_release"`
}
type PolicyResp struct {
	OwnerResp   OwnerResp   `yaml:"owner" json:"owner"`
	RepoResp    RepoResp    `yaml:"repo" json:"repo"`
	CommitResp  CommitResp  `yaml:"commit" json:"commit"`
	ReleaseResp ReleaseResp `yaml:"release" json:"release"`
}

func (p *Response) ToYaml() string {
	bytes, _ := yaml.Marshal(p)
	return string(bytes)
}

func (p *Response) ToJson() string {
	bytes, _ := json.MarshalIndent(p, "", " ")
	return string(bytes)
}
