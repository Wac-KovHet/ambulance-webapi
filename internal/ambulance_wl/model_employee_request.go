/*
 * Ambulance Api
 *
 * Ambulance management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xkovhet@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance_wl

type EmployeeRequest struct {

	// Name of employee
	Name string `json:"name"`

	// Surname of employee
	Surname string `json:"surname"`

	// Date of birth of employee
	DateOfBirth string `json:"dateOfBirth"`

	// Position of employee
	Position string `json:"position"`

	// Wage
	Wage int32 `json:"wage"`
}
