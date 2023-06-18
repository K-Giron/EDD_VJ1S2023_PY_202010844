package main

import (
	"PF1_202010844/estructuras"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ListaSimple = &estructuras.ListaSimple{Inicio: nil, Longitud: 0}
var ListaDoble = &estructuras.ListaDoble{Inicio: nil, Longitud: 0}
var ListaCircular = &estructuras.ListaCircular{Inicio: nil, Longitud: 0}
var Cola = &estructuras.Cola{Primero: nil, Longitud: 0}
var Pila = &estructuras.Pila{Primero: nil, Longitud: 0}
var matriz = &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
var matriz_csv = &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--------------------------------------")
	fmt.Println("Proyecto Fase 1 - Estructuras de Datos")
	fmt.Println("Bienvenido al sistema de inicio de sesión")
	ListaSimple.Insertar("202010844", "Jorge", "Estudiante", "1234")
	// Loop hasta que se inicie sesión correctamente o se seleccione la opción de salir
	for {
		fmt.Println("----------------------------------")
		fmt.Println("\nPor favor, elija una opción:")
		fmt.Println("1. Iniciar sesión")
		fmt.Println("2. Salir")
		fmt.Println("----------------------------------")
		fmt.Print("Opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("\nIngrese sus credenciales:")
			fmt.Print("ID de usuario: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Contraseña: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)
			//Si es admin
			if id == "admin_202010844" && password == "admin" {
				fmt.Println("Inicio de sesión exitoso. ¡Bienvenido, administrador!")
				menuAdmin()
				return
			}

			// Verificar las credenciales ingresadas
			currentNode := ListaSimple.Inicio
			for currentNode != nil {
				empleado := currentNode.Empleado
				if empleado.IdEmpleado == id && empleado.Password == password {
					fmt.Printf("Inicio de sesión exitoso. ¡Bienvenido, %s!\n", empleado.Nombre)
					menuEmpleado(empleado.IdEmpleado)
					return
				}
				currentNode = currentNode.Siguiente
			}

			fmt.Println("Credenciales incorrectas. Por favor, intente nuevamente.")
		case "2":
			fmt.Println("Finalizando programa...")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Por favor, elija nuevamente.")
		}
	}
}

func menuAdmin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nPor favor, elija una opción:")
		fmt.Println("1. Cargar archivo de empleados")
		fmt.Println("2. Cargar archivo de imágenes")
		fmt.Println("3. Cargar archivo de Clientes")
		fmt.Println("4. Actualizar Cola")
		fmt.Println("5. Reportes Estructuras")
		fmt.Println("6. Cerrar sesión")
		fmt.Println("----------------------------------")
		fmt.Print("Opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			estructuras.LeerEmpleados("archivos_prueba/empleados.csv", ListaSimple)
			fmt.Println("Archivo cargado exitosamente")
		case "2":
			estructuras.LeerImagen("archivos_prueba/imagenes.csv", ListaDoble)
			fmt.Println("Archivo cargado exitosamente")
		case "3":
			estructuras.LeerClientes("archivos_prueba/clientes_registrados.csv", ListaCircular)
			fmt.Println("Archivo cargado exitosamente")
		case "4":
			clientesNuevos := estructuras.LeerCola("archivos_prueba/clientes_cola.csv", Cola)
			if clientesNuevos.Longitud > 0 {
				aux := clientesNuevos.Inicio
				for {
					ListaCircular.Insertar(aux.Cliente.IdCliente, aux.Cliente.Nombre)
					fmt.Println("Se inserto el cliente: ", aux.Cliente.IdCliente, aux.Cliente.Nombre, " a la lista circular de clientes")
					aux = aux.Siguiente
					if aux == clientesNuevos.Inicio {
						break
					}
				}
			}
			fmt.Println("Archivo cargado exitosamente")
		case "5":
			ListaSimple.ReporteEmpleados()
			ListaDoble.ReporteImagen()
			ListaCircular.ReporteClientes()
			Cola.Graficar()
			Pila.Graficar()
			Pila.CrearContenidoJSON()
			fmt.Println("Reportes generados exitosamente en la carpeta Reportes")
		case "6":
			fmt.Println("Sesión cerrada exitosamente")
			main()
		default:
			fmt.Println("Opción no válida. Por favor, elija nuevamente.")
		}
	}
}

func menuEmpleado(idEmpleado string) {
	fmt.Println("Bienvenido al sistema de empleados", idEmpleado)
	reader := bufio.NewReader(os.Stdin)
	nombreImagen := ""
	idCliente := ""
	imagen := "ave"
	for {
		fmt.Println("\nPor favor, elija una opción:")
		fmt.Println("1. Ver imágenes Cargadas")
		fmt.Println("2. Realizar pedido")
		fmt.Println("3. Cerrar sesión")
		fmt.Println("----------------------------------")
		fmt.Print("Opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("Imprimiendo imágenes cargadas")
			ListaDoble.Mostrar()
			fmt.Println("Elige una imagen ingresando su nombre:")
			fmt.Print("Opción: ")
			nombreImagen, _ = reader.ReadString('\n')
			nombreImagen = strings.TrimSpace(nombreImagen)

		case "2":
			fmt.Println("Imprimiendo el nombre de la imagen: ", nombreImagen)
			idCliente = Cola.Descolar()
			fmt.Println("El cliente que realizo el pedido es: ", idCliente)
			Pila.Push(nombreImagen, idCliente)
			matriz_csv.LeerInicial("archivos_prueba/"+imagen+"/inicial.csv", imagen)
			matriz_csv.GenerarImagen(imagen)

		case "3":
			fmt.Println("Sesión cerrada exitosamente")
			main()
		default:
			fmt.Println("Opción no válida. Por favor, elija nuevamente.")
		}
	}
}

/*fmt.Println("Hello, World! Esto es una prueba de go")
ListaSimple := &estructuras.ListaSimple{Inicio: nil, Longitud: 0}
ListaSimple.Insertar("202010844", "Jorge", "Estudiante", "1234")
ListaSimple.Insertar("202010845", "Luis", "Trabajador", "12345")
ListaSimple.Insertar("202010846", "Pedro", "Limpiador", "123456")
ListaSimple.Insertar("202010847", "Guicho", "Gerente", "1234567")
ListaSimple.Insertar("202010848", "Kevin", "Estudiante", "12345678")

estructuras.LeerEmpleados("archivos_prueba/empleados.csv", ListaSimple)
ListaSimple.ReporteEmpleados()

fmt.Println("-----------------------------------------")
ListaDoble := &estructuras.ListaDoble{Inicio: nil, Longitud: 0}
ListaDoble.Insertar("imagen1", "1")
ListaDoble.Insertar("imagen2", "2")
ListaDoble.Insertar("imagen3", "3")
ListaDoble.Insertar("imagen4", "4")
estructuras.LeerImagen("archivos_prueba/imagenes.csv", ListaDoble)
ListaDoble.ReporteImagen()

fmt.Println("-----------------------------------------")
ListaCircular := &estructuras.ListaCircular{Inicio: nil, Longitud: 0}
ListaCircular.Insertar("202010844", "Matias")
ListaCircular.Insertar("202010845", "Ana")
ListaCircular.Insertar("202010846", "Susana")
ListaCircular.Insertar("202010847", "jeremias")
estructuras.LeerClientes("archivos_prueba/clientes_registrados.csv", ListaCircular)
ListaCircular.Mostrar()
ListaCircular.ReporteClientes()

fmt.Println("-----------------------------------------")
Cola := &estructuras.Cola{Primero: nil, Longitud: 0}
Cola.Encolar("4343", "Cristian")
Cola.Encolar("4344", "Juan")
Cola.Encolar("4345", "Pedro")
Cola.Encolar("4346", "Maria")
clientesNuevos := estructuras.LeerCola("archivos_prueba/clientes_cola.csv", Cola)
Cola.Graficar()

fmt.Println("-----------------------------------------")
if clientesNuevos.Longitud > 0 {
	aux := clientesNuevos.Inicio
	for {
		ListaCircular.Insertar(aux.Cliente.IdCliente, aux.Cliente.Nombre)
		fmt.Println("Se inserto el cliente: ", aux.Cliente.IdCliente, aux.Cliente.Nombre, " a la lista circular de clientes")
		aux = aux.Siguiente
		if aux == clientesNuevos.Inicio {
			break
		}
	}
}

fmt.Println("-----------------------------------------")
fmt.Println("Mostrando la cola")
Cola.Mostrar()
ListaCircular.Mostrar()
*/
