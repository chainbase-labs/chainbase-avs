package node

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"

	"github.com/cbergoon/merkletree"
)

func extractJobID(log string) (string, error) {
	// Define the regular expression pattern
	pattern := `Job has been submitted with JobID (\w+)`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the match in the log
	match := re.FindStringSubmatch(log)

	// Check if a match was found
	if len(match) < 2 {
		return "", fmt.Errorf("JobID not found in the log")
	}

	// Return the captured JobID
	return match[1], nil
}

// ResultContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type ResultContent struct {
	result string
}

// CalculateHash hashes the values of a ResultContent
func (r ResultContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(r.result)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (r ResultContent) Equals(other merkletree.Content) (bool, error) {
	orc, ok := other.(ResultContent)
	if !ok {
		return false, errors.New("value is not of type ResultContent")
	}
	return r.result == orc.result, nil
}

func isValidNodeSocket(nodeSocket string) bool {
	pattern := `^([0-9]{1,3}\.){3}[0-9]{1,3}:[0-9]{1,5}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(nodeSocket)
}
