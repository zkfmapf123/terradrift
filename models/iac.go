package models

type DriftResultsParams struct {
	Add     int
	Change  int
	Destroy int
}

type IaCParams struct {
	PlanPath []string
	Results  map[string]DriftResultsParams
}

type DriftResultFuncs interface {
	Push(path string)
	Plan(concurrency int)
}
