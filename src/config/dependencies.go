package config

import "apparat-api/src/www"

type Container struct {
	EmployeeController *www.EmployeeController
	Configuration      *Configuration
}

var DI Container

func SetupDI() {
	configuration := &Configuration{}

	employeeController := &www.EmployeeController{}

	DI = Container{
		EmployeeController: employeeController,
		Configuration:      configuration,
	}
}
