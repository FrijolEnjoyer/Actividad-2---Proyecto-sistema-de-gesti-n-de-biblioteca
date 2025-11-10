# Sistema de Gestión de Biblioteca (Etapa 2)

Prototipo funcional usando estructuras de datos lineales en Go (backend) y React (frontend), orquestado con Docker Compose.

## Estructuras de datos
- Árboles binarios de búsqueda (`internal/ds/tree.go`):
  - Libros y usuarios ordenados por ID para inserción/búsqueda/eliminación en O(log n) promedio.
  - Préstamos activos indexados por ID de libro para validar disponibilidad y devoluciones.
- Pila (`internal/ds/stack.go`): historial de operaciones recientes.
- Cola (`internal/ds/queue.go`): (etapa anterior) solicitudes en secuencia, conservada como referencia.
- Arreglo (`internal/ds/array.go`): destacados con capacidad fija.

## Operaciones
- Libros: registrar, listar en orden, buscar por texto, prestar, devolver, eliminar (impide borrar si está prestado).
- Usuarios: registrar, listar en orden, eliminar (impide borrar si tiene préstamos activos).

## Arquitectura
- Backend Go: `backend/`
- Frontend React (Vite): `frontend/`
- Docker: `docker-compose.yml`

## Cómo ejecutar (Docker)
1. Requisitos: Docker y Docker Compose
2. Ejecutar:
```
docker compose up --build
```
3. Frontend: http://localhost:5173
4. Backend API: http://localhost:8080/api

## Cómo ejecutar local (opcional)
- Backend:
```
cd backend
go run ./cmd/server
```
- Frontend:
```
cd frontend
npm install
npm run dev
```

## Endpoints principales
- `POST /api/users` crear usuario
- `GET /api/users` listar usuarios
- `DELETE /api/users?id=USER_ID` eliminar usuario
- `POST /api/books` crear libro
- `GET /api/books` listar libros
- `GET /api/books/search?q=texto` buscar por título o autor
- `DELETE /api/books?id=BOOK_ID` eliminar libro (si no está prestado)
- `POST /api/loans/borrow` prestar
- `POST /api/loans/return` devolver

## Pruebas
- Backend (estructuras y servicio):
```
cd backend
go test ./...
```

## Decisiones de diseño
- Se migró el modelo central a árboles de búsqueda binaria para optimizar la gestión de libros, usuarios y préstamos activos.
- Se mantienen estructuras lineales para historial, destacados y como referencia de la etapa previa.
- Sin base de datos: almacenamiento en memoria con estructuras diseñadas.
- CORS habilitado para React.
- UI con tema oscuro, tarjetas y botones con estados. Listas con recarga automática tras crear elementos (hot reload) y tras prestar/devolver.
- Nuevas operaciones de eliminación: `RemoveBook` evita borrar si el libro está prestado; `RemoveUser` elimina por ID.

## Próximos pasos
- Persistencia
- Autenticación básica
- Paginación y validaciones más estrictas
- Balanceo del BST (AVL/Red-Black) si el patrón de inserciones produce degradación a O(n).

## Ejemplos (cURL)
- Crear usuario:
```
curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"id":"u1","name":"Ana"}'
```
- Eliminar usuario:
```
curl -X DELETE "http://localhost:8080/api/users?id=u1"
```
- Crear libro:
```
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d '{"id":"b1","title":"Go","author":"Gopher","isbn":"123"}'
```
- Buscar libro:
```
curl "http://localhost:8080/api/books/search?q=go"
```
- Eliminar libro (solo si no está prestado):
```
curl -X DELETE "http://localhost:8080/api/books?id=b1"
```
