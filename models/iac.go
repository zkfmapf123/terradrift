package models

type DriftResultsParams struct {
	Add     string
	Change  string
	Destroy string
}

type IaCParams struct {
	PlanPath []string
	Results  map[string]DriftResultsParams
}

type DriftResultFuncs interface {
	AllPush(paths []string)
	Push(path string)
	Plan(int)
}
