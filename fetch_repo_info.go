package ossinspector

import (
	"context"
	"errors"
	"os"
	"sync"

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
	TotalCommits int64 // TODO: still need to fetch total commits
}

// ReleaseInfo
type ReleaseInfo struct {
	LastReleaseAt int64
	LastCreatedAt int64
}

func FetchRepoInfo(owner, repo string) (*RepoInfo, error) {
	repoInfo := new(RepoInfo)
	var err error
	logger.Println("*")
	err = addRepoInfo(owner, repo, repoInfo)
	if err != nil {
		return nil, err
	}
	logger.Println("**")
	_ = addOwnerInfo(owner, repoInfo)
	// if err != nil {
	// 	return nil, err
	// }
	logger.Println("***")
	_ = addReleaseInfo(owner, repo, &repoInfo.ReleaseInfo)
	// if err != nil {
	// 	return nil, err
	// }
	logger.Println("****")
	_ = addCommitInfo(owner, repo, &repoInfo.CommitInfo)
	// if err != nil {
	// 	return nil, err
	// }
	logger.Printf("Info fetched successfully...")
	return repoInfo, nil
}
func addRepoInfo(owner, repo string, repoInfo *RepoInfo) error {
	client.Lock()
	repos, resp, err := client.git.Repositories.Get(context.Background(), owner, repo)
	client.Unlock()
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
	// repoInfo.WatcherCount = uint(repos.GetStargazersCount())
	repoInfo.ContributorsCount = uint(getContributorsCount(owner, repo))

	return nil
}
func addOwnerInfo(owner string, repoInfo *RepoInfo) error {
	client.Lock()
	owner_info, _, err := client.git.Users.Get(context.Background(), owner)
	client.Unlock()
	if err != nil {
		logger.Printf("Unable to fetch owner %s", owner)
		return err
	}
	repoInfo.OwnerInfo.CreatedAt = owner_info.GetCreatedAt().Unix()
	repoInfo.OwnerInfo.UpdatedAt = owner_info.GetUpdatedAt().Unix()
	repoInfo.OwnerInfo.ReposCount = owner_info.GetPublicRepos()
	repoInfo.OwnerInfo.FollowersCount = getFollowerCounts(owner)
	// TODO: need to add total contribution counts of owner on github
	return nil
}

func addReleaseInfo(owner, repo string, releaseInfo *ReleaseInfo) error {
	client.Lock()
	release, resp, err := client.git.Repositories.GetLatestRelease(context.Background(), owner, repo)
	client.Unlock()
	if err != nil {
		releaseInfo.LastCreatedAt = 0
		releaseInfo.LastReleaseAt = 0
		logger.Printf("Unable to fetch release info: %s/%s\n %v", owner, repo, err)
		return err
	}

	if resp.StatusCode != 200 {
		logger.Printf("it seems relase info for %s/%s doesn't exists", owner, repo)
		return errors.New("repo doesn't exists")
	}
	releaseInfo.LastCreatedAt = release.GetCreatedAt().Unix()
	releaseInfo.LastReleaseAt = release.GetPublishedAt().Unix()
	return nil
}
func addCommitInfo(owner, repo string, commitInfo *CommitInfo) error {
	// NOTE: fetches information related commit
	client.Lock()
	commits, resp, err := client.git.Repositories.ListCommits(context.Background(), owner, repo, nil)
	client.Unlock()
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		logger.Printf("unable to fetch commits for %s/%s", owner, repo)
		return errors.New("unable to fetch commits")
	}
	for _, rc := range commits {
		// we are only interested in first commit
		commitInfo.LastCommitAt = rc.GetCommit().Author.Date.Unix()
		logger.Printf("commitInfo.SHA: %v\n", rc.GetSHA())
		break
	}

	return nil
}

type ClientT struct {
	sync.Mutex
	git *github.Client
}

var client *ClientT

func MakeClient() *ClientT {
	if client == nil {
		logger.Println("creating universal client")
		token := os.Getenv("gh_token")
		if token != "" {
			logger.Printf("gh_token environment variable found")
			ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
			tc := oauth2.NewClient(context.Background(), ts)
			git := github.NewClient(tc)

			client = new(ClientT)
			client.git = git
			return client
		} else {
			logger.Printf("gh_token environment variable not found. Setup gh_token to evade rate limiting")
			git := github.NewClient(nil)
			client = new(ClientT)
			client.git = git
			return client
		}

	}
	return client
}

func getFollowerCounts(user string) int {
	count := 0

	page := 1
	for {
		if count > 500 {
			logger.Println("[!] Owner's follower are more than 500")
			count = 99999
			break
		}
		resp, _, err := client.git.Users.ListFollowers(context.Background(), user, &github.ListOptions{Page: page, PerPage: 100})
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

func getContributorsCount(owner, repo string) int {
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
		client.Lock()
		resp, _, err := client.git.Repositories.ListContributors(context.Background(), owner, repo, options)
		client.Unlock()
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
