package ambulance_wl

type Ambulance struct {
	// Unique id of the ambulance
	Id string `json:"id"`

	// Name of ambulance
	Name string `json:"name"`

	// Location of ambulance
	Location string `json:"location"`

	// Contact for ambulance
	Contact string `json:"contact"`

	// List of employees
	Employees []Employee `json:"employees,omitempty"`
}