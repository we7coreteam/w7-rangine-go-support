package provider

type Provider interface {
	GetPackageName() string
	Register()
}

type Abstract struct {
	Provider
	PackageName string
}

func (abstract *Abstract) GetPackageName() string {
	return abstract.PackageName
}
