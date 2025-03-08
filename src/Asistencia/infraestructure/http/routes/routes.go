package routes

import (
	"driver-attendance/src/Asistencia/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/asistencia")

	RegistrarAsistencia := dependencies.RegistrarAsistenciaCtrl()

	routes.POST("/registrar", RegistrarAsistencia.Run)
}