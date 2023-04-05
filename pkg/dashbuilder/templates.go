package dashbuilder

import (
	"encoding/json"

	"github.com/ai-zelenin/grafana-operator/pkg/grafana"
)

type Templates struct {
	cfg   *Config
	jsons map[string][]byte
}

func NewTemplates(cfg *Config) (*Templates, error) {
	result := make(map[string][]byte)
	for k, panel := range cfg.Templates {
		data, err := json.MarshalIndent(panel, "", "\t")
		if err != nil {
			return nil, err
		}
		result[k] = data
	}
	return &Templates{
		cfg:   cfg,
		jsons: result,
	}, nil
}

func (t *Templates) GetTemplate(key string) (template *grafana.Panel, templateJson []byte, ok bool) {
	template, ok = t.cfg.Templates[key]
	if !ok {
		return
	}
	templateJson, ok = t.jsons[key]
	if !ok {
		return
	}
	return
}
