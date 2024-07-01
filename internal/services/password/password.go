package password

type Password struct {
	commonPatterns []string
}

func New(commonPatternsPath string) (*Password, error) {
	commonPatterns, err := loadCommonPatterns(commonPatternsPath)
	if err != nil {
		return nil, err
	}
	
	return &Password{
		commonPatterns: commonPatterns,
	}, nil
}
