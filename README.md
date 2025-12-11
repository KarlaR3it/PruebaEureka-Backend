# Backend GO - Registro de Áreas y Personas

## Descripción

API REST desarrollada en Golang para el registro de personas y áreas de trabajo, usando MySQL como base de datos. Permite crear y consultar personas y áreas, con validaciones y mensajes claros de éxito o error.

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

2. Copia el archivo de variables de entorno:

   ```sh
   cp .env.example .env
   ```

3. Levanta los servicios (backend y base de datos):

   ```sh
   docker compose up -d --build
   ```

4. (Opcional) Limpia los volúmenes y datos:
   ```sh
   docker compose down -v
   ```

## Endpoints principales

- **Áreas**

  - `POST   /areas`    Crear área
  - `GET    /areas`    Listar áreas
  - `GET    /areas/count` Contar personas por área

- **Personas**
  - `POST   /persons`   Crear persona
  - `GET    /persons`   Listar personas

## Ejemplos de uso

### Crear un área

```json
{
  "name": "Recursos Humanos"
}
```

### Crear una persona

```json
{
  "name": "Juan Pérez",
  "email": "juan.perez@email.com",
  "area_id": 1
}
```

## Probar con Postman

- Base URL: `http://localhost:8084`
- Ejemplo: `GET http://localhost:8084/areas`

## Colección de Postman

Puedes importar y probar todos los endpoints usando la colección incluida en la raíz del proyecto:

- `PruebaEureka.postman_collection.json`

Importa este archivo en Postman para tener acceso rápido a todas las peticiones del backend.
