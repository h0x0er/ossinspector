package ossinspector

func Validate(config *Config, repoInfo *RepoInfo) {

	// TODO: verify repo rule
	for _, repoRule := range config.RepoTrustRule.RepoRules {
		validateRepoRule(&repoRule, repoInfo)
	}

	//TODO; verify author rule.

	// TODO: to verify commit rule

}

func validateRepoRule(repoRule *RepoRule, repoInfo *RepoInfo) bool {

	return true

}
