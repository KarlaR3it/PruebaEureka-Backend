# Backend GO - Registro de Áreas y Personas

## Descripción

API REST desarrollada en Go para el registro de áreas de trabajo y personas asociadas a cada área. Utiliza MySQL como base de datos. Permite crear y consultar áreas y personas, con validaciones y mensajes claros de éxito o error. El proyecto incluye algunos tests unitarios. Si deseas una interfaz de usuario, puedes clonar el frontend de este proyecto aquí: https://github.com/KarlaR3it/PruebaEureka-Frontend.git

## Características principales

- Crear y listar personas (nombre, correo electrónico único, área de trabajo).
- Crear y listar áreas de trabajo (ej: Ventas, Recursos Humanos).
- Consultar la cantidad de personas por área.
- Todos los campos son obligatorios.
- Respuestas claras de éxito o error en cada operación.

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

3. Levanta los servicios (backend y base de datos):

   ```sh
   docker compose up -d --build
   ```

   Nota: Usa docker compose up -d si solo quieres levantar los servicios sin reconstruir la imagen.

4. (Opcional) Detener y limpiar los contenedores y volúmenes:
   ```sh
   docker compose down -v
   ```

## Ejecutar Tests Unitarios

El proyecto incluye tests unitarios para las capas de repositorio, servicio y handler de la entidad Area.

Para ejecutar todos los tests:

```sh
go test -v ./...
```

Para ejecutar test de una función específica:

```sh
go test -v ./test -run TestAreaHandler_Create
```

Los tests usan mocks (testify/mock y sqlmock) y no requieren una base de datos real.

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
