package ossinspector

import (
	"strings"
	"time"
)

func Validate(policy *Policy, repoInfo *RepoInfo) (bool, *Response) {
	response := new(Response)

	ok1 := validateRepoRule(&policy.Repo, repoInfo, &response.PolicyResp.RepoResp)

	ok2 := validateOwnerRule(&policy.Owner, &repoInfo.OwnerInfo, &response.PolicyResp.OwnerResp)

	ok3 := validateCommitRule(&policy.Commit, &repoInfo.CommitInfo, &response.PolicyResp.CommitResp)

	ok4 := validateReleaseRule(&policy.Release, &repoInfo.ReleaseInfo, &response.PolicyResp.ReleaseResp)

	return (ok1 && ok2 && ok3 && ok4), response

}

func validateOwnerRule(ownerPolicy *Owner, ownerInfo *OwnerInfo, ownerResp *OwnerResp) bool {
	var val int
	ownerResp.Age, val = checkExpr(ownerPolicy.Age, uint(ownerInfo.CreatedAt))
	logger.Printf("ownerInfo.Age: %v\n", val)
	ownerResp.Followers, val = checkExpr(ownerPolicy.Followers, uint(ownerInfo.FollowersCount))
	logger.Printf("ownerInfo.Followers: %v\n", val)
	ownerResp.Repos, val = checkExpr(ownerPolicy.Repos, uint(ownerInfo.ReposCount))
	logger.Printf("ownerInfo.Repos: %v\n", val)
	// NOTE:  still need to fetch contributions made by owner
	// resp.Contrib:u ons = checkExpr(ownerPolicy.Contributions, ownerInfo.Contributions)
	return ownerResp.Age && ownerResp.Followers && ownerResp.Repos // NOTE: make sure all rules were followed
}

func validateRepoRule(repoPolicy *Repo, repo *RepoInfo, repoResp *RepoResp) bool {
	var val int
	repoResp.Stars, val = checkExpr(repoPolicy.Stars, repo.StaggersCount)
	logger.Printf("repoResp.Stars: %v\n", val)
	repoResp.Watchers, val = checkExpr(repoPolicy.Watchers, repo.WatcherCount)
	logger.Printf("repoResp.Watchers: %v\n", val)
	repoResp.Forks, val = checkExpr(repoPolicy.Forks, repo.ForkCount)
	logger.Printf("repoResp.Forks: %v\n", val)
	repoResp.Age, val = checkExpr(repoPolicy.Age, uint(repo.CreatedAt))
	logger.Printf("repoResp.Age: %v\n", val)
	repoResp.Contributors, val = checkExpr(repoPolicy.Contributors, repo.ContributorsCount)
	logger.Printf("repoResp.Contributors: %v\n", val)
	// NOTE: combined repoResponse to check; whether all rules were followed or not
	return repoResp.Stars && repoResp.Watchers && repoResp.Age && repoResp.Forks

}

func validateCommitRule(commitPolicy *Commit, commitInfo *CommitInfo, resp *CommitResp) bool {
	var val int
	resp.LastCommitAge, val = checkExpr(commitPolicy.LastCommitAge, uint(commitInfo.LastCommitAt))
	logger.Printf("commitInfo.LastCommitAt: %v\n", commitInfo.LastCommitAt)
	logger.Printf("commitInfo.LastCommitAge: %v\n", val)
	return resp.LastCommitAge
}

func validateReleaseRule(releasePolicy *Release, releaseInfo *ReleaseInfo, resp *ReleaseResp) bool {
	var val int
	resp.LastRelease, val = checkExpr(releasePolicy.LastRelease, uint(releaseInfo.LastReleaseAt))
	logger.Printf("releaseInfo.LastReleaseAt: %v\n", releaseInfo.LastReleaseAt)
	logger.Printf("releaseInfo.LastReleaseAge: %v\n", val)
	return resp.LastRelease
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
