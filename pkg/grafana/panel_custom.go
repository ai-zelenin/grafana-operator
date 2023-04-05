package grafana

type PanelCustom map[string]any

type ScaleDistributionType string
type StackingMode string

const (
	ScaleDistributionTypeLinear ScaleDistributionType = "linear"
	StackingModeNone            StackingMode          = "none"
)

const (
	axisCenteredZero  = "axisCenteredZero"
	axisColorMode     = "axisColorMode"
	axisPlacement     = "axisPlacement"
	barAlignment      = "barAlignment"
	DrawStyle         = "drawStyle"
	fillOpacity       = "fillOpacity"
	gradientMode      = "gradientMode"
	lineInterpolation = "lineInterpolation"
	lineWidth         = "lineWidth"
	pointSize         = "pointSize"
	scaleDistribution = "scaleDistribution"
	showPoints        = "showPoints"
	spanNulls         = "spanNulls"
	stacking          = "stacking"
)

func (c PanelCustom) SetAxisCenteredZero(val bool) PanelCustom {
	c[axisCenteredZero] = val
	return c
}

func (c PanelCustom) SetAxisColorMode(val string) PanelCustom {
	c[axisColorMode] = val
	return c
}

func (c PanelCustom) SetAxisPlacement(val string) PanelCustom {
	c[axisPlacement] = val
	return c
}

func (c PanelCustom) SetBarAlignment(val int) PanelCustom {
	c[barAlignment] = val
	return c
}

func (c PanelCustom) SetDrawStyle(val string) PanelCustom {
	c[DrawStyle] = val
	return c
}

func (c PanelCustom) SetFillOpacity(val int) PanelCustom {
	c[fillOpacity] = val
	return c
}

func (c PanelCustom) SetGradientMode(val string) PanelCustom {
	c[gradientMode] = val
	return c
}

func (c PanelCustom) SetLineInterpolation(val string) PanelCustom {
	c[lineInterpolation] = val
	return c
}

func (c PanelCustom) SetLineWidth(val int) PanelCustom {
	c[lineWidth] = val
	return c
}

func (c PanelCustom) SetPointSize(val int) PanelCustom {
	c[pointSize] = val
	return c
}

func (c PanelCustom) SetScaleDistribution(val ScaleDistributionType) PanelCustom {
	c[scaleDistribution] = map[string]any{"type": val}
	return c
}

func (c PanelCustom) SetShowPoints(val string) PanelCustom {
	c[showPoints] = val
	return c
}

func (c PanelCustom) SetSpanNulls(val bool) PanelCustom {
	c[spanNulls] = val
	return c
}

func (c PanelCustom) SetStacking(val StackingMode) PanelCustom {
	c[stacking] = map[string]any{"group": "A", "mode": val}
	return c
}
