# EDD_VJ1S2023_PY_202010844
## Manual Técnico

El manual técnico tiene como finalidad dar a conocer al lector que pueda requerir hacer modificaciones futuras al software el desarrollo de la aplicación indicando el IDE utilizado para su creación, versión, requerimientos del sistema, etc.
La aplicación tiene como objetivo realizar un sistema que registre los empleados, clientes, colas y pedidos de imagenes para la creación de las mismas imágenes. 

### Estructuras de datos para cada objeto
Se utilizan diferentes estructuras de datos para guardar cada tipo de dato.


### Carga de datos
La manera de acceder a los datos es con archivos de tipo .csv . En estos archivos se encuentran todos los registros los cuales deben ser registrados y posteriormente deben ser guardados en las estructuras de datos establecidos para cada tipo de objeto. Se presenta la carpeta de los archivos .csv


### Funcionalidad de estructuras
En la carpeta de estructuras se encuentra cada módulo que contiene todos los métodos y funciones para la funcionalidad de cada estructura de datos, también esta el módulo de Nodo en el cual esta la estructura de cada uno de los nodos para las listas.

se encuentra también el archivo CrearArchivos que se relaciona directamente a la parte de la generación de reportes en cada una de las estructuras, en cada módulo se tiene la estructura del tipo de dato a almacenar.


----
## Manual de Usuario
Un Manual de usuario permite al usuario la experiencia, familiarización y visualización de la aplicación y de los diferentes componentes que se utilizan en la aplicación con el fin de utilizar de mejor manera el mismo. La aplicación cuenta con todo el uso por medio de consola, recibiendo y mostrando datos con un sistema amigable y comprensible para un mejor uso del sistema.

### Sistema de Login

El menú inicial en el sistema es un sistema de Login, este menú de login funciona tanto para el usuario admin establecido estáticamente con las credenciales id: admin_202010844, password: admin, así mismo con los usuarios ingresados de manera dinámica una vez ya cargados datos con usuario admin.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/ea5a3522-afcb-4f3f-a23b-51386e89cfc6" width="600">
<br>
### Menú de administrador

El siguiente menú que aparece en el sistema es el menú administrador, una vez ya se haya logeado el administrador pues se le da acceso a este menú el cual tiene varias opciones de carga de archivos en la cual se lee el archivo automáticamente y se carga a las listas correspondientes también de manera automática, cabe recalcar que dichos archivos deben estar en una carpeta llamada "archivos_prueba" y con los nombres ya establecidos anteriormente. Las 3 primeras acciones cargan los achivos correspondientes y la opción 4 carga una cola y actualiza la lista circular de clientes, la última de las opciones cierra la sesión y devuelve al usuario al login.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/092cac4d-3d35-4dbf-a1db-a275ef45a071" width="600">
<br>
### Menú de Empleados

El menú de empleados aparece solamente cuando se logean con los credenciales de un empleado, aparecen las opciones de ver imágenes cargadas la cual lista el nombre de las imágenes con sus capas y se debe elegir una ingresando el nombre, posteriormente, en la opción 2 del menú que se realiza el pedido se extrae el dato id de la cola que esta primero para guardar ese dato junto al nombre de la imagen. Por último tiene la opción de cerrar sesión que regresa al Login.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/c53b11e4-a361-4205-ba5b-4196d8154fb1" width="600">
<br>

### Reportes

Se recalca que los reportes se crean en una carpeta llamada reportes dentro de la raíz principal y se crea su respectivo dot con su jpg de cada una de las listas que son utilizadas en el programa.

<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/966b87af-2ddf-48da-b010-150130e2badb" width="400">
<br>

En los reportes existen las pilas que manejan los pedidos.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/e2c9848d-5a80-497e-bd39-413170d0a113" width="400">
<br>

Las listas circulares que maneja los clientes.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/fdc2c3c2-72a0-4ac8-b89b-3e31053e4a00" width="400">
<br>

La lista doble enlazada que maneja las imagenes.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/738e53e4-45cd-4e53-b65e-7d7d2a3ead13" width="400">
<br>
La enlazada simple que almacena a los empleados.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/bbcf1d30-349a-4c34-a8e9-4368b04f7ec1" width="400">
<br>
La cola de clientes que esperan a ser atendidos.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/05eb9670-f99e-4d71-a59e-0a0cdedfe7d7" width="400">
<br>
El sistema se utiliza con una matriz dispersa para una funcionalidad la generación de imágenes.


