package loglint

import (
	"github.com/Davidianol/loglint/internal/config"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

func New(settings any) (register.LinterPlugin, error) {
	cfg, err := config.Parse(settings)
	if err != nil {
		return nil, err
	}
	return &plugin{cfg: cfg}, nil
}

type plugin struct {
	cfg *config.Config
}

var _ register.LinterPlugin = new(plugin)

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{NewAnalyzer(p.cfg)}, nil
}

func (*plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
