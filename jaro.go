package jaro

import (
	"math"
	"strings"
)

type Options struct {
	CaseSensitive bool
}

func Distance(s1 string, s2 string, options Options) float64 {
	m := 0

	// Exit early if either are empty.
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	// Set default values for options if they are not provided.
	if !options.CaseSensitive {
		options.CaseSensitive = true
	}

	// Convert to upper if case-sensitive is false.
	if !options.CaseSensitive {
		s1 = strings.ToUpper(s1)
		s2 = strings.ToUpper(s2)
	}

	// Exit early if they're an exact match.
	if s1 == s2 {
		return 1
	}

	rangeVal := (int(math.Floor(math.Max(float64(len(s1)), float64(len(s2))) / 2))) - 1
	s1Matches := make([]bool, len(s1))
	s2Matches := make([]bool, len(s2))

	for i := 0; i < len(s1); i++ {
		low := int(math.Max(float64(i-rangeVal), 0)) // Use math.Max to handle negative indices
		high := int(math.Min(float64(i+rangeVal), float64(len(s2)-1)))

		for j := low; j <= high; j++ {
			if !s1Matches[i] && !s2Matches[j] && s1[i] == s2[j] {
				m++
				s1Matches[i] = true
				s2Matches[j] = true
				break
			}
		}
	}

	// Exit early if no matches were found.
	if m == 0 {
		return 0
	}

	// Count the transpositions.
	k := 0
	numTrans := 0

	for i := 0; i < len(s1); i++ {
		if s1Matches[i] {
			for j := k; j < len(s2); j++ {
				if s2Matches[j] {
					k = j + 1
					break
				}

				if s1[i] != s2[j] {
					numTrans++
				}
			}
		}
	}

	weight := (float64(m)/float64(len(s1)) + float64(m)/float64(len(s2)) + (float64(m)-float64(numTrans)/2)/float64(m)) / 3

	l := 0
	p := 0.1

	if weight > 0.7 {
		for l < 4 && l < len(s1) && l < len(s2) && s1[l] == s2[l] {
			l++
		}

		weight = weight + float64(l)*p*(1-weight)
	}

	return weight
}
