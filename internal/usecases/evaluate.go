package usecases

func (u *UseCases) EvaluatePassword(password string) int {
	lengthScore := len(password)
	varietyScore := u.password.CalculateVariety(password)
	entropyScore := u.password.CalculateEntropy(password)
	commonPatternScore := u.password.CheckCommonPatterns(password)

	totalScore := lengthScore + varietyScore + entropyScore - commonPatternScore
	if totalScore < 0 {
		return 0
	}

	return totalScore
}
