package meta

type TagOptionConfig struct {
	Separator     string   `json:"separator"`
	AcceptedModes []string `json:"accepted_modes"`
}

type TagOption struct {
	config        TagOptionConfig
	AcceptedModes StringSet
}

func (tc TagOptionConfig) ToOptions() (to TagOption) {
	return TagOption{
		config:        tc,
		AcceptedModes: NewStringSet(tc.AcceptedModes),
	}
}
