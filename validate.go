package ossinspector

import (
	"strings"
	"time"
)

func Validate(policy *Policy, repoInfo *RepoInfo) (bool, *Response) {
	response := new(Response)

	// TODO: verify repo rules
	ok1 := validateRepoRule(&policy.Repo, repoInfo, &response.PolicyResp.RepoResp)
	//TODO; verify author rule.
	ok2 := validateAuthorRule(&policy.Owner, &repoInfo.OwnerInfo, &response.PolicyResp.AuthorResp)

	// TODO: to verify commit rule
	// validateCommitRule()
	return (ok1 && ok2), response

}

func validateAuthorRule(ownerPolicy *Owner, ownerInfo *OwnerInfo, resp *AuthorResp) bool {
	var val int
	resp.Age, val = checkExpr(ownerPolicy.Age, uint(ownerInfo.CreatedAt))
	logger.Printf("Owner Age: %v\n", val)
	resp.Followers, val = checkExpr(ownerPolicy.Followers, uint(ownerInfo.FollowersCount))
	logger.Printf("Owner Followers: %v\n", val)
	resp.Repos, val = checkExpr(ownerPolicy.Repos, uint(ownerInfo.ReposCount))
	logger.Printf("Owner Repos: %v\n", val)
	// NOTE:  still need to fetch contributions made by owner
	// resp.Contrib:u ons = checkExpr(ownerPolicy.Contributions, ownerInfo.Contributions)
	return resp.Age && resp.Followers && resp.Repos
}

func validateRepoRule(repoPolicy *Repo, repo *RepoInfo, resp *RepoResp) bool {
	var val int
	resp.Stars, val = checkExpr(repoPolicy.Stars, repo.StaggersCount)
	logger.Printf("Repo stars: %v\n", val)
	resp.Watchers, val = checkExpr(repoPolicy.Watchers, repo.WatcherCount)
	logger.Printf("Repo Watchers: %v\n", val)
	resp.Forks, val = checkExpr(repoPolicy.Forks, repo.ForkCount)
	logger.Printf("Repo Forks: %v\n", val)
	resp.Age, val = checkExpr(repoPolicy.Age, uint(repo.CreatedAt))
	logger.Printf("Repo Age: %v\n", val)
	resp.Contributors, val = checkExpr(repoPolicy.Contributors, repo.ContributorsCount)
	logger.Printf("Repo Contributors: %v\n", val)

	// TODO: output needed to be improved
	return resp.Stars && resp.Watchers && resp.Age && resp.Forks

}

func checkExpr(checkString string, value uint) (bool, int) {
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
		value = days

	case DAYS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		days := (diff / (60 * 60 * 24))
		resp = (days >= uint(pvalue))
		value = days

	case MONTHS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		months := (diff / (60 * 60 * 24 * 30))
		resp = (months < uint(pvalue))
		value = months

	case MONTHS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		months := (diff / (60 * 60 * 24 * 30))
		resp = (months >= uint(pvalue))
		value = months

	case YEARS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		years := (diff / (60 * 60 * 24 * 30 * 12))
		resp = (years < uint(pvalue))
		value = years

	case YEARS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		years := (diff / (60 * 60 * 24 * 30 * 12))
		resp = (years >= uint(pvalue))
		value = years

	}
	return resp, int(value)
}

func checkBool(fromPolicy bool, fromRepo bool) bool {
	return fromPolicy == fromRepo
}
