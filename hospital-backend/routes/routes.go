package routes

import (
	"hospital-backend/aseguradora"
	"hospital-backend/cirugia"
	"hospital-backend/cita"
	"hospital-backend/consulta"
	"hospital-backend/diagnostico"
	facturacion "hospital-backend/factura"
	"hospital-backend/hospitalizacion"
	"hospital-backend/inventario"
	"hospital-backend/medicamento"
	"hospital-backend/medico"
	"hospital-backend/middleware"
	"hospital-backend/paciente"
	"hospital-backend/pago"
	"hospital-backend/personal"
	"hospital-backend/poliza"
	"hospital-backend/prueba"
	"hospital-backend/receta"
	"hospital-backend/rol"
	"hospital-backend/sala"
	"hospital-backend/tipo_prueba"
	"hospital-backend/tratamiento"
	"hospital-backend/usuario"
	"hospital-backend/usuario_rol"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	// ---------- AUTH ----------
	userService := usuario.NewService(db)
	userHandler := usuario.NewHandler(userService)
	api.POST("/auth/login", userHandler.Login)

	// Rutas públicas (solo lectura) — accesibles desde frontend sin token
	// ---------- MEDICOS ----------
	medicoRepoPub := medico.NewRepository(db)
	medicoServicePub := medico.NewService(medicoRepoPub)
	medicoHandlerPub := medico.NewHandler(medicoServicePub)
	medicosPub := api.Group("/medicos")
	medicosPub.GET("", medicoHandlerPub.GetAll)
	medicosPub.GET("/:id", medicoHandlerPub.GetByID)

	// ---------- PACIENTES ----------
	pacienteRepoPub := paciente.NewPacienteRepository(db)
	pacienteServicePub := paciente.NewPacienteService(pacienteRepoPub)
	pacienteHandlerPub := paciente.NewPacienteHandler(pacienteServicePub)
	pacientesPub := api.Group("/pacientes")
	pacientesPub.GET("", pacienteHandlerPub.ObtenerTodos)
	pacientesPub.GET("/:id", pacienteHandlerPub.ObtenerPorID)

	// ---------- MEDICAMENTOS ----------
	medRepoPub := medicamento.NewRepository(db)
	medSvcPub := medicamento.NewService(medRepoPub)
	medHandlerPub := medicamento.NewHandler(medSvcPub)
	medsPub := api.Group("/medicamentos")
	medsPub.GET("", medHandlerPub.GetAll)
	medsPub.GET("/:id", medHandlerPub.GetByID)

	// ---------- SALAS ----------
	salaRepo := sala.NewRepository(db)
	salaSvc := sala.NewService(salaRepo)
	salaHandler := sala.NewHandler(salaSvc)
	salasPub := api.Group("/salas")
	salasPub.GET("", salaHandler.GetAll)
	salasPub.GET("/:id", salaHandler.GetByID)

	// ---------- PRUEBAS ----------
	pruebaHandlerPub := prueba.NewHandler(db)
	pruebasPub := api.Group("/pruebas")
	pruebasPub.GET("", pruebaHandlerPub.GetAll)
	pruebasPub.GET("/:id", pruebaHandlerPub.GetByID)
	pruebasPub.GET("/paciente/:pacienteID", pruebaHandlerPub.GetByPaciente)

	// ---------- POLIZAS ----------
	polizaRepo := poliza.NewRepository(db)
	polizaSvc := poliza.NewService(polizaRepo)
	polizaHandler := poliza.NewHandler(polizaSvc)
	polizasPub := api.Group("/polizas")
	polizasPub.GET("", polizaHandler.GetAll)
	polizasPub.GET("/:id", polizaHandler.GetByID)

	// ---------- ASEGURADORAS ----------
	asegRepo := aseguradora.NewRepository(db)
	asegSvc := aseguradora.NewService(asegRepo)
	asegHandler := aseguradora.NewHandler(asegSvc)
	asegsPub := api.Group("/aseguradoras")
	asegsPub.GET("", asegHandler.GetAll)
	asegsPub.GET("/:id", asegHandler.GetByID)

	// ---------- HOSPITALIZACIONES ----------
	hospRepo := hospitalizacion.NewRepository(db)
	hospSvc := hospitalizacion.NewService(hospRepo)
	hospHandler := hospitalizacion.NewHandler(hospSvc)
	hospPub := api.Group("/hospitalizaciones")
	hospPub.GET("", hospHandler.GetAll)
	hospPub.GET("/:id", hospHandler.GetByID)

	// ---------- PROTECTED ----------
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())

	// ---------- USUARIO ROLES ----------
	usuarioRolHandler := usuario_rol.NewHandler(db)
	usuarioRolHandler.RegisterRoutes(protected)

	// ---------- MEDICOS (DOCTORES) - protegido para modificación ----------
	medicos := protected.Group("/medicos")
	medicos.POST("", medicoHandlerPub.Create)
	medicos.PUT("/:id", medicoHandlerPub.Update)
	medicos.DELETE("/:id", medicoHandlerPub.Delete)

	// ---------- PACIENTES - protegido para creación/edición ----------
	pacientes := protected.Group("/pacientes")
	pacientes.POST("", pacienteHandlerPub.Crear)
	pacientes.PUT("/:id", pacienteHandlerPub.Actualizar)
	pacientes.DELETE("/:id", pacienteHandlerPub.Eliminar)

	// ---------- CITAS ----------
	citaRepo := cita.NewCitaRepository(db)
	citaService := cita.NewCitaService(citaRepo)
	citaHandler := cita.NewCitaHandler(citaService)
	citaHandler.Registrar(protected)

	// ---------- CONSULTAS ----------
	consultaRepo := consulta.NewConsultaRepository(db)
	consultaService := consulta.NewConsultaService(consultaRepo)
	consultaHandler := consulta.NewConsultaHandler(consultaService)
	consultas := protected.Group("/consultas")
	consultas.GET("", consultaHandler.PorRango)
	consultas.GET("/:id", consultaHandler.Obtener)
	consultas.POST("", consultaHandler.Crear)
	consultas.PUT("/:id", consultaHandler.Actualizar)
	consultas.DELETE("/:id", consultaHandler.Eliminar)

	// ---------- CIRUGIAS ----------
	cirugiaHandler := cirugia.NewHandler(db)
	cirugiaHandler.RegisterRoutes(protected)

	// ---------- DIAGNOSTICOS ----------
	diagnHandler := diagnostico.NewHandler(db)
	diagnHandler.RegisterRoutes(protected)

	// ---------- PRUEBAS (create/update/delete) ----------
	pruebaHandler := prueba.NewHandler(db)
	pruebaHandler.RegisterRoutes(protected)

	// ---------- MEDICAMENTOS CRUD ----------
	meds := protected.Group("/medicamentos")
	meds.POST("", medHandlerPub.Create)
	meds.PUT("/:id", medHandlerPub.Update)
	meds.DELETE("/:id", medHandlerPub.Delete)

	// ---------- SALAS CRUD ----------
	salas := protected.Group("/salas")
	salas.POST("", salaHandler.Create)
	salas.PUT("/:id", salaHandler.Update)
	salas.DELETE("/:id", salaHandler.Delete)

	// ---------- PAGOS ----------
	pagoRepo := pago.NewRepository(db)
	pagoSvc := pago.NewService(pagoRepo)
	pagoHandler := pago.NewHandler(pagoSvc)
	pagoHandler.RegisterRoutes(protected)

	// ---------- PERSONAL ----------
	personalRepo := personal.NewRepository(db)
	personalSvc := personal.NewService(personalRepo)
	personalHandler := personal.NewHandler(personalSvc)
	personalGroup := protected.Group("/personal")
	personalGroup.GET("", personalHandler.GetAll)
	personalGroup.GET("/:id", personalHandler.GetByID)
	personalGroup.POST("", personalHandler.Create)
	personalGroup.PUT("/:id", personalHandler.Update)
	personalGroup.DELETE("/:id", personalHandler.Delete)

	// ---------- RECETAS ----------
	recRepo := receta.NewRepository(db)
	recSvc := receta.NewService(recRepo)
	recHandler := receta.NewHandler(recSvc)
	recGroup := protected.Group("/recetas")
	recGroup.GET("", recHandler.GetAll)
	recGroup.GET("/:id", recHandler.GetByID)
	recGroup.POST("", recHandler.Create)
	recGroup.PUT("/:id", recHandler.Update)
	recGroup.DELETE("/:id", recHandler.Delete)

	// ---------- FACTURACION (facturas) ----------
	facRepo := facturacion.NewRepository(db)
	facSvc := facturacion.NewService(facRepo)
	facHandler := facturacion.NewHandler(facSvc)
	facGroup := protected.Group("/facturas")
	facGroup.GET("", facHandler.GetAllFacturas)
	facGroup.GET("/:id", facHandler.GetFactura)
	facGroup.POST("", facHandler.CreateFactura)
	facGroup.PUT("/:id", facHandler.UpdateFactura)
	facGroup.DELETE("/:id", facHandler.DeleteFactura)

	// ---------- INVENTARIO ----------
	invRepo := inventario.NewRepository(db)
	invSvc := inventario.NewService(invRepo)
	invHandler := inventario.NewHandler(invSvc)
	invGroup := protected.Group("/inventario")
	invGroup.GET("", invHandler.GetAll)
	invGroup.GET("/:id", invHandler.GetByID)
	invGroup.POST("", invHandler.Create)
	invGroup.PUT("/:id", invHandler.Update)
	invGroup.DELETE("/:id", invHandler.Delete)

	// ---------- POLIZAS ----------
	// polGroup := protected.Group("/polizas")
	// polGroup.GET("", polizaHandler.GetAll)
	// polGroup.GET("/:id", polizaHandler.GetByID)
	// polGroup.POST("", polizaHandler.Create)
	// polGroup.PUT("/:id", polizaHandler.Update)
	// polGroup.DELETE("/:id", polizaHandler.Delete)

	// Reemplazamos por: solo endpoints de modificación bajo protected (GETs ya están registrados como públicos arriba)
	polGroup := protected.Group("/polizas")
	polGroup.POST("", polizaHandler.Create)
	polGroup.PUT("/:id", polizaHandler.Update)
	polGroup.DELETE("/:id", polizaHandler.Delete)

	// ---------- TIPO_PRUEBA ----------
	tipoHandler := tipo_prueba.NewHandler(db)
	tipoHandler.RegisterRoutes(protected)

	// ---------- TRATAMIENTOS ----------
	trRepo := tratamiento.NewRepository(db)
	trSvc := tratamiento.NewService(trRepo)
	// servicio simple: exponer CRUD via protected
	trGroup := protected.Group("/tratamientos")
	trGroup.GET("", func(c *gin.Context) {
		list, _ := trSvc.GetAll()
		c.JSON(200, list)
	})
	trGroup.GET("/:id", func(c *gin.Context) {
		// ...simple wrapper...
		id := c.Param("id")
		_ = id
		// existing handler functions can be created if needed
		c.Status(204)
	})
	// roles
	rolHandler := rol.NewHandler(db)
	rolHandler.RegisterRoutes(protected)
}
