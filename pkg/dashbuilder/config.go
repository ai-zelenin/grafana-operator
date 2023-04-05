package dashbuilder

import (
	"github.com/ai-zelenin/grafana-operator/pkg/grafana"
)

type Config struct {
	GrafanaURL        string                    `json:"grafana_url" yaml:"grafana_url"`
	GrafanaToken      string                    `json:"grafana_token" yaml:"grafana_token"`
	PrometheusURL     string                    `json:"prometheus_url" yaml:"prometheus_url"`
	Templates         map[string]*grafana.Panel `json:"templates" yaml:"templates"`
	PanelLoaderConfig []*PanelConfig            `json:"panels" yaml:"panels"`
}

type PanelConfig struct {
	Name         string         `json:"name" yaml:"name"`
	Folder       string         `json:"folder" yaml:"folder"`
	TemplateName string         `json:"template" yaml:"template"`
	Customs      map[string]any `json:"customs" yaml:"customs"`
}
