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

    // CreateAmbulanceEntry - Saves new entry into ambulance list
   CreateAmbulanceEntry(ctx *gin.Context)

    // DeleteAmbulanceEntry - Deletes specific entry
   DeleteAmbulanceEntry(ctx *gin.Context)

    // GetAmbulanceEntry - Provides details about the ambulance
   GetAmbulanceEntry(ctx *gin.Context)

    // GetAmbulanceList - Provides the ambulance list
   GetAmbulanceList(ctx *gin.Context)

    // UpdateAmbulanceEntry - Updates specific entry
   UpdateAmbulanceEntry(ctx *gin.Context)

 }

// partial implementation of AmbulanceListAPI - all functions must be implemented in add on files
type implAmbulanceListAPI struct {

}

func newAmbulanceListAPI() AmbulanceListAPI {
  return &implAmbulanceListAPI{}
}

func (this *implAmbulanceListAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/ambulances", this.CreateAmbulanceEntry)
  routerGroup.Handle( http.MethodDelete, "/ambulances/:ambulanceId", this.DeleteAmbulanceEntry)
  routerGroup.Handle( http.MethodGet, "/ambulances/:ambulanceId", this.GetAmbulanceEntry)
  routerGroup.Handle( http.MethodGet, "/ambulances", this.GetAmbulanceList)
  routerGroup.Handle( http.MethodPut, "/ambulances/:ambulanceId", this.UpdateAmbulanceEntry)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // CreateAmbulanceEntry - Saves new entry into ambulance list
// func (this *implAmbulanceListAPI) CreateAmbulanceEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteAmbulanceEntry - Deletes specific entry
// func (this *implAmbulanceListAPI) DeleteAmbulanceEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetAmbulanceEntry - Provides details about the ambulance
// func (this *implAmbulanceListAPI) GetAmbulanceEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetAmbulanceList - Provides the ambulance list
// func (this *implAmbulanceListAPI) GetAmbulanceList(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateAmbulanceEntry - Updates specific entry
// func (this *implAmbulanceListAPI) UpdateAmbulanceEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

