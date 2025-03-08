package controllers

import (
	"driver-attendance/src/Asistencia/application"
	"driver-attendance/src/Asistencia/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegistrarAsistenciaCtrl struct {
	uc *application.RegistrarAsistenciaUC
}

func NewRegistrarAsistenciaCtrl(uc *application.RegistrarAsistenciaUC) *RegistrarAsistenciaCtrl {
	return &RegistrarAsistenciaCtrl{uc: uc}
}

func (ctrl *RegistrarAsistenciaCtrl) Run(c *gin.Context) {
	choferID, _ := strconv.Atoi(c.Query("chofer_id"))
    busID, _ := strconv.Atoi(c.Query("bus_id"))

	asistencia := domain.Asistencia{
		ChoferID: choferID,
		BusID:    busID,
	}
	err := ctrl.uc.Run(asistencia)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Asistencia registrada"})
}