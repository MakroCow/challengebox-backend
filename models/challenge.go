package models

type Challenge struct {
	Title             string
	Description       string
	Id                int
	SportValue        int
	IntelligenceValue int
	CulinaryValue     int
	SocialValue       int
	SelfcareValue     int
}

func NewChallenge(title string, description string, sportValue int, intelligenceValue int, culinaryValue int, socialValue int, selfcareValue int) *Challenge {
	return &Challenge{Title: title, Description: description, SportValue: sportValue, IntelligenceValue: intelligenceValue, CulinaryValue: culinaryValue, SocialValue: socialValue, SelfcareValue: selfcareValue}
}

type ChallengeBuilder struct {
	c Challenge
}

func (cb *ChallengeBuilder) Title(title string) *ChallengeBuilder {
	cb.c.Title = title
	return cb
}

func (cb *ChallengeBuilder) Description(description string) *ChallengeBuilder {
	cb.c.Description = description
	return cb
}

func (cb *ChallengeBuilder) Build() *Challenge {
	return &cb.c
}
