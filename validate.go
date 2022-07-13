package ossinspector

import (
	"log"
	"strings"
	"time"
)

func Validate(policy *Policy, repoInfo *RepoInfo) bool {
	response := new(Response)

	// TODO: verify repo rules
	ok := validateRepoRule(&policy.Repo, repoInfo, &response.PolicyResp.RepoResp)
	//TODO; verify author rule.
	validateAuthorRule(&policy.Owner, &repoInfo.OwnerInfo, &response.PolicyResp.AuthorResp)

	log.Println(response.ToJson())
	return ok

	// TODO: to verify commit rule
	// validateCommitRule()

}

func validateAuthorRule(ownerPolicy *Owner, ownerInfo *OwnerInfo, resp *AuthorResp) bool {

	resp.Age = checkExpr(ownerPolicy.Age, uint(ownerInfo.CreatedAt))
	resp.Followers = checkExpr(ownerPolicy.Followers, uint(ownerInfo.FollowersCount))
	resp.Repos = checkExpr(ownerPolicy.Repos, uint(ownerInfo.ReposCount))
	// NOTE:  still need to fetch contributions made by owner
	// resp.Contributions = checkExpr(ownerPolicy.Contributions, ownerInfo.Contributions)
	return resp.Age && resp.Followers && resp.Repos
}

func validateRepoRule(repoPolicy *Repo, repo *RepoInfo, resp *RepoResp) bool {

	resp.Stars = checkExpr(repoPolicy.Stars, repo.StaggersCount)
	resp.Watchers = checkExpr(repoPolicy.Watchers, repo.WatcherCount)
	resp.Forks = checkExpr(repoPolicy.Forks, repo.ForkCount)
	resp.Age = checkExpr(repoPolicy.Age, uint(repo.CreatedAt))

	log.Printf("stars_resp: %v\n", resp.Stars)
	log.Printf("watcher_resp: %v\n", resp.Watchers)
	log.Printf("forks_resp: %v\n", resp.Forks)
	log.Printf("age_resp: %v\n", resp.Age)

	// TODO: output needed to be improved
	return resp.Stars && resp.Watchers && resp.Age && resp.Forks

}

func checkExpr(checkString string, value uint) bool {
	checkString = strings.ReplaceAll(checkString, " ", "")
	resp := false
	ctype, pvalue := evaluate(checkString) // pvalue = policy value

	switch ctype {
	case LESSER_THAN:
		resp = (value < uint(pvalue))
	case GREATER_THAN:
		resp = (value >= uint(pvalue))

	case DAYS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := current - value
		days := (diff / (60 * 60 * 24))
		resp = (days < uint(pvalue))

	case DAYS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		days := (diff / (60 * 60 * 24))
		resp = (days >= uint(pvalue))

	case MONTHS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		months := (diff / (60 * 60 * 24 * 30))
		resp = (months < uint(pvalue))

	case MONTHS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		months := (diff / (60 * 60 * 24 * 30))
		resp = (months >= uint(pvalue))

	case YEARS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		years := (diff / (60 * 60 * 24 * 30 * 12))
		resp = (years < uint(pvalue))

	case YEARS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		years := (diff / (60 * 60 * 24 * 30 * 12))
		resp = (years >= uint(pvalue))

	}
	return resp
}

func checkBool(fromPolicy bool, fromRepo bool) bool {
	return fromPolicy == fromRepo
}
