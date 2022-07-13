package ossinspector

import (
	"context"
	"errors"
	"fmt"
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

	OwnerInfo   OwnerInfo
	CommitInfo  CommitInfo
	ReleaseInfo ReleaseInfo
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
type CommitInfo struct {
	// TODO: add commit information
	LastCommitAt int64
	TotalCommits int64
}

// ReleaseInfo
type ReleaseInfo struct {
	LastReleaseAt int64
}

func FetchRepoInfo(owner, repo string) (*RepoInfo, error) {
	client := getClient()
	repoInfo := new(RepoInfo)

	addRepoInfo(client, owner, repo, repoInfo)
	addOwnerInfo(client, owner, repoInfo)

	return repoInfo, nil
}
func addRepoInfo(client *github.Client, owner, repo string, repoInfo *RepoInfo) error {

	repos, resp, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		log.Printf("Unable to fetch package: %s/%s\n %v", owner, repo, err)
		return err
	}

	if resp.StatusCode != 200 {
		log.Printf("it seems %s/%s doesn't exists", owner, repo)
		return errors.New("repo doesn't exists")
	}
	// NOTE: unable to fetch total contributor
	// contrib, _, err := client.Repositories.ListContributors(context.Background(), owner, repo, nil)

	repoInfo.CreatedAt = repos.GetCreatedAt().Unix()
	repoInfo.LastUpdatedAt = repos.GetUpdatedAt().Unix()
	repoInfo.LastPushedAt = repos.PushedAt.Unix()

	repoInfo.IsArchived = repos.GetArchived()
	repoInfo.IsDisabled = repos.GetDisabled()
	repoInfo.IsForked = repos.GetFork()

	repoInfo.ForkCount = uint(repos.GetForksCount())
	repoInfo.StaggersCount = uint(repos.GetStargazersCount())
	repoInfo.WatcherCount = uint(repos.GetStargazersCount())

	return nil
}
func addOwnerInfo(client *github.Client, owner string, repoInfo *RepoInfo) error {
	owner_info, _, err := client.Users.Get(context.Background(), owner)
	if err != nil {
		log.Printf("Unable to fetch owner %s", owner)
		return err
	}
	repoInfo.OwnerInfo.CreatedAt = owner_info.GetCreatedAt().Unix()
	repoInfo.OwnerInfo.UpdatedAt = owner_info.GetUpdatedAt().Unix()
	repoInfo.OwnerInfo.ReposCount = owner_info.GetPublicRepos()
	//TODO: add more owner info
	repoInfo.OwnerInfo.FollowersCount = getFollowerCounts(client, owner)
	fmt.Printf("owner_info.GetFollowers(): %v\n", repoInfo.OwnerInfo.FollowersCount)
	return nil
}

func addReleaseInfo(client *github.Client, owner, repo string, repoInfo *RepoInfo) error {

	_, resp, err := client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		log.Printf("Unable to fetch package: %s/%s\n %v", owner, repo, err)
		return err
	}

	if resp.StatusCode != 200 {
		log.Printf("it seems %s/%s doesn't exists", owner, repo)
		return errors.New("repo doesn't exists")
	}

	return nil
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

func getFollowerCounts(client *github.Client, user string) int {
	count := 0

	page := 1
	for {
		resp, _, err := client.Users.ListFollowers(context.Background(), user, &github.ListOptions{Page: page, PerPage: 100})
		if err == nil {
			l := len(resp)
			count += l
			if l < 100 {
				break
			}
			page += 1
		}
	}

	return count
}
