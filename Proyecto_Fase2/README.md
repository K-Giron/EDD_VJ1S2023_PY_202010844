# Proyecto Fase 2
## Manual Técnico

El manual técnico tiene como finalidad dar a conocer al lector que pueda requerir hacer modificaciones futuras al software el desarrollo de la aplicación indicando el IDE utilizado para su creación, versión, requerimientos del sistema, etc.
La aplicación tiene como objetivo realizar un sistema que registre los empleados, los pedidos y a base de ello junto con imágenes crear un sistema de filtros para las imágenes y generación de facturas para los pedidos.

### Estructuras de datos para cada objeto
Se utilizan diferentes estructuras de datos para guardar cada tipo de dato.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/da85cae1-6a4c-469d-975e-b7846d258582">
<br>

### Carga de datos
La manera de acceder a los datos es con archivos de tipo .csv y archivo de tipo JSON . En estos archivos se encuentran todos los registros los cuales deben ser registrados y posteriormente deben ser guardados en las estructuras de datos establecidos para cada tipo de objeto. Se presenta la carpeta de los archivos .csv
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/72cb42ec-7283-487b-9545-651e4535e720">
<br>


### Módulo de Peticiones
En este módulo se encuentran estructuras utilizadas para las peticiones que se reciben desde el front y así mismo para la funcionalidad del backend
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/16d24e83-c667-429f-aa00-ae81dad0f826">
<br>

### Módulos
Los módulos contienen la estructura a guardar de los objetos y también el de la estructura en específico del módulo que se este trabajando además de todas las funcionalidades detalladas en las siguientes imágenes.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/9af4d586-b43a-4586-8a8b-1933ae57f665" width="600">
<br>

### Clase Módulo
En este módulo se encuentra toda la lógica y funcionalidad del sistema, se encuentran las peticiones post y get que recibe del frontend y así mismo todo lo relacionado con las estructuras.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/0674dc10-3a70-412b-9ee9-a106e0ef7c22" width="600">
<br>

### Inicialización de estructuras
Se muestra la inicialización de las estructuras de maneral global para trabajarse posteriormente en los métodos.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/b945993a-10d3-43fe-9ffd-e33766b2de90" width="600">
<br>



## Manual de Usuario
Un Manual de usuario permite al usuario la experiencia, familiarización y visualización de la aplicación y de los diferentes componentes que se utilizan en la aplicación con el fin de utilizar de mejor manera el mismo. La aplicación cuenta con todo el uso por medio de consola, recibiendo y mostrando datos con un sistema amigable y comprensible para un mejor uso del sistema.

### Sistema de Login

El menú inicial en el sistema es un sistema de Login, este menú de login funciona tanto para el usuario admin establecido estáticamente con las credenciales id: ADMIN_202010844, password: admin, así mismo con los usuarios ingresados de manera dinámica una vez ya cargados datos con usuario admin.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/5ea98878-3e0c-43c1-b834-a5476f31a36f" width="600">
<br>
### Menú de administrador

El siguiente menú que aparece en el sistema es el menú administrador, una vez ya se haya logeado el administrador pues se le da acceso a este menú el cual tiene varias opciones de carga de archivos en la cual se lee el archivo automáticamente y se carga a las estructuras correspondientes también de manera automática. Las opciones de carga de pedidos es con archivo JSON y los empleados con el csv como se manejaba en la anterior forma, tiene un botón de reportes que envía a una vista de reportes de algunas estructuras.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/320043b1-9556-4a17-bdb5-71571c175fc5" width="600">
<br>
### Menú de Empleados

El menú de empleados aparece solamente cuando se logean con los credenciales de un empleado, aparecen las opciones de aplicación de filtros aplicados la cual direcciona a otra vista para la aplicacion de los filtros, la opción de generar factura la cual envía a una vista de generación de factura y la opción de ver facturas la cual envía a una tabla para ver las facturas generadas.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/75eda3f5-589d-44f5-8b21-f47076ac0584" width="600">
<br>

### Reportes

Los reportes se muestra en el espacio de la imágen estática.

<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/fb3f553e-6e9a-47cb-9398-8dcf9626716f" width="400">
<br>
### Vista Filtros
La vista de filtros se muestran una selección para posteriormente generar o aplicar el filtro el cual se genera dentro de una carpeta dentro del backend
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/1187b61b-ddf3-4ddd-b96a-ff56db7ef917" width="400">
<br>

### Vista  Facturas
La vista de facturas muestra todas las facturas generadas por el usuario.
<br>
<img align='center' src="https://github.com/K-Giron/EDD_VJ1S2023_PY_202010844/assets/68490914/92ea67b6-9a12-40bd-9bc3-0daf237f0ce3" width="400">
<br>
