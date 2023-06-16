# EDD_VJ1S2023_PY_202010844
## Manual Técnico
hola
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
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/c53b11e4-a361-4205-ba5b-4196d8154fb1" width="600">
<br>

En los reportes existen las pilas que manejan los pedidos, las listas circulares que maneja los clientes, la lista doble enlazada que maneja las imagenes, la enlazada simple que almacena a los empleados, por último la cola de clientes que esperan a ser atendidos.


