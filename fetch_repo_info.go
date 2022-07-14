package ossinspector

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// RepoInfo info about the target repo
type RepoInfo struct {
	CreatedAt         int64 // timestamp of creation
	LastPushedAt      int64
	LastUpdatedAt     int64
	IsArchived        bool
	IsDisabled        bool
	IsForked          bool
	StaggersCount     uint
	WatcherCount      uint
	ForkCount         uint
	ContributorsCount uint

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
	var err error
	err = addRepoInfo(client, owner, repo, repoInfo)
	if err != nil {
		return nil, err
	}
	err = addOwnerInfo(client, owner, repoInfo)
	if err != nil {
		return nil, err
	}

	return repoInfo, nil
}
func addRepoInfo(client *github.Client, owner, repo string, repoInfo *RepoInfo) error {

	repos, resp, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		logger.Printf("Unable to fetch package: %s/%s\n %v", owner, repo, err)
		return err
	}

	if resp.StatusCode != 200 {
		logger.Printf("it seems %s/%s doesn't exists", owner, repo)
		return errors.New("repo doesn't exists")
	}

	repoInfo.CreatedAt = repos.GetCreatedAt().Unix()
	repoInfo.LastUpdatedAt = repos.GetUpdatedAt().Unix()
	repoInfo.LastPushedAt = repos.PushedAt.Unix()

	repoInfo.IsArchived = repos.GetArchived()
	repoInfo.IsDisabled = repos.GetDisabled()
	repoInfo.IsForked = repos.GetFork()

	repoInfo.ForkCount = uint(repos.GetForksCount())
	repoInfo.StaggersCount = uint(repos.GetStargazersCount())
	repoInfo.WatcherCount = uint(repos.GetStargazersCount())
	repoInfo.ContributorsCount = uint(getContributorsCount(client, owner, repo))

	return nil
}
func addOwnerInfo(client *github.Client, owner string, repoInfo *RepoInfo) error {
	owner_info, _, err := client.Users.Get(context.Background(), owner)
	if err != nil {
		logger.Printf("Unable to fetch owner %s", owner)
		return err
	}
	repoInfo.OwnerInfo.CreatedAt = owner_info.GetCreatedAt().Unix()
	repoInfo.OwnerInfo.UpdatedAt = owner_info.GetUpdatedAt().Unix()
	repoInfo.OwnerInfo.ReposCount = owner_info.GetPublicRepos()
	repoInfo.OwnerInfo.FollowersCount = getFollowerCounts(client, owner)
	// TODO: need to add total contribution counts of owner on github
	return nil
}

func addReleaseInfo(client *github.Client, owner, repo string, repoInfo *RepoInfo) error {

	_, resp, err := client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		logger.Printf("Unable to fetch package: %s/%s\n %v", owner, repo, err)
		return err
	}

	if resp.StatusCode != 200 {
		logger.Printf("it seems %s/%s doesn't exists", owner, repo)
		return errors.New("repo doesn't exists")
	}

	return nil
}
func addCommitInfo(client *github.Client, owner, repo string, commitInfo *CommitInfo) error {
	// NOTE: fetches information related commit
	client.Repositories.
	return nil
}

func getClient() *github.Client {
	token := os.Getenv("gh_token")
	if token != "" {
		logger.Printf("gh_token environment variable found")
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(context.Background(), ts)
		client := github.NewClient(tc)
		return client
	}
	logger.Printf("gh_token environment variable not found. Setup gh_token to evade rate limiting")
	client := github.NewClient(nil)
	return client

}

func getFollowerCounts(client *github.Client, user string) int {
	count := 0

	page := 1
	for {
		if count > 500 {
			count = 99999
			break
		}
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

func getContributorsCount(client *github.Client, owner, repo string) int {
	count := 0
	page := 1
	for {
		options := new(github.ListContributorsOptions)
		options.Anon = "false"
		options.Page = page
		options.PerPage = 100
		if count > 500 {
			logger.Println("contributors are more than 500")
			// NOTE
			// if number or contributors are more than 500
			// then it is definite the repo is trustable
			count = 99999 // setting to maximum count; so that check can be passed
			break
		}
		resp, _, err := client.Repositories.ListContributors(context.Background(), owner, repo, options)

		if err == nil {
			c := len(resp)
			count += c
			if c < 100 {
				break
			}
			page += 1
		}

	}

	return count

}
