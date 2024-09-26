package config

type Cronjob struct {
	TaskSpec string `yaml:"task_spec" default:"@every 10m"`
	TaskJob  bool   `yaml:"task_job" default:"false"`
}
