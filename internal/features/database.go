package features

func init() {
	Register(&databaseFeature{})
}

type databaseFeature struct{}

func (f *databaseFeature) ID() string {
	return "database"
}

func (f *databaseFeature) Name() string {
	return "Database - Conexión Go-Supabase"
}

func (f *databaseFeature) Description() string {
	return "Go-Supabase connection with models"
}

func (f *databaseFeature) CompatibleStacks() []string {
	return []string{"web", "mobile", "web-fullstack"}
}
