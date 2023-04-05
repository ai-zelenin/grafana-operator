package dashbuilder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/ai-zelenin/grafana-operator/pkg/grafana"
)

type Builder struct {
	*Templates
}

func NewBuilder(t *Templates) *Builder {
	return &Builder{
		Templates: t,
	}
}

func (b *Builder) BuildPanel(pc PanelContext) (resultPanel *grafana.Panel, err error) {
	panelData := string(pc.TemplateJSON)
	if pc.Customs != nil {
		customs := pc.Customs
		customsJson, err := json.MarshalIndent(customs, "", "\t")
		if err != nil {
			return nil, err
		}
		gjson.ForEachLine(string(customsJson), func(line gjson.Result) bool {
			walk(line, "", func(path string, value any) {
				before := gjson.Get(panelData, path).String()
				log.Printf("[%s] %v -> %v \n", path, before, value)

				panelData, err = sjson.Set(panelData, path, value)
				if err != nil {
					log.Fatal("sjson.Set", err)
					return
				}
			})
			return true
		})
	}

	resultPanel = new(grafana.Panel)
	err = json.Unmarshal([]byte(panelData), resultPanel)
	if err != nil {
		return nil, err
	}
	// UID
	uid := resultPanel.UID
	if uid != "" {
		uid, err = Template(uid, pc)
		if err != nil {
			return nil, err
		}
	}
	resultPanel.UID = uid

	//title
	panelTitle := resultPanel.Title
	if panelTitle != "" {
		panelTitle, err = Template(panelTitle, pc)
		if err != nil {
			return nil, err
		}
	}
	resultPanel.Title = panelTitle

	// description
	panelDescription := resultPanel.Description
	if panelDescription != "" {
		panelDescription, err = Template(panelDescription, pc)
		if err != nil {
			return nil, err
		}
	}
	resultPanel.Description = panelDescription

	// unit
	unit := resultPanel.FieldConfig.Defaults.Unit
	if unit != "" {
		unit, err = Template(unit, pc)
		if err != nil {
			return nil, err
		}
	}
	resultPanel.FieldConfig.Defaults.Unit = unit

	// targets
	for i, target := range resultPanel.Targets {
		b.InitTargetDefault(pc.MD.Type, pc.Name, resultPanel.Targets[i], pc)
		expr, err := Template(target.Expr, pc)
		if err != nil {
			return nil, err
		}
		resultPanel.Targets[i].Expr = expr
	}

	return resultPanel, nil
}

func (b *Builder) InitTargetDefault(metricType, metricName string, t *grafana.Target, pc PanelContext) {
	if t.Expr == "" {
		switch metricType {
		case "gauge":
			t.Expr = fmt.Sprintf(`%s{env="$env"}`, metricName)
		case "counter":
			t.Expr = fmt.Sprintf(`rate(%s{env="$env"}[$__rate_interval])`, metricName)
		case "histogram":
			t.Expr = fmt.Sprintf(`rate(%s{env="$env"}[$__rate_interval])`, metricName)
			t.Format = grafana.TargetFormatHeatmap
		default:
			t.Expr = metricName
		}
	}

	if t.LegendFormat == "" {
		t.LegendFormat = "__auto"
	}
	if t.Datasource.Uid == "" {
		t.Datasource.Uid = "{DS}"
	}
	if t.Datasource.Type == "" {
		t.Datasource.Type = grafana.DatasourceTypePrometheus
	}
	if t.EditorMode == "" {
		t.EditorMode = "code"
	}
	if t.LegendFormat == "" {
		t.LegendFormat = "__auto"
	}
	if t.RefId == "" {
		t.RefId = "A"
	}
}

func (b *Builder) InitPanelDefault(p *grafana.Panel) {
	if p.Type == "" {
		p.Type = grafana.PanelTypeTimeseries
	}
	if p.GridPos == nil {
		p.GridPos = grafana.DefaultGrid
	}
	if p.FieldConfig.Defaults.Color.Mode == "" {
		p.FieldConfig.Defaults.Color.Mode = grafana.DefaultColorMod
	}
	if p.FieldConfig.Defaults.Custom[grafana.DrawStyle] == "" {
		p.FieldConfig.Defaults.Custom.SetDrawStyle(grafana.DefaultDrawStyle)
	}
	if p.FieldConfig.Defaults.Unit == "" {
		p.FieldConfig.Defaults.Unit = string(grafana.UnitTypeShort)
	}
}

func Template(tmpl string, templateContext any) (string, error) {
	t, err := template.New("").Delims("$$", "$$").Parse(tmpl)
	if err != nil {
		return "", err
	}
	bb := bytes.NewBuffer(nil)
	err = t.Execute(bb, templateContext)
	if err != nil {
		return "", err
	}
	return bb.String(), nil
}

func walk(r gjson.Result, path string, cb func(path string, value any)) {
	r.ForEach(func(key, value gjson.Result) bool {
		var np = strings.TrimPrefix(path+"."+key.String(), ".")
		switch value.Type {
		case gjson.JSON:
			walk(value, np, cb)
		default:
			cb(np, value.Value())
		}
		return true
	})
}
