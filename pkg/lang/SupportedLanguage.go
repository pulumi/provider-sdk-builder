package lang

import (
	"fmt"
	"github.com/pulumi/provider-sdk-builder/pkg/shell"
	"strings"
)

const ALL_LANGUAGES = "all"
const DECOMPOSED_ALL_LANGUAGES = "go,python,dotnet,nodejs,java"

type SupportedLanguage string

type Language interface {
	String() string
	GenerateSdkRecipe(providerName, path, outputPath string) shell.ShellCommandSequence
	CompileSdkRecipe() shell.ShellCommandSequence
	PackageSdkRecipie() shell.ShellCommandSequence
}

var (
	dotNet Language = DotNet{}
	goLang Language = GoLang{}
	java   Language = Java{}
	nodeJS Language = NodeJS{}
	python Language = Python{}
)

var LanguageMap = map[string]Language{
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
	value, isValidLanguage := LanguageMap[rawLangStr]
	if !isValidLanguage {
		return nil, fmt.Errorf("invalid language: %s. Supported languages are: go, python, dotnet, nodejs, java", rawLangStr)
	}
	return value, nil
}
