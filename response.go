package ossinspector

import (
	"encoding/json"
	"fmt"

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
	Age   bool `yaml:"age" json:"age"`
	Stars bool `yaml:"stars" json:"stars"`
	Forks bool `yaml:"forks" json:"forks"`
	// Watchers     bool `yaml:"watchers" json:"watchers"`

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

type PolicyScore struct {
	Score Score `json:"policy_score" yaml"policy_Score"`
}
type Score struct {
	Owner   string `json:"owner" yaml:"owner"`
	Repo    string `yaml:"repo" json:"repo"`
	Commit  string `yaml:"commit" json:"commit"`
	Release string `yaml:"release" json:"release"`
}

func (p *Response) GetScores() string {
	var pscore PolicyScore
	var score Score
	owner := getInt(p.PolicyResp.OwnerResp.Age) + getInt(p.PolicyResp.OwnerResp.Followers) + getInt(p.PolicyResp.OwnerResp.Repos)
	logger.Printf("owner_score: %v\n", owner)
	repo := getInt(p.PolicyResp.RepoResp.Age) + getInt(p.PolicyResp.RepoResp.Stars) + getInt(p.PolicyResp.RepoResp.Forks) + getInt(p.PolicyResp.RepoResp.Contributors)
	logger.Printf("repo-score: %v\n", repo)

	commit := getInt(p.PolicyResp.CommitResp.LastCommitAge)
	logger.Printf("commit_score: %v\n", commit)
	release := getInt(p.PolicyResp.ReleaseResp.LastRelease)
	logger.Printf("release_score: %v\n", release)

	score.Commit = fmt.Sprintf("%d/1", commit)
	score.Owner = fmt.Sprintf("%d/3", owner)
	score.Repo = fmt.Sprintf("%d/4", repo)
	score.Release = fmt.Sprintf("%d/1", release)
	pscore.Score = score
	bytes, _ := json.MarshalIndent(pscore, "", " ")
	return string(bytes)
}

func getInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
