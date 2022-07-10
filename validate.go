package ossinspector

import (
	"log"
	"strings"
	"time"
)

func Validate(policy *Policy, repoInfo *RepoInfo) bool {

	// TODO: verify repo rules
	return validateRepoRule(&policy.Repo, repoInfo)

	//TODO; verify author rule.
	// validateAuthorRule()
	// TODO: to verify commit rule
	// validateCommitRule()

}

func validateRepoRule(repoPolicy *Repo, repo *RepoInfo) bool {

	stars_resp := checkExpr(repoPolicy.Stars, repo.StaggersCount)
	watcher_resp := checkExpr(repoPolicy.Watchers, repo.WatcherCount)
	forks_resp := checkExpr(repoPolicy.Forks, repo.ForkCount)
	age_resp := checkExpr(repoPolicy.Age, uint(repo.CreatedAt))

	log.Printf("stars_resp: %v\n", stars_resp)
	log.Printf("watcher_resp: %v\n", watcher_resp)
	log.Printf("forks_resp: %v\n", forks_resp)
	log.Printf("age_resp: %v\n", age_resp)
	
	// TODO: output needed to be improved
	return stars_resp && watcher_resp && forks_resp && age_resp

}

func checkExpr(checkString string, value uint) bool {
	checkString = strings.ReplaceAll(checkString, " ", "")
	resp := false
	ctype, pvalue := evaluate(checkString) // pvalue = policy value

	switch ctype {
	case LESSER_THAN:
		resp = (value < uint(pvalue))
	case GREATER_THAN:
		resp = (value > uint(pvalue))

	case DAYS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := current - value
		days := (diff / (60 * 60 * 24))
		resp = (days < uint(pvalue))

	case DAYS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		days := (diff / (60 * 60 * 24))
		resp = (days > uint(pvalue))

	case MONTHS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		months := (diff / (60 * 60 * 24 * 30))
		resp = (months < uint(pvalue))

	case MONTHS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		months := (diff / (60 * 60 * 24 * 30))
		resp = (months > uint(pvalue))

	case YEARS_LESSER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		years := (diff / (60 * 60 * 24 * 30 * 365))
		resp = (years < uint(pvalue))

	case YEARS_GREATER_THAN:
		current := uint(time.Now().Unix())
		diff := (current - value)
		years := (diff / (60 * 60 * 24 * 30 * 365))
		resp = (years > uint(pvalue))

	}
	return resp
}

func checkBool(fromPolicy bool, fromRepo bool) bool {
	return fromPolicy == fromRepo
}
