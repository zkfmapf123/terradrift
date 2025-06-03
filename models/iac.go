package models

type IacPlanPath string

type DriftResultsParams struct {
	Add     int
	Change  int
	Destroy int
}

type IaCParams struct {
	Method   string
	PlanPath []string
	Results  map[IacPlanPath]DriftResultsParams
}

type DriftResultFuncs interface {
	Plan()
}
