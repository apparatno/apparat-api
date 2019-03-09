package www

import (
	"apparat-api/src/app"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

type EmployeeController struct{}

func (c EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files, err := ioutil.ReadDir("data/employees")
	if err != nil {
		log.Printf("could not read employees: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	employees := make(map[string]app.Employee)
	for _, file := range files {
		fileName := file.Name()

		fileBytes, err := ioutil.ReadFile(fmt.Sprintf("data/employees/%s", fileName))
		if err != nil {
			log.Printf("could not read employee file: %s", err)
			continue
		}

		var employee app.Employee
		err = json.Unmarshal(fileBytes, &employee)
		if err != nil {
			log.Printf("could not parse bytes to employee: %s", err)
			continue
		}

		fileNameWithoutExtension := fileName[:len(fileName)-5]
		employees[fileNameWithoutExtension] = employee
	}

	employeesBytes, err := json.Marshal(employees)
	if err != nil {
		log.Printf("could not marshal employees: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(employeesBytes)
	if err != nil {
		log.Printf("could not write employees to response: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func (c EmployeeController) GetEmployeeByName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeName := p.ByName("name")

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("data/employees/%s.json", employeeName))
	if err != nil {
		log.Printf("could not read employee file: %s", err)
		http.Error(w, fmt.Sprintf("no employee named '%s'", employeeName), http.StatusBadRequest)
		return
	}

	var employee app.Employee
	err = json.Unmarshal(fileBytes, &employee)
	if err != nil {
		log.Printf("could not parse bytes to employee: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	employeeBytes, err := json.Marshal(employee)
	if err != nil {
		log.Printf("could not marshal employees: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(employeeBytes)
	if err != nil {
		log.Printf("could not write employee to response: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
