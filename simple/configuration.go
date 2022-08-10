package simple

type Configuration struct {
	Name string
}
type Application struct {
	*Configuration
}

// provider func
func NewApplication() *Application {
	return &Application{
		// Configuration akan we jdkan as provider caranya di injector.go gunakan wire.FieldOf
		Configuration: &Configuration{
			Name: "ini berasal dari field dari suatu struct (Configuration), sedangkan provider ini dari struct application",
		},
	}
}
