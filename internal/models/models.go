package models

type Profiles struct {
	profile map[string]Profile `yaml:"profile"`
}

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}
