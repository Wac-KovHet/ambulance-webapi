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

type AmbulanceListAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // CreateAmbulance - Saves new ambulance
   CreateAmbulance(ctx *gin.Context)

    // DeleteAmbulance - Deletes specific ambulance
   DeleteAmbulance(ctx *gin.Context)

    // GetAmbulance - Provides details about the ambulance
   GetAmbulance(ctx *gin.Context)

    // GetAmbulanceList - Provides the ambulance list
   GetAmbulanceList(ctx *gin.Context)

    // UpdateAmbulance - Updates specific ambulance
   UpdateAmbulance(ctx *gin.Context)

 }

// partial implementation of AmbulanceListAPI - all functions must be implemented in add on files
type implAmbulanceListAPI struct {

}

func newAmbulanceListAPI() AmbulanceListAPI {
  return &implAmbulanceListAPI{}
}

func (this *implAmbulanceListAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/ambulances", this.CreateAmbulance)
  routerGroup.Handle( http.MethodDelete, "/ambulances/:ambulanceId", this.DeleteAmbulance)
  routerGroup.Handle( http.MethodGet, "/ambulances/:ambulanceId", this.GetAmbulance)
  routerGroup.Handle( http.MethodGet, "/ambulances", this.GetAmbulanceList)
  routerGroup.Handle( http.MethodPut, "/ambulances/:ambulanceId", this.UpdateAmbulance)
}
