package grafana

type DatasourceType string

const (
	DatasourceTypePrometheus = "prometheus"
)

type Datasource struct {
	Type string `json:"type" yaml:"type"`
	Uid  string `json:"uid" yaml:"uid"`
}

func NewDatasource(dsType string, uid string) *Datasource {
	return &Datasource{Type: dsType, Uid: uid}
}

func (d Datasource) NewTarget(expr, legend string) *Target {
	return &Target{
		Datasource:   d,
		EditorMode:   "code",
		Expr:         expr,
		Range:        true,
		LegendFormat: legend,
		RefId:        "A",
	}
}

type TargetFormat string

const (
	TargetFormatTimeSeries = "time_series"
	TargetFormatHeatmap    = "heatmap"
)

type Target struct {
	Datasource   Datasource   `json:"datasource" yaml:"datasource"`
	EditorMode   string       `json:"editorMode" yaml:"editorMode"`
	Expr         string       `json:"expr" yaml:"expr"`
	LegendFormat string       `json:"legendFormat" yaml:"legendFormat"`
	Format       TargetFormat `json:"format,omitempty" yaml:"format"`
	Range        bool         `json:"range" yaml:"range"`
	RefId        string       `json:"refId" yaml:"refId"`
}
