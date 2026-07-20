package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Paso 1: Definir la estructura del dato (modelo)
type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	EnStock     bool    `json:"en_stock"`
}

// Paso 2: Crear la base de datos simulada en memoria (slice)
var productos = []Producto{
	{ID: 1, Nombre: "Producto 1", Descripcion: "Descripción del producto 1", Precio: 19.99, EnStock: true},
	{ID: 2, Nombre: "Producto 2", Descripcion: "Descripción del producto 2", Precio: 29.99, EnStock: false},
	{ID: 3, Nombre: "Producto 3", Descripcion: "Descripción del producto 3", Precio: 39.99, EnStock: true},
}

// Paso 3: Crear el controlador (handler) que procesa la petición y responde json
func obtenerProductosHandler(w http.ResponseWriter, r *http.Request) {

	// Configurar la cabecera para avisar que enviamos un json
	w.Header().Set("Content-Type", "application/json")

	// Convertir el slice de productos a json
	json.NewEncoder(w).Encode(productos)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// Configurar la cabecera para avisar que enviamos un json
	w.Header().Set("Content-Type", "application/json")

	// Crear un mensaje de estado
	mensaje := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(mensaje)
}

func main() {

	// Paso 4: Configurar el servidor HTTP y asociarla con su controlador
	http.HandleFunc("/api/productos", obtenerProductosHandler)
	http.HandleFunc("/health", healthCheckHandler)

	//Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
