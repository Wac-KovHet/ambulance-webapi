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

func MapAmbulanceToResponse(ambulance Ambulance) AmbulanceResponse {
	var nurseCount, doctorCount int32
	for _, employee := range ambulance.Employees {
		switch employee.Position {
		case "Nurse":
			nurseCount++
		case "Doctor":
			doctorCount++
		}
	}

	return AmbulanceResponse{
		Id:          ambulance.Id,
		Name:        ambulance.Name,
		Location:    ambulance.Location,
		Contact:     ambulance.Contact,
		NurseCount:  nurseCount,
		DoctorCount: doctorCount,
	}
}