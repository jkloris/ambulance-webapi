/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: <your_email>
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package ambulance_wl

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type AmbulanceConditionsAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // GetConditions - Provides the list of conditions associated with ambulance
   GetConditions(ctx *gin.Context)

 }

 // partial implementation of AmbulanceConditionsAPI - all functions must be implemented in add on files
type implAmbulanceConditionsAPI struct {

}

func newAmbulanceConditionsAPI() AmbulanceConditionsAPI {
  return &implAmbulanceConditionsAPI{}
}

func (this *implAmbulanceConditionsAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodGet, "/waiting-list/:ambulanceId/condition", this.GetConditions)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // GetConditions - Provides the list of conditions associated with ambulance
// func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

