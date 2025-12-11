package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/repository"
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/service"
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/handler"
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/routes"
)

func main() {
    // Cargar variables de entorno desde .env
    _ = godotenv.Load()

    // Construir DSN desde variables de entorno
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    // Conectar a la base de datos
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }
    fmt.Println("Base de datos conectada exitosamente")

    // Migrar modelos
    if err := db.AutoMigrate(&models.Area{}, &models.Person{}); err != nil {
        log.Fatal("Error en migración de modelos:", err)
    }

    // Inicializar repositorios, servicios y handlers
    areaRepo := repository.NewAreaRepository(db)
    personRepo := repository.NewPersonRepository(db)
    areaService := service.NewAreaService(areaRepo)
    personService := service.NewPersonService(personRepo)
    areaHandler := handler.NewAreaHandler(areaService)
    personHandler := handler.NewPersonHandler(personService)

    // Configurar rutas
    router := routes.SetupRoutes(areaHandler, personHandler)

    // Obtener puerto del servicio
    port := os.Getenv("SERVICE_PORT")
    if port == "" {
        port = "8080"
    }

    // Iniciar el servidor
    if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
        log.Fatal("Error al iniciar el servidor:", err)
    }
}