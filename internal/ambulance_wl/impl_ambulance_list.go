package ambulance_wl

import (
	"net/http"

	"github.com/Wac-KovHet/ambulance-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAmbulance - Saves new ambulance
func (this *implAmbulanceListAPI) CreateAmbulance(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Ambulance])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	var ambulanceRequest AmbulanceRequest
	err := ctx.BindJSON(&ambulanceRequest)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// Validate the request
	err = validate.Struct(ambulanceRequest)
	if err != nil {
		errors := formatValidationError(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Validation failed",
				"error": "Validation failed",
				"errors": errors,
			})
		return
	}

	ambulance := Ambulance{
		Id:       uuid.New().String(),
		Name:     ambulanceRequest.Name,
		Location: ambulanceRequest.Location,
		Contact:  ambulanceRequest.Contact,
	}

	err = db.CreateDocument(ctx, ambulance.Id, &ambulance)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			MapAmbulanceToResponse(ambulance),
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Ambulance already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create ambulance in database",
				"error":   err.Error(),
			},
		)
	}
}

// DeleteAmbulance - Deletes specific ambulance
func (this *implAmbulanceListAPI) DeleteAmbulance(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
					"status":  "Internal Server Error",
					"message": "db_service not found",
					"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Ambulance])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
					"status":  "Internal Server Error",
					"message": "db_service context is not of type db_service.DbService",
					"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	ambulanceId := ctx.Param("ambulanceId")
	err := db.DeleteDocument(ctx, ambulanceId)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to delete ambulance from database",
					"error":   err.Error(),
			})
	}
}

// GetAmbulance - Provides details about the ambulance
func (this *implAmbulanceListAPI) GetAmbulance(ctx *gin.Context) {
 	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Ambulance])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	ambulanceId := ctx.Param("ambulanceId")
	ambulance, err := db.FindDocument(ctx, ambulanceId)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			MapAmbulanceToResponse(*ambulance),
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Ambulance not found",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve ambulance from database",
				"error":   err.Error(),
			})
	}
}

// GetAmbulanceList - Provides the ambulance list
func (this *implAmbulanceListAPI) GetAmbulanceList(ctx *gin.Context) {
 	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Ambulance])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	ambulances, err := db.FindAllDocuments(ctx)
	switch err {
	case nil:
		var response []AmbulanceResponse
		for _, ambulance := range ambulances {
			response = append(response, MapAmbulanceToResponse(ambulance))
		}
		ctx.JSON(
			http.StatusOK,
			response,
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve ambulances from database",
				"error":   err.Error(),
			})
	}
}

// UpdateAmbulance - Updates specific ambulance
func (this *implAmbulanceListAPI) UpdateAmbulance(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Ambulance])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	ambulanceId := ctx.Param("ambulanceId")

	var ambulanceRequest AmbulanceRequest
	err := ctx.BindJSON(&ambulanceRequest)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// Validate the request
	err = validate.Struct(ambulanceRequest)
	if err != nil {
		errors := formatValidationError(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Validation failed",
				"error": "Validation failed",
				"errors": errors,
			})
		return
	}

	ambulance, err := db.FindDocument(ctx, ambulanceId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
				})
		default:
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to retrieve ambulance from database",
					"error":   err.Error(),
				})
		}
		return
	}

	ambulance.Name = ambulanceRequest.Name
	ambulance.Location = ambulanceRequest.Location
	ambulance.Contact = ambulanceRequest.Contact

	err = db.UpdateDocument(ctx, ambulanceId, ambulance)
	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			MapAmbulanceToResponse(*ambulance),
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Ambulance not found",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update ambulance in database",
				"error":   err.Error(),
			})
	}
}
