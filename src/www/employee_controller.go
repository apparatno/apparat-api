package www

import (
	"apparat-api/src/app"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type EmployeeController struct{}

func (c EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files, err := getEmployeeFiles()
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

		employeeName := getEmployeeName(fileName)
		employees[employeeName] = employee
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

func (c EmployeeController) GetEmployeeNames(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files, err := getEmployeeFiles()
	if err != nil {
		log.Printf("could not read employees: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	var employeeNames []string
	for _, file := range files {
		employeeName := getEmployeeName(file.Name())
		employeeNames = append(employeeNames, employeeName)
	}

	namesJson, err := json.Marshal(employeeNames)
	if err != nil {
		log.Printf("could not marshal employee names: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(namesJson)
	if err != nil {
		log.Printf("could not write names to response: %s", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func getEmployeeFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir("data/employees")
	if err != nil {
		return nil, err
	}

	excludeFiles := []string{"default.json"}

	var employees []os.FileInfo
	for _, file := range files {
		exclude := false

		for _, ex := range excludeFiles {
			if file.Name() == ex {
				exclude = true
				break
			}
		}

		if !exclude {
			employees = append(employees, file)
		}
	}

	return employees, nil
}

func getEmployeeName(fileName string) string {
	return getFileNameWithoutExtension(fileName, ".json")
}

func getFileNameWithoutExtension(fileName string, ext string) string {
	return fileName[:len(fileName)-len(ext)]
}
