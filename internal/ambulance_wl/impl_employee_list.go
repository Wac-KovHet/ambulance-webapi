package ambulance_wl

import (
	"net/http"

	"github.com/Wac-KovHet/ambulance-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateEmployee - Saves new employee
func (this *implEmployeeListAPI) CreateEmployee(ctx *gin.Context) {
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

	ambulanceId := ctx.Param("ambulanceId")
	ambulance, err := db.FindDocument(ctx, ambulanceId)
	if err != nil {
		if err == db_service.ErrNotFound {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
				})
		} else {
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

	var employeeRequest EmployeeRequest
	err = ctx.BindJSON(&employeeRequest)
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
	err = validate.Struct(employeeRequest)
	if err != nil {
		errors := FormatValidationError(err)
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

	employee := Employee{
		Id:          uuid.New().String(),
		Name:        employeeRequest.Name,
		Surname:     employeeRequest.Surname,
		DateOfBirth: employeeRequest.DateOfBirth,
		Position:    employeeRequest.Position,
		Wage:        employeeRequest.Wage,
	}

	ambulance.Employees = append(ambulance.Employees, employee)
	err = db.UpdateDocument(ctx, ambulanceId, ambulance)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			MapEmployeeToResponse(employee),
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create employee in database",
				"error":   err.Error(),
			},
		)
	}
}

// DeleteEmployee - Deletes specific employee
func (this *implEmployeeListAPI) DeleteEmployee(ctx *gin.Context) {
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
	employeeId := ctx.Param("employeeId")

	ambulance, err := db.FindDocument(ctx, ambulanceId)
	if err != nil {
		if err == db_service.ErrNotFound {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
				})
		} else {
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

	for i, employee := range ambulance.Employees {
		if employee.Id == employeeId {
			ambulance.Employees = append(ambulance.Employees[:i], ambulance.Employees[i+1:]...)
			err = db.UpdateDocument(ctx, ambulanceId, ambulance)
			if err != nil {
				ctx.JSON(
					http.StatusBadGateway,
					gin.H{
						"status":  "Bad Gateway",
						"message": "Failed to update ambulance in database",
						"error":   err.Error(),
					})
				return
			}
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
	}

	ctx.JSON(
		http.StatusForbidden,
		gin.H{
			"status":  "Forbidden",
			"message": "Employee does not fall under the specified ambulance",
			"error":   "Employee does not fall under the specified ambulance",
		},
	)
}

// GetEmployee - Provides details about the employee
func (this *implEmployeeListAPI) GetEmployee(ctx *gin.Context) {
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
	employeeId := ctx.Param("employeeId")

	ambulance, err := db.FindDocument(ctx, ambulanceId)
	if err != nil {
		if err == db_service.ErrNotFound {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
				})
		} else {
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

	for _, employee := range ambulance.Employees {
		if employee.Id == employeeId {
			ctx.JSON(
				http.StatusOK,
				MapEmployeeToResponse(employee),
			)
			return
		}
	}

	ctx.JSON(
		http.StatusForbidden,
		gin.H{
			"status":  "Forbidden",
			"message": "Employee does not fall under the specified ambulance",
			"error":   "Employee does not fall under the specified ambulance",
		},
	)
}

// GetEmployeeList - Provides details the ambulance employee list
func (this *implEmployeeListAPI) GetEmployeeList(ctx *gin.Context) {
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
	if err != nil {
		if err == db_service.ErrNotFound {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
				})
		} else {
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

	var response []EmployeeResponse
	for _, employee := range ambulance.Employees {
		response = append(response, MapEmployeeToResponse(employee))
	}
	ctx.JSON(
		http.StatusOK,
		response,
	)
}

// UpdateEmployee - Updates specific employee
func (this *implEmployeeListAPI) UpdateEmployee(ctx *gin.Context) {
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
	employeeId := ctx.Param("employeeId")

	ambulance, err := db.FindDocument(ctx, ambulanceId)
	if err != nil {
		if err == db_service.ErrNotFound {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Ambulance not found",
					"error":   err.Error(),
				})
		} else {
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

	var employeeRequest EmployeeRequest
	err = ctx.BindJSON(&employeeRequest)
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
	err = validate.Struct(employeeRequest)
	if err != nil {
		errors := FormatValidationError(err)
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

	for i, emp := range ambulance.Employees {
		if emp.Id == employeeId {
			ambulance.Employees[i] = Employee{
				Id:          emp.Id,
				Name:        employeeRequest.Name,
				Surname:     employeeRequest.Surname,
				DateOfBirth: employeeRequest.DateOfBirth,
				Position:    employeeRequest.Position,
				Wage:			 employeeRequest.Wage,
			}
			err = db.UpdateDocument(ctx, ambulanceId, ambulance)
			if err != nil {
				ctx.JSON(
					http.StatusBadGateway,
					gin.H{
						"status":  "Bad Gateway",
						"message": "Failed to update employee in database",
						"error":   err.Error(),
					})
				return
			}
			ctx.JSON(
				http.StatusOK,
				MapEmployeeToResponse(ambulance.Employees[i]),
			)
			return
		}
	}

	ctx.JSON(
		http.StatusForbidden,
		gin.H{
			"status":  "Forbidden",
			"message": "Employee does not fall under the specified ambulance",
			"error":   "Employee does not fall under the specified ambulance",
		},
	)
}
