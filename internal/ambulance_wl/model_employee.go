package ambulance_wl

type Employee struct {

	// Unique id of the employee
	Id string `json:"id"`

	// Name of employee
	Name string `json:"name"`

	// Surname of employee
	Surname string `json:"surname"`

	// Date of birth of employee
	DateOfBirth string `json:"dateOfBirth"`

	// Position of employee
	Position string `json:"position"`

	// Wage of employee
	Wage int32 `json:"wage"`
}

func MapEmployeeToResponse(employee Employee) EmployeeResponse {
	return EmployeeResponse{
		Id:          employee.Id,
		Name:        employee.Name,
		Surname:     employee.Surname,
		DateOfBirth: employee.DateOfBirth,
		Position:    employee.Position,
		Wage:        employee.Wage,
	}
}