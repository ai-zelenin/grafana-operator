package grafana

const SchemaVersion = 37

type Dashboard struct {
	ID            int64         `json:"id,omitempty" yaml:"id"`
	UID           string        `json:"uid,omitempty" yaml:"uid"`
	Title         string        `json:"title,omitempty" yaml:"title"`
	Tags          []string      `json:"tags,omitempty" yaml:"tags"`
	Style         string        `json:"style,omitempty" yaml:"style"`
	Timezone      string        `json:"timezone,omitempty" yaml:"timezone"`
	Editable      bool          `json:"editable,omitempty" yaml:"editable"`
	GraphTooltip  int           `json:"graphTooltip,omitempty" yaml:"graphTooltip"`
	Panels        []*Panel      `json:"panels,omitempty" yaml:"panels"`
	Time          *Time         `json:"time,omitempty" yaml:"time"`
	Timepicker    *Timepicker   `json:"timepicker,omitempty" yaml:"timepicker"`
	Templating    *Templating   `json:"templating,omitempty" yaml:"templating"`
	Annotations   *Annotations  `json:"annotations,omitempty" yaml:"annotations"`
	Refresh       string        `json:"refresh,omitempty" yaml:"refresh"`
	SchemaVersion int           `json:"schemaVersion,omitempty" yaml:"schemaVersion"`
	Version       int           `json:"version,omitempty" yaml:"version"`
	Links         []interface{} `json:"links,omitempty" yaml:"links"`
}

func NewDashboard(title string) *Dashboard {
	return &Dashboard{
		Title:         title,
		Panels:        make([]*Panel, 0),
		Tags:          make([]string, 0),
		Timezone:      "browser",
		SchemaVersion: SchemaVersion,
	}
}

func (d *Dashboard) AddPanel(p *Panel) *Dashboard {
	d.Panels = append(d.Panels, p)
	return d
}

type Time struct {
	From string `json:"from,omitempty" yaml:"from"`
	To   string `json:"to,omitempty" yaml:"to"`
}

type Timepicker struct {
	TimeOptions      []interface{} `json:"time_options,omitempty" yaml:"time_options"`
	RefreshIntervals []string      `json:"refresh_intervals,omitempty" yaml:"refresh_intervals"`
	Collapse         bool          `json:"collapse,omitempty" yaml:"collapse"`
	Enable           bool          `json:"enable,omitempty" yaml:"enable"`
	Notice           bool          `json:"notice,omitempty" yaml:"notice"`
	Now              bool          `json:"now,omitempty" yaml:"now"`
	Status           string        `json:"status,omitempty" yaml:"status"`
	Type             string        `json:"type,omitempty" yaml:"type"`
}

type Annotations struct {
	List []interface{} `json:"list,omitempty" yaml:"list"`
}

type Templating struct {
	Enable bool       `json:"enable,omitempty" yaml:"enable"`
	List   []Variable `json:"list,omitempty" yaml:"list"`
}

type Current struct {
	Tags  []interface{} `json:"tags,omitempty" yaml:"tags"`
	Text  string        `json:"text,omitempty" yaml:"text"`
	Value string        `json:"value,omitempty" yaml:"value"`
}

type DashboardOptions struct {
	Selected bool   `json:"selected,omitempty" yaml:"selected"`
	Text     string `json:"text,omitempty" yaml:"text"`
	Value    string `json:"value,omitempty" yaml:"value"`
}

type Variable struct {
	AllFormat   string             `json:"allFormat,omitempty" yaml:"allFormat"`
	Current     Current            `json:"current,omitempty" yaml:"current"`
	Datasource  interface{}        `json:"datasource,omitempty" yaml:"datasource"`
	IncludeAll  bool               `json:"includeAll,omitempty" yaml:"includeAll"`
	Name        string             `json:"name,omitempty" yaml:"name"`
	Options     []DashboardOptions `json:"options,omitempty" yaml:"options"`
	Query       string             `json:"query,omitempty" yaml:"query"`
	Refresh     bool               `json:"refresh,omitempty" yaml:"refresh"`
	Type        string             `json:"type,omitempty" yaml:"type"`
	Multi       bool               `json:"multi,omitempty" yaml:"multi"`
	MultiFormat string             `json:"multiFormat,omitempty" yaml:"multiFormat"`
	Regex       string             `json:"regex,omitempty" yaml:"regex"`
}
