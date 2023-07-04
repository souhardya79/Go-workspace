package symin

import (
	"flag"

	"github.com/syself/backend/internal/service"
	"github.com/syself/backend/internal/testlogservice"
	"github.com/syself/backend/pkg/flagext"
	utillog "github.com/syself/backend/pkg/util/log"
)

// Config is the main configuration object for symin.
type Config struct {
	Target flagext.StringSliceCSV `yaml:"target"`

	LogLevel utillog.Level `yaml:"logLevel"`

	TestService    service.Config        `yaml:"testService"`
	TestLogService testlogservice.Config `yaml:"testLogService"`
}

func (c *Config) RegisterFlags(f *flag.FlagSet) {
	// Set the default module list to 'all'
	c.Target = []string{Service, TestLogService}
	f.Var(&c.Target, "target",
		"A comma-separated list of components to run. "+
			"The default value 'all' runs Symin in single binary mode.",
	)

	// Set default
	c.LogLevel = utillog.LevelInfo
	f.Func("loglevel", "Defines the log level.", func(s string) error {
		c.LogLevel = utillog.Level(s)
		return nil
	})

	c.TestService.RegisterFlags(f)
	c.TestLogService.RegisterFlags(f)
}
