# Backend GO - Registro de Áreas y Personas

## Descripción

API REST desarrollada en Go para el registro de áreas de trabajo y personas asociadas a cada área. Utiliza MySQL como base de datos y sigue una arquitectura limpia con capas de handler, service y repository.

**Frontend disponible:** https://github.com/KarlaR3it/PruebaEureka-Frontend.git

## Características principales

- **Gestión de Personas**: Crear y listar personas con validación de email único
- **Gestión de Áreas**: Crear y listar áreas con validación de nombre único
- **Estadísticas**: Consultar cantidad de personas por área
- **Validaciones**: Todos los campos obligatorios con mensajes claros de error
- **Tests Unitarios**: Cobertura completa de handlers, services y repositories
- **Arquitectura Limpia**: Separación clara de responsabilidades

## Requisitos

- Go 1.20+
- Docker y Docker Compose
- (Opcional) DBeaver o MySQL Workbench

## Instalación y Ejecución

1. Clona el repositorio:

   ```sh
   git clone https://github.com/KarlaR3it/PruebaEureka-Backend
   cd Backend-GO
   ```

2. Crea el archivo de variables de entorno:

   Usa `.env.example` como referencia.

3. Levanta los servicios:

   ```sh
   docker compose up -d --build
   ```

## Comandos Útiles de Docker

```sh
# Iniciar servicios en segundo plano
docker compose up -d

# Reconstruir e iniciar servicios
docker compose up -d --build

# Detener y limpiar contenedores y volúmenes
docker compose down -v

# Ver logs en tiempo real
docker compose logs -f

# Ver estado de los contenedores
docker compose ps

# Acceder al contenedor del backend
docker compose exec backend sh

# Acceder a la base de datos MySQL
docker compose exec db mysql -u root -p
```

## Ejecutar Tests Unitarios

El proyecto incluye tests unitarios completos para ambas entidades (Area y Person) en las capas de handler, service y repository.

```sh
# Ejecutar todos los tests
go test -v ./...

# Ejecutar tests de un paquete específico
go test -v ./test

# Ejecutar un test específico
go test -v ./test -run TestPersonHandler_Create
```

**Nota:** Los tests usan mocks (testify/mock y sqlmock) y no requieren una base de datos real.

## Colección de Postman

Puedes importar y probar todos los endpoints usando la colección incluida en la raíz del proyecto:

- `PruebaEureka.postman_collection.json`

Importa este archivo en Postman para tener acceso rápido a todas las peticiones del backend.

## Estructura del Proyecto

```
Backend-GO/
├── cmd/
│   └── main.go              # Punto de entrada de la aplicación
├── internal/
│   ├── models/              # Modelos de datos (Area, Person)
│   ├── repository/          # Capa de acceso a datos
│   ├── service/             # Lógica de negocio
│   ├── handler/             # Controladores HTTP
│   └── routes/              # Configuración de rutas
├── test/                    # Tests unitarios
├── docker-compose.yml       # Configuración de Docker
├── Dockerfile               # Imagen del backend
├── .env.example             # Variables de entorno de ejemplo
└── README.md                # Este archivo
```
