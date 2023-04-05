package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ai-zelenin/grafana-operator/pkg/dashbuilder"
	"github.com/ai-zelenin/grafana-operator/pkg/grafana"
	"github.com/ai-zelenin/grafana-operator/pkg/prometheus"
)

func main() {
	var configFile = flag.String("c", "config.yml", "path to config file")
	flag.Parse()
	ctx := context.Background()
	data, err := os.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}
	cfg := &dashbuilder.Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	grafanaCli := grafana.NewClient(cfg.GrafanaURL, cfg.GrafanaToken)
	promCli, err := prometheus.NewClient(cfg.PrometheusURL)
	if err != nil {
		panic(err)
	}
	l := dashbuilder.NewLoader(promCli)
	templates, err := dashbuilder.NewTemplates(cfg)
	if err != nil {
		panic(err)
	}
	pcs, err := l.LoadPanelContext(ctx, templates, cfg.PanelLoaderConfig)
	if err != nil {
		panic(err)
	}
	b := dashbuilder.NewBuilder(templates)
	for _, panelContext := range pcs {
		panel, err := b.BuildPanel(panelContext)
		if err != nil {
			panic(err)
		}
		//fmt.Println(panel)
		//d, err := l.MarshalTemplates(map[string]*grafana.Panel{"a": panel})
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(string(d["a"]))
		//id, err := grafanaCli.DeleteLibraryPanel(panel.UID)
		//if err != nil {
		//	log.Println(err)
		//}
		//resp, err := grafanaCli.CreateLibraryPanel(panel, panelContext.Folder)
		//if err != nil {
		//	log.Println(err)
		//}
		resp, err := grafanaCli.SaveLibraryPanel(panel, panelContext.Folder)
		if err != nil {
			panic(fmt.Errorf("SaveLibraryPanel %v", err))
		}
		fmt.Println(dashbuilder.DumpJSON(resp))
	}
}
