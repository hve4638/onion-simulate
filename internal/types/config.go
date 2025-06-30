package types

type Config struct {
	NodeCount NodeCountConfig `yaml:"node_count"`
	Server    []ServerConfig  `yaml:"server"`
	User      []UserConfig    `yaml:"user"`
	Routines  []RoutineConfig `yaml:"routines"`
}

type NodeCountConfig struct {
	Amount         int `yaml:"amount"`
	AllowExitCount int `yaml:"allow_exit_count"`
}

type ServerConfig struct {
	Id    string `yaml:"id"`
	Outer bool   `yaml:"outer"`
}

type UserConfig struct {
	Id       string   `yaml:"id"`
	Routines []string `yaml:"routines"`
}

type RoutineConfig struct {
	Id                    string       `yaml:"id"`
	RepeatCountRange      int64        `yaml:"repeat_count"`
	RepeatIntervalRange   IntRange     `yaml:"repeat_interval_range"`
	Period                PeriodConfig `yaml:"period"`
	URL                   string       `yaml:"url"`
	CommunicateCountRange IntRange     `yaml:"communicate_count_range"`
}

type PeriodConfig struct {
	Week      int `yaml:"week"`
	TimeRange struct {
		Min struct {
			Hour   int64 `yaml:"hour"`
			Minute int64 `yaml:"minute"`
			Second int64 `yaml:"second"`
		} `yaml:"min"`
		Max struct {
			Hour   int64 `yaml:"hour"`
			Minute int64 `yaml:"minute"`
			Second int64 `yaml:"second"`
		} `yaml:"max"`
	} `yaml:"time_range"`
}

type IntRange struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}
