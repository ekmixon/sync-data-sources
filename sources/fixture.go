package syncdatasources

import "fmt"

// Config holds data source config options
type Config struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Endpoint holds data source endpoint options
type Endpoint struct {
	Name string `yaml:"name"`
}

// DataSource contains data source spec from dev-analytics-api
type DataSource struct {
	Slug      string     `yaml:"slug"`
	Config    []Config   `yaml:"config"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

// Fixture contains full YAML structure of dev-analytics-api fixture files
type Fixture struct {
	Native      map[string]string `yaml:"native"`
	DataSources []DataSource      `yaml:"data_sources"`
	Fn          string
	Slug        string
}

// Task holds single endpoint task and its context (required config, fixture filename etc.)
type Task struct {
	Endpoint Endpoint
	Config   []Config
	DsSlug   string
	FxSlug   string
	FxFn     string
}

func (t Task) String() string {
	return fmt.Sprintf("{Endpoint:%s DS:%s Slug:%s File:%s NConfigs:%d}", t.Endpoint.Name, t.DsSlug, t.FxSlug, t.FxFn, len(t.Config))
}
