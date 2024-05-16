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

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeListAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // CreateEmployee - Saves new employee
   CreateEmployee(ctx *gin.Context)

    // DeleteEmployee - Deletes specific employee
   DeleteEmployee(ctx *gin.Context)

    // GetEmployee - Provides details about the employee
   GetEmployee(ctx *gin.Context)

    // GetEmployeeList - Provides details the ambulance employee list
   GetEmployeeList(ctx *gin.Context)

    // UpdateEmployee - Updates specific employee
   UpdateEmployee(ctx *gin.Context)

 }

// partial implementation of EmployeeListAPI - all functions must be implemented in add on files
type implEmployeeListAPI struct {

}

func newEmployeeListAPI() EmployeeListAPI {
  return &implEmployeeListAPI{}
}

func (this *implEmployeeListAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/ambulances/:ambulanceId/employees", this.CreateEmployee)
  routerGroup.Handle( http.MethodDelete, "/ambulances/:ambulanceId/employees/:employeeId", this.DeleteEmployee)
  routerGroup.Handle( http.MethodGet, "/ambulances/:ambulanceId/employees/:employeeId", this.GetEmployee)
  routerGroup.Handle( http.MethodGet, "/ambulances/:ambulanceId/employees", this.GetEmployeeList)
  routerGroup.Handle( http.MethodPut, "/ambulances/:ambulanceId/employees/:employeeId", this.UpdateEmployee)
}