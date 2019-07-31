package study

// Study is double blind study. Instances of a study are called trials.
type Study struct {
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}

// Group represents a study group. A study will generally have a Control Group
// and one or more Experiment Groups.
type Group struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}
