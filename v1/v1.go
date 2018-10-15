package v1

import "golang.org/x/text/language"

type V1 struct{}

func (v V1) String() string {
	return language.BritishEnglish.String()
}
