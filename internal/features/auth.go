package features

func init() {
	Register(&authFeature{})
}

type authFeature struct{}

func (f *authFeature) ID() string {
	return "auth"
}

func (f *authFeature) Name() string {
	return "Auth - Login/registro via backend"
}

func (f *authFeature) Description() string {
	return "Authentication via Go backend with JWT"
}

func (f *authFeature) CompatibleStacks() []string {
	return []string{"web", "mobile", "web-fullstack"}
}
