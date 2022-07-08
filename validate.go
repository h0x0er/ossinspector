package ossinspector

func Validate(config *Config, repoInfo *RepoInfo) {

	// TODO: verify repo rules
	validateRepoRule(&config.TrustRules.RepoRules, repoInfo)

	//TODO; verify author rule.

	// TODO: to verify commit rule

}

func validateRepoRule(repoRule *RepoRule, repoInfo *RepoInfo) bool {

	return true

}
