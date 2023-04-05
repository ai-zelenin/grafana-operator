package dashbuilder

type Desc interface {
	Name() string
	Description() string
	Unit() string
	MetricType() string
}
