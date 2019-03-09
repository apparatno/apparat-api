package config

import "apparat-api/src/www"

type InversionOfControlContainer struct {
	EmployeeController *www.EmployeeController
	Configuration      *Configuration
}

var DI InversionOfControlContainer

func SetupDependencies() {
	configuration := &Configuration{}

	employeeController := &www.EmployeeController{}

	DI = InversionOfControlContainer{
		EmployeeController: employeeController,
		Configuration:      configuration,
	}
}
