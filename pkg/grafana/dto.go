package grafana

import (
	"time"
)

type SaveDashboardRequest struct {
	Dashboard *Dashboard `json:"dashboard,omitempty" yaml:"dashboard"`
	FolderId  int        `json:"folderId,omitempty" yaml:"folderId"`
	Message   string     `json:"message,omitempty" yaml:"message"`
	Overwrite bool       `json:"overwrite,omitempty" yaml:"overwrite"`
}

type SaveDashboardResponse struct {
	Id      int    `json:"id" yaml:"id"`
	Slug    string `json:"slug" yaml:"slug"`
	Status  string `json:"status" yaml:"status"`
	Uid     string `json:"uid" yaml:"uid"`
	Url     string `json:"url" yaml:"url"`
	Version int    `json:"version" yaml:"version"`
}

type GetLibraryPanelsResponse struct {
	Result struct {
		TotalCount int           `json:"totalCount"`
		Elements   []interface{} `json:"elements"`
		Page       int           `json:"page"`
		PerPage    int           `json:"perPage"`
	} `json:"result"`
}

type SaveLibraryPanelRequest struct {
	Uid       string `json:"uid,omitempty"`
	FolderUid string `json:"folderUid,omitempty"`
	Name      string `json:"name,omitempty"`
	Model     any    `json:"model,omitempty"`
	Kind      int    `json:"kind,omitempty"`
	Version   int    `json:"version,omitempty"`
}

type PanelLibraryResponse struct {
	Result struct {
		Id          int    `json:"id"`
		OrgId       int    `json:"orgId"`
		FolderId    int    `json:"folderId"`
		FolderUid   string `json:"folderUid"`
		Uid         string `json:"uid"`
		Name        string `json:"name"`
		Kind        int    `json:"kind"`
		Type        string `json:"type"`
		Description string `json:"description"`
		Model       any    `json:"model"`
		Version     int    `json:"version"`
		Meta        struct {
			FolderName          string    `json:"folderName"`
			FolderUid           string    `json:"folderUid"`
			ConnectedDashboards int       `json:"connectedDashboards"`
			Created             time.Time `json:"created"`
			Updated             time.Time `json:"updated"`
			CreatedBy           struct {
				Id        int    `json:"id"`
				Name      string `json:"name"`
				AvatarUrl string `json:"avatarUrl"`
			} `json:"createdBy"`
			UpdatedBy struct {
				Id        int    `json:"id"`
				Name      string `json:"name"`
				AvatarUrl string `json:"avatarUrl"`
			} `json:"updatedBy"`
		} `json:"meta"`
	} `json:"result"`
}

type Panels struct {
	Datasource struct {
		Type string `json:"type"`
		Uid  string `json:"uid"`
	} `json:"datasource"`
	Description string `json:"description"`
	FieldConfig struct {
		Defaults struct {
			Color struct {
				Mode string `json:"mode"`
			} `json:"color"`
			Custom struct {
				AxisCenteredZero bool   `json:"axisCenteredZero"`
				AxisColorMode    string `json:"axisColorMode"`
				AxisLabel        string `json:"axisLabel"`
				AxisPlacement    string `json:"axisPlacement"`
				BarAlignment     int    `json:"barAlignment"`
				DrawStyle        string `json:"DrawStyle"`
				FillOpacity      int    `json:"fillOpacity"`
				GradientMode     string `json:"gradientMode"`
				HideFrom         struct {
					Legend  bool `json:"legend"`
					Tooltip bool `json:"tooltip"`
					Viz     bool `json:"viz"`
				} `json:"hideFrom"`
				LineInterpolation string `json:"lineInterpolation"`
				LineWidth         int    `json:"lineWidth"`
				PointSize         int    `json:"pointSize"`
				ScaleDistribution struct {
					Type string `json:"type"`
				} `json:"scaleDistribution"`
				ShowPoints string `json:"showPoints"`
				SpanNulls  bool   `json:"spanNulls"`
				Stacking   struct {
					Group string `json:"group"`
					Mode  string `json:"mode"`
				} `json:"stacking"`
				ThresholdsStyle struct {
					Mode string `json:"mode"`
				} `json:"thresholdsStyle"`
			} `json:"custom"`
			Mappings   []interface{} `json:"mappings"`
			Thresholds struct {
				Mode  string `json:"mode"`
				Steps []struct {
					Color string `json:"color"`
					Value *int   `json:"value"`
				} `json:"steps"`
			} `json:"thresholds"`
		} `json:"defaults"`
		Overrides []interface{} `json:"overrides"`
	} `json:"fieldConfig"`
	GridPos struct {
		H int `json:"h"`
		W int `json:"w"`
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gridPos"`
	Id           int `json:"id"`
	LibraryPanel struct {
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"libraryPanel"`
	Options struct {
		Legend struct {
			Calcs       []interface{} `json:"calcs"`
			DisplayMode string        `json:"displayMode"`
			Placement   string        `json:"placement"`
			ShowLegend  bool          `json:"showLegend"`
		} `json:"legend"`
		Tooltip struct {
			Mode string `json:"mode"`
			Sort string `json:"sort"`
		} `json:"tooltip"`
	} `json:"options"`
	Targets []struct {
		Datasource struct {
			Type string `json:"type"`
			Uid  string `json:"uid"`
		} `json:"datasource"`
		EditorMode   string `json:"editorMode"`
		Expr         string `json:"expr"`
		LegendFormat string `json:"legendFormat"`
		Range        bool   `json:"range"`
		RefId        string `json:"refId"`
	} `json:"targets"`
	Title string `json:"title"`
	Type  string `json:"type"`
}
