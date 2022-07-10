package ossinspector

func Validate(policy *Policy, repoInfo *RepoInfo) bool {

	// TODO: verify repo rules
	return validateRepoRule(&policy.Repo, repoInfo)

	//TODO; verify author rule.
	// validateAuthorRule()
	// TODO: to verify commit rule
	// validateCommitRule()

}

func validateRepoRule(policy *Repo, repo *RepoInfo) bool {

	stars_resp := checkExpr(policy.Stars, repo.StaggersCount)
	watcher_resp := checkExpr(policy.Watchers, repo.WatcherCount)
	forks_resp := checkExpr(policy.Forks, repo.ForkCount)


	// TODO: output needed to be improved
	return stars_resp && watcher_resp && forks_resp

}

func checkExpr(checkString string, value uint) bool {
	resp := false
	ctype, pvalue := evaluate(checkString)
	switch ctype {
	case LESSER_THAN:
		resp = (value < uint(pvalue))
	case GREATER_THAN:
		resp = (value > uint(pvalue))
	}
	return resp
}

func checkBool(fromPolicy bool, fromRepo bool) bool {
	return fromPolicy == fromRepo
}
