package main

// Pipeline
// Ref: https://concourse-ci.org/pipelines.html
type Pipeline struct {
	Groups        []Group        `yaml:"groups,omitempty"`
	VarSources    []VarSource    `yaml:"var_sources,omitempty"`
	ResourceTypes []ResourceType `yaml:"resource_types,omitempty"`
	Resources     []Resource     `yaml:"resources,omitempty"`
	Jobs          []Job          `yaml:"jobs,omitempty"`
}

// Group
// Ref: https://concourse-ci.org/pipelines.html#schema.pipeline.groups
type Group struct {
	Name string   `yaml:"name,omitempty"`
	Jobs []string `yaml:"jobs,omitempty"`
}

// VarSource
// Ref: https://concourse-ci.org/vars.html#schema.var_source
type VarSource struct {
	Name   string      `yaml:"name,omitempty"`
	Type   string      `yaml:"type,omitempty"`
	Config interface{} `yaml:"config,omitempty"`
}

// ResourceType
// Ref: https://concourse-ci.org/resource-types.html#schema.resource_type
type ResourceType struct {
	Name                 string      `yaml:"name,omitempty"`
	Type                 string      `yaml:"type,omitempty"`
	Source               interface{} `yaml:"source,omitempty"`
	Privileged           bool        `yaml:"privileged,omitempty"`
	Params               interface{} `yaml:"params,omitempty"`
	CheckEvery           string      `yaml:"check_every,omitempty"`
	Tags                 []string    `yaml:"tags,omitempty"`
	UniqueVersionHistory bool        `yaml:"unique_version_history,omitempty"`
}

// Resource
// Ref: https://concourse-ci.org/resources.html#schema.resource
type Resource struct {
	Name         string      `yaml:"name,omitempty"`
	OldName      string      `yaml:"old_name,omitempty"`
	Type         string      `yaml:"type,omitempty"`
	Icon         string      `yaml:"icon,omitempty"`
	Source       interface{} `yaml:"source,omitempty"`
	Version      interface{} `yaml:"version,omitempty"`
	CheckEvery   string      `yaml:"check_every,omitempty"`
	Tags         []string    `yaml:"tags,omitempty"`
	Public       bool        `yaml:"public,omitempty"`
	WebhookToken string      `yaml:"webhook_token,omitempty"`
}

// Job
// Ref: https://concourse-ci.org/jobs.html
type Job struct {
	Name                 string            `yaml:"name,omitempty"`
	OldName              string            `yaml:"old_name,omitempty"`
	Serial               bool              `yaml:"serial,omitempty"`
	SerialGroups         []string          `yaml:"serial_groups,omitempty"`
	BuildLogRetention    BuildLogRetention `yaml:"build_log_retention,omitempty"`
	BuildLogsToRetain    int               `yaml:"build_logs_to_retain,omitempty"`
	MaxInFlight          int               `yaml:"max_in_flight,omitempty"`
	Public               bool              `yaml:"public,omitempty"`
	DisableManualTrigger bool              `yaml:"disable_manual_trigger,omitempty"`
	Interruptible        bool              `yaml:"interruptible,omitempty"`
	Plan                 []Step            `yaml:"plan,omitempty"`
	OnSuccess            Step              `yaml:"on_success,omitempty"`
	OnFailure            Step              `yaml:"on_failure,omitempty"`
	OnError              Step              `yaml:"on_error,omitempty"`
	OnAbort              Step              `yaml:"on_abort,omitempty"`
	Ensure               Step              `yaml:"ensure,omitempty"`
}

// BuildLogRetention
// Ref: https://concourse-ci.org/jobs.html#schema.job.build_log_retention
type BuildLogRetention struct {
	Days               int `yaml:"days,omitempty"`
	Builds             int `yaml:"builds,omitempty"`
	MinSucceededBuilds int `yaml:"minimum_succeeded_builds,omitempty"`
}

// Step
// Ref: https://concourse-ci.org/jobs.html#schema.step
type Step struct {
	// Top level
	Get         string         `yaml:"get,omitempty"`
	Put         string         `yaml:"put,omitempty"`
	Task        string         `yaml:"task,omitempty"`
	SetPipeline string         `yaml:"set_pipeline,omitempty"`
	LoadVar     string         `yaml:"load_var,omitempty"`
	InParallel  InParallelStep `yaml:"in_parallel,omitempty"`
	Aggregate   []Step         `yaml:"aggregate,omitempty"`
	Do          []Step         `yaml:"do,omitempty"`
	Try         *Step          `yaml:"try,omitempty"`
	// Sub-options
	Resource      string      `yaml:"resource,omitempty"`
	Passed        []string    `yaml:"passed,omitempty"`
	File          string      `yaml:"file,omitempty"`
	Image         string      `yaml:"image,omitempty"`
	Params        interface{} `yaml:"params,omitempty"`
	Trigger       bool        `yaml:"trigger,omitempty"`
	Version       interface{} `yaml:"version,omitempty"`
	Inputs        string      `yaml:"inputs,omitempty"`
	GetParams     interface{} `yaml:"get_params,omitempty"`
	Config        TaskConfig  `yaml:"config,omitempty"`
	Privileged    bool        `yaml:"privileged,omitempty"`
	Vars          interface{} `yaml:"vars,omitempty"`
	InputMapping  interface{} `yaml:"input_mapping,omitempty"`
	OutputMapping interface{} `yaml:"output_mapping,omitempty"`
	Format        string      `yaml:"format,omitempty"`
	Reveal        bool        `yaml:"reveal,omitempty"`
	Timeout       string      `yaml:"timeout,omitempty"`
	Attempts      int         `yaml:"attempts,omitempty"`
	Tags          []string    `yaml:"tags,omitempty"`
	OnSuccess     *Step       `yaml:"on_success,omitempty"`
	OnFailure     *Step       `yaml:"on_failure,omitempty"`
	OnError       *Step       `yaml:"on_error,omitempty"`
	OnAbort       *Step       `yaml:"on_abort,omitempty"`
	Ensure        *Step       `yaml:"ensure,omitempty"`
}

// InParallelStep
// Ref: https://concourse-ci.org/jobs.html#schema.in_parallel_config
type InParallelStep struct {
	Steps    []Step `yaml:"steps,omitempty"`
	Limit    int    `yaml:"limit,omitempty"`
	FailFast bool   `yaml:"fail_fast,omitempty"`
}

// Task
// Ref: https://concourse-ci.org/tasks.html
type TaskConfig struct {
	Platform        string            `yaml:"platform,omitempty"`
	ImageSource     AnonymousResource `yaml:"image_source,omitempty"`
	Inputs          []Input           `yaml:"inputs,omitempty"`
	Outputs         []Output          `yaml:"outputs,omitempty"`
	Caches          []Cache           `yaml:"cache,omitempty"`
	Params          interface{}       `yaml:"params,omitempty"`
	Run             Command           `yaml:"run,omitempty"`
	RootfsUri       string            `yaml:"roofs_uri,omitempty"`
	ContainerLimits ContainerLimits   `yaml:"container_limits,omitempty"`
}

// AnonymousResource
// Ref: https://concourse-ci.org/tasks.html#schema.anonymous_resource
type AnonymousResource struct {
	Type    string      `yaml:"type,omitempty"`
	Source  interface{} `yaml:"source,omitempty"`
	Params  interface{} `yaml:"params,omitempty"`
	Version interface{} `yaml:"version,omitempty"`
}

// Input
// Ref: https://concourse-ci.org/tasks.html#schema.input
type Input struct {
	Name     string `yaml:"name,omitempty"`
	Path     string `yaml:"path,omitempty"`
	Optional bool   `yaml:"optional,omitempty"`
}

// Output
// Ref: https:/concourse-ci.org/tasks.html#schema.output
type Output struct {
	Name string `yaml:"name,omitempty"`
	Path string `yaml:"path,omitempty"`
}

// Cache
// Ref: https://concourse-ci.org/tasks.html#schema.cache
type Cache struct {
	Path string `yaml:"path,omitempty"`
}

// Command
// Ref: https://concourse-ci.org/tasks.html#schema.command
type Command struct {
	Path string   `yaml:"path,omitempty"`
	Args []string `yaml:"args,omitempty"`
	Dir  string   `yaml:"dir,omitempty"`
	User string   `yaml:"user,omitempty"`
}

// ContainerLimits
// Ref: https://concourse-ci.org/tasks.html#schema.container_limits
type ContainerLimits struct {
	Cpu    int `yaml:"cpu,omitempty"`
	Memory int `yaml:"memory,omitempty"`
}
