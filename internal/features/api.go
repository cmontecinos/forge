package features

func init() {
	Register(&apiFeature{})
}

type apiFeature struct{}

func (f *apiFeature) ID() string {
	return "api"
}

func (f *apiFeature) Name() string {
	return "API - Router, middlewares, handlers"
}

func (f *apiFeature) Description() string {
	return "Echo router with middlewares and handlers"
}

func (f *apiFeature) CompatibleStacks() []string {
	return []string{"web", "mobile"}
}
