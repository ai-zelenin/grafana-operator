package dashbuilder

import (
	"github.com/ai-zelenin/grafana-operator/pkg/grafana"
	"github.com/ai-zelenin/grafana-operator/pkg/prometheus"
)

type PanelContext struct {
	Name         string
	Folder       string
	MD           prometheus.Metadata
	Template     grafana.Panel
	TemplateJSON []byte
	Customs      map[string]any
}
