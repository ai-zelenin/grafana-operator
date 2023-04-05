package grafana

type PanelType string

const (
	PanelTypeRow        PanelType = "row"
	PanelTypeBarchart   PanelType = "barchart"
	PanelTypeTimeseries PanelType = "timeseries"
)

type UnitType string

const (
	UnitTypeShort   UnitType = "short"
	UnitTypeBytes   UnitType = "bytes"
	UnitTypeSeconds UnitType = "s"
)

var (
	DefaultGrid = &GridPos{
		H: 8,
		W: 12,
	}
	DefaultColorMod  = "palette-classic"
	DefaultDrawStyle = "line"
)

type Panel struct {
	ID          int       `json:"id,omitempty" yaml:"id"`
	UID         string    `json:"uid,omitempty" yaml:"uid"`
	Type        PanelType `json:"type,omitempty" yaml:"type"`
	Title       string    `json:"title,omitempty" yaml:"title"`
	Description string    `json:"description,omitempty" yaml:"description"`
	GridPos     *GridPos  `json:"gridPos,omitempty" yaml:"gridPos"`

	Repeat          string `json:"repeat,omitempty" yaml:"repeat"`
	RepeatDirection string `json:"repeatDirection,omitempty" yaml:"repeatDirection"`
	MaxPerRow       int    `json:"maxPerRow,omitempty" yaml:"maxPerRow"`

	Mode        string         `json:"mode,omitempty" yaml:"mode"`
	Content     string         `json:"content,omitempty" yaml:"content"`
	Targets     []*Target      `json:"targets,omitempty" yaml:"targets"`
	Options     map[string]any `json:"options,omitempty" yaml:"options"`
	FieldConfig FieldConfig    `json:"fieldConfig,omitempty" yaml:"fieldConfig"`
	Datasource  Datasource     `json:"datasource,omitempty" yaml:"datasource"`
}

type FieldConfig struct {
	Defaults  PanelDefaults `json:"defaults,omitempty" yaml:"defaults"`
	Overrides []any         `json:"overrides,omitempty" yaml:"overrides"`
}

type PanelOptions struct {
	Legend  PanelLegend  `json:"legend,omitempty" yaml:"legend"`
	Tooltip PanelTooltip `json:"tooltip,omitempty" yaml:"tooltip"`
}

type PanelLegend struct {
	Calcs       []string `json:"calcs,omitempty" yaml:"calcs"`
	DisplayMode string   `json:"displayMode,omitempty" yaml:"displayMode"`
	Placement   string   `json:"placement,omitempty" yaml:"placement"`
	ShowLegend  bool     `json:"showLegend,omitempty" yaml:"showLegend"`
}

type PanelTooltip struct {
	Mode string `json:"mode,omitempty" yaml:"mode"`
	Sort string `json:"sort,omitempty" yaml:"sort"`
}

type PanelDefaults struct {
	Color  PanelColor  `json:"color,omitempty" yaml:"color"`
	Custom PanelCustom `json:"custom,omitempty" yaml:"custom"`
	Unit   string      `json:"unit,omitempty" yaml:"unit"`
}

type PanelColor struct {
	Mode string `json:"mode" yaml:"mode"`
}
type LibraryPanel struct {
	Name string `json:"name" yaml:"name"`
	Uid  string `json:"uid" yaml:"uid"`
}

func NewPanel(panelType PanelType, title string) *Panel {
	custom := PanelCustom{}
	custom.SetDrawStyle("line")
	return &Panel{
		Type:    panelType,
		Title:   title,
		GridPos: DefaultGrid,
		FieldConfig: FieldConfig{
			Defaults: PanelDefaults{
				Color: PanelColor{
					Mode: DefaultColorMod,
				},
				Custom: custom,
			},
		},
	}
}

func NewRowPanel(title string) *Panel {
	return NewPanel(PanelTypeRow, title)
}

func (p *Panel) AddTarget(target *Target) *Panel {
	p.Targets = append(p.Targets, target)
	return p
}

func (p *Panel) SetSize(h, w, x, y int) {
	p.GridPos = &GridPos{
		H: h,
		W: w,
		X: x,
		Y: y,
	}
}

type GridPos struct {
	X int `json:"x,omitempty" yaml:"x"`
	Y int `json:"y,omitempty" yaml:"y"`
	W int `json:"w,omitempty" yaml:"w"`
	H int `json:"h,omitempty" yaml:"h"`
}
