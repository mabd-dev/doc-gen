package pipeline

import (
	"fmt"
	"regexp"
)

// getDocsOnly takes output from llm and extract only the kdoc part
func getDocsOnly(docs string) (string, error) {
	kdocRegex := regexp.MustCompile(`/\*\*(.|[\r\n])*?\*/`)
	matches := kdocRegex.FindAllString(docs, -1)

	if len(matches) > 0 {
		return matches[0], nil
	}

	err := fmt.Errorf("could not find kdoc in llm response")
	return "", err
}
