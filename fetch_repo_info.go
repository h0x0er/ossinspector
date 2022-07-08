package ossinspector

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// RepoInfo info about the target repo
type RepoInfo struct {
	CreatedAt     int64 // timestamp of creation
	LastPushedAt  int64
	LastUpdatedAt int64
	IsArchived    bool
	IsDisabled    bool
	IsForked      bool
	StaggersCount uint
	WatcherCount  uint
	ForkCount     uint

	OwnerInfo  OwnerInfo
	CommitInfo CommitInfo
}

// OwnerInfo info about actual creator of the repo.
type OwnerInfo struct {
	Owner          string
	CreatedAt      int64
	UpdatedAt      int64
	ReposCount     int // count of public repos
	FollowersCount int
}

// CommitInfo info about commits
type CommitInfo struct{
	// TODO: add commit information
}

func FetchRepoInfo(owner, repo string) (*RepoInfo, error) {
	client := getClient()
	repos, resp, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		log.Fatalf("Unable to fetch package: %s/%s\n %v", owner, repo, err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Fatalf("it seems %s/%s doesn't exists", owner, repo)
		return nil, errors.New("repo doesn't exists")
	}

	owner_info, _, err := client.Users.Get(context.Background(), owner)
	if err != nil {
		log.Fatalf("Unable to fetch owner %s", owner)
		return nil, err
	}

	var repoInfo RepoInfo

	repoInfo.CreatedAt = repos.GetCreatedAt().Unix()
	repoInfo.LastUpdatedAt = repos.GetUpdatedAt().Unix()
	repoInfo.LastPushedAt = repos.PushedAt.Unix()

	repoInfo.IsArchived = repos.GetArchived()
	repoInfo.IsDisabled = repos.GetDisabled()
	repoInfo.IsForked = repos.GetFork()

	repoInfo.ForkCount = uint(repos.GetForksCount())
	repoInfo.StaggersCount = uint(repos.GetStargazersCount())
	repoInfo.WatcherCount = uint(repos.GetStargazersCount())

	repoInfo.OwnerInfo.CreatedAt = owner_info.GetCreatedAt().Unix()
	repoInfo.OwnerInfo.UpdatedAt = owner_info.GetUpdatedAt().Unix()
	repoInfo.OwnerInfo.ReposCount = owner_info.GetPublicRepos()

	// TODO: Add more owner related info
	//client.Repositories.

	return &repoInfo, nil
}

func getClient() *github.Client {
	token := os.Getenv("gh_token")
	if token != "" {
		log.Printf("oauth token found")
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(context.Background(), ts)
		client := github.NewClient(tc)
		return client
	}
	client := github.NewClient(nil)
	return client

}
