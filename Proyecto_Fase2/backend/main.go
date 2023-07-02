package main

import (
	"Fase2/estructuras/ArbolAVL"
	"Fase2/estructuras/ColaPedidos"
	"Fase2/estructuras/Facturas"
	"Fase2/estructuras/Grafo"
	"Fase2/estructuras/Lista"
	"Fase2/estructuras/Matriz"
	"Fase2/estructuras/Peticiones"
	"Fase2/estructuras/TablaHash"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleado *Lista.ListaSimple
var ArbolPedidos *ArbolAVL.Arbol
var MatrizOriginal *Matriz.Matriz
var MatrizFiltro *Matriz.Matriz
var PedidosCola *ColaPedidos.Cola
var FacturasRealizadas *Facturas.BlockChain
var VerFacturasRealizadas *TablaHash.TablaHash
var FiltrosColocados string
var EmpleadoLogeado string
var GrafosEmpleados map[string]Grafo.Grafo
var GenerarGrafo *Grafo.Grafo

func main() {
	ListaEmpleado = &Lista.ListaSimple{Inicio: nil, Longitud: 0}
	ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}
	MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	FacturasRealizadas = &Facturas.BlockChain{Bloques_Creados: 0}
	VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
	PedidosCola = &ColaPedidos.Cola{Primero: nil, Longitud: 0}
	GenerarGrafo = &Grafo.Grafo{Principal: nil}
	FiltrosColocados = ""
	EmpleadoLogeado = ""

	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", func(c *fiber.Ctx) error {
		var usuario Peticiones.Login
		c.BodyParser(&usuario)
		if usuario.Username == "ADMIN_202010844" && usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status": "400",
			})
		} else {
			if ListaEmpleado.Inicio != nil {
				if ListaEmpleado.Buscar(usuario.Username, usuario.Password) {
					VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
					VerFacturasRealizadas.NewTablaHash()
					EmpleadoLogeado = usuario.Username
					return c.JSON(&fiber.Map{
						"status": "200",
					})
				}
			}
		}
		return c.JSON(&fiber.Map{
			"status": "404",
		})
	})

	app.Post("/cargarempleados", func(c *fiber.Ctx) error {
		var nombreArchivo Peticiones.Archivo
		c.BodyParser(&nombreArchivo)
		Lista.LeerEmpleados(nombreArchivo.Nombre, ListaEmpleado)
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Post("/cargarpedidos", func(c *fiber.Ctx) error {
		var pedidosNuevos Peticiones.ArbolPeticion
		c.BodyParser(&pedidosNuevos)
		fmt.Println(pedidosNuevos)
		for i := 0; i < len(pedidosNuevos.Pedidos); i++ {
			ArbolPedidos.InsertarElemento(pedidosNuevos.Pedidos[i].Id_Cliente, pedidosNuevos.Pedidos[i].Nombre_Imagen)
		}
		ArbolPedidos.Graficar()
		ArbolPedidos.RecorridoInorden(ArbolPedidos.Raiz, PedidosCola)
		return c.JSON(&fiber.Map{
			"status": 200,
			"cola":   PedidosCola,
		})
	})

	app.Get("/reporte-arbol", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reportes/arbolAVL.jpg"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Get("/reporteGrafo", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reportes/grafo.jpg"}
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 400,
			})
		}
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Post("/aplicarfiltro", func(c *fiber.Ctx) error {
		MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		var tipo Peticiones.PeticionFiltro
		tipo.NombreImagen = PedidosCola.Primero.Pedido.Nombre_Imagen
		imagen := tipo.NombreImagen
		c.BodyParser(&tipo)

		switch tipo.Tipo {
		case 1:
			MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
			MatrizFiltro.Negativo(imagen + "Negativo")
			FiltrosColocados += "Negativo, "
		case 2:
			MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
			MatrizFiltro.EscalaGrises(imagen + "Grises")
			FiltrosColocados += "Grises, "
		case 3:
			MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
			MatrizFiltro.RotacionX()
			MatrizFiltro.GenerarImagen(imagen + "RX")
			FiltrosColocados += "Eje X, "
		case 4:
			MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
			MatrizFiltro.RotacionY()
			MatrizFiltro.GenerarImagen(imagen + "RY")
			FiltrosColocados += "Eje Y, "
		case 5:
			MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
			MatrizFiltro.RotacionDoble()
			MatrizFiltro.GenerarImagen(imagen + "RDoble")
			FiltrosColocados += "Doble,  "
		}

		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Get("/obtenerPedido", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"datos": PedidosCola.Primero.Pedido,
		})
	})

	app.Post("/generarfactura", func(c *fiber.Ctx) error {
		var nuevoBloque Peticiones.BloquePeticion
		c.BodyParser(&nuevoBloque)
		FacturasRealizadas.InsertarBloque(nuevoBloque.Timestamp, nuevoBloque.Biller, nuevoBloque.Customer, nuevoBloque.Payment)
		/*Ingresar al grafo, tomar los valores de nuevoBloque.Biller, nuevoBloque.Customer, PedidosCola.Primero.Pedido.Nombre_Imagen,Filtros_colocados */
		GenerarGrafo.InsertarValores(EmpleadoLogeado, strconv.Itoa(PedidosCola.Primero.Pedido.Id_Cliente), PedidosCola.Primero.Pedido.Nombre_Imagen, FiltrosColocados)
		GenerarGrafo.Reporte()
		PedidosCola.Descolar()
		VerFacturasRealizadas.NewTablaHash()
		FacturasRealizadas.InsertarTabla(VerFacturasRealizadas, EmpleadoLogeado)
		MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}

		return c.JSON(&fiber.Map{
			"datos": FacturasRealizadas.Bloques_Creados,
		})
	})

	app.Get("/facturaempleado", func(c *fiber.Ctx) error {

		return c.JSON(&fiber.Map{
			"status":  200,
			"factura": VerFacturasRealizadas.Tabla,
		})
	})

	app.Listen(":3001")
}
