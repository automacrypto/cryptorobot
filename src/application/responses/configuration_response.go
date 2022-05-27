package application

type StepResponse struct {
	name     string
	interval uint
}

type ConfigurationResponse struct {
	app   string
	steps []StepResponse
}
