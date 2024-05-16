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

type AmbulanceRequestEntry struct {

	// Unique id of the ambulance
	Id string `json:"id"`

	// Name of ambulance
	Name string `json:"name"`

	// Location of ambulance
	Location string `json:"location"`

	// Contact for ambulance
	Contact string `json:"contact"`
}
