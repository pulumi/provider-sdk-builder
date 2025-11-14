package lang

import (
	"fmt"
	"strings"
)

const ALL_LANGUAGES = "all"

type Language interface {
	String() string
	GenerateSdkRecipe(schemaPath, outputPath, version, providerPath string) []string
	CompileSdkRecipe(outputPath, providerPath string) []string
	InstallSdkRecipe(sdkLocation, installLocation string) []string
}

var (
	dotNet Language = DotNet{}
	goLang Language = GoLang{}
	java   Language = Java{}
	nodeJS Language = NodeJS{}
	python Language = Python{}
)

var _languages = map[string]Language{
	"dotnet": dotNet,
	"go":     goLang,
	"java":   java,
	"nodejs": nodeJS,
	"python": python,
}

func AllSupportedLanguages() []Language {
	return []Language{dotNet, goLang, java, nodeJS, python}
}

func ParseRequestedLanguages(rawString string) ([]Language, error) {

	// All should be translated to a list of languages here
	if strings.EqualFold(ALL_LANGUAGES, rawString) {
		return AllSupportedLanguages(), nil
	}

	var languageTokens []string = strings.Split(rawString, ",")
	var result []Language

	for _, value := range languageTokens {
		lang, err := parseLanguageToken(value)
		if err != nil {
			return nil, err
		}
		result = append(result, lang)
	}
	return result, nil
}

func parseLanguageToken(rawLangStr string) (Language, error) {
	value, isValidLanguage := _languages[rawLangStr]
	if !isValidLanguage {
		return nil, fmt.Errorf("invalid language: %s. Supported languages are: go, python, dotnet, nodejs, java", rawLangStr)
	}
	return value, nil
}
