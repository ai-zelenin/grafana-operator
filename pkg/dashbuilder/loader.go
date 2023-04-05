package dashbuilder

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ai-zelenin/grafana-operator/pkg/prometheus"
)

type Loader struct {
	cli *prometheus.Client
}

func NewLoader(cli *prometheus.Client) *Loader {
	return &Loader{
		cli: cli,
	}
}

func (l *Loader) LoadPanelContext(ctx context.Context, templates *Templates, panels []*PanelConfig) (map[string]PanelContext, error) {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	sem := make(chan struct{}, 10)
	result := make(map[string]PanelContext)
	for _, panelCfg := range panels {
		wg.Add(1)
		go func(p *PanelConfig) {
			defer func() {
				wg.Done()
				<-sem
			}()
			sem <- struct{}{}
			key := p.Name
			//meta, err := l.LoadMetricMetadata(ctx, key)
			//if err != nil {
			//	log.Println(err)
			//	return
			//}
			template, templateJson, ok := templates.GetTemplate(p.TemplateName)
			if !ok {
				log.Fatal(fmt.Errorf("cannot find template '%s'", p.TemplateName))
			}
			pc := PanelContext{
				Name:         key,
				Folder:       p.Folder,
				MD:           prometheus.Metadata{},
				Customs:      p.Customs,
				TemplateJSON: templateJson,
				Template:     *template,
			}
			mx.Lock()
			result[key] = pc
			mx.Unlock()
		}(panelCfg)
	}
	wg.Wait()
	return result, nil
}

func (l *Loader) LoadMetricMetadata(ctx context.Context, m string) (*prometheus.Metadata, error) {
	resp, err := l.cli.Metadata(ctx, m, "")
	if err != nil {
		return nil, err
	}
	metric, ok := resp[m]
	if !ok {
		return nil, fmt.Errorf("not found %s in prometheus", m)
	}
	if len(metric) == 0 {
		return nil, fmt.Errorf("not found %s in prometheus", m)
	}
	mm := metric[0]
	return &prometheus.Metadata{
		Type: string(mm.Type),
		Unit: mm.Unit,
		Help: mm.Help,
	}, nil
}
