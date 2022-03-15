package model

import "unicode/utf8"

type Status int

const (
	UNCHECKED Status = iota
	UNUSED
	BITE
	EAT
)

type Results []*Result

type Result struct {
	Answer   string
	Statuses []Status
}

func NewResult(answer, target string) *Result {
	return &Result{
		Answer:   answer,
		Statuses: getStatuses(answer, target),
	}
}

func getStatuses(answer, target string) []Status {
	statuses := make([]Status, utf8.RuneCountInString(answer))
	for ansIdx, ansLetter := range answer {
		status := UNUSED
		for targetIdx, targetLetter := range target {
			if ansLetter == targetLetter {
				if ansIdx == targetIdx {
					status = EAT
					break
				} else {
					status = BITE
				}
			}
		}
		statuses[ansIdx] = status
	}
	return statuses
}

func (rs Results) Solved() bool {
	for _, r := range rs {
		if r.Solved() {
			return true
		}
	}
	return false
}

func (r *Result) Solved() bool {
	for _, status := range r.Statuses {
		if status != EAT {
			return false
		}
	}
	return true
}

func (rs Results) GetStatusMap() map[rune]Status {
	statusMap := rs.newStatusMap()
	for _, r := range rs {
		for ansIdx, ansLetter := range r.Answer {
			status := r.Statuses[ansIdx]
			if status > statusMap[ansLetter] {
				statusMap[ansLetter] = status
			}
		}
	}
	return statusMap
}

func (rs Results) newStatusMap() map[rune]Status {
	statusMap := map[rune]Status{}
	for i := 0; i < 26; i++ {
		letter := rune('a' + i)
		statusMap[letter] = UNCHECKED
	}
	return statusMap
}
