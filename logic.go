package gohamcrest

type not struct {
	trueMatcher Matcher
}

func (this *not)Match(actual interface{}) bool {
	return !this.trueMatcher.Match(actual)
}

func (this *not)FailReason(actual interface{})string {
	return this.trueMatcher.NegationFailReason(actual)
}
func (this *not)NegationFailReason(actual interface{})string {
	return ""
}

func Not(matcher Matcher) Matcher {
	return &not{
		trueMatcher:matcher,
	}
}