package v2

import "golang.org/x/text/language"

type V2 struct{}

func (v V2) String() string {
	return language.AmericanEnglish.String()
}
