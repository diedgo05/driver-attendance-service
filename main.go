package main

import (
	dependenciesAsistencia "driver-attendance/src/Asistencia/infraestructure/dependencies"
	routesAsistencia "driver-attendance/src/Asistencia/infraestructure/http/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	//Asistencia
	dependenciesAsistencia.Init()
	routesAsistencia.Routes(r)

	r.Run(":8081")

}