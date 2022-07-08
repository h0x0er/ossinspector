package ossinspector

func Validate(config *Config, repoInfo *RepoInfo) bool {

	// TODO: verify repo rules
	return validateRepoRule(&config.TrustRules.RepoRules, repoInfo)

	//TODO; verify author rule.
	// validateAuthorRule()
	// TODO: to verify commit rule
	// validateCommitRule()

}

func validateRepoRule(policy *RepoRule, repo *RepoInfo) bool {

	stars_resp := checkExpr(policy.Stars, repo.StaggersCount)
	watcher_resp := checkExpr(policy.Watchers, repo.WatcherCount)
	forks_resp := checkExpr(policy.Forks, repo.ForkCount)

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
