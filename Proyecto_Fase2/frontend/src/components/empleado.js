import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'



export const Empleado = () => {
    const [imagen, setImagen] = useState('https://icones.pro/wp-content/uploads/2021/02/symbole-masculin-icone-l-utilisateur-violet.png')
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/","_self")
    }

    const validar = (data) =>{
        console.log(data)
    }

    const aplicarFiltros = async(e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/filtros","_self")
    }

    const generarFacturas = async(e) => {
        e.preventDefault();
        window.open("/factura","_self")
    }
    
    const verFacturas = async(e) => {
        e.preventDefault();
        window.open("/verfactura","_self")
    }

    return(
        <div className="form-signin1">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 className="cabecera">Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <div className="botones">
                        <center><button id='button' onClick={aplicarFiltros}>Aplicacion de Filtros</button></center>
                        <br/>
                        <center><button id='button' onClick={generarFacturas}>Generar Facturas</button></center>
                        <br/>
                        <center><button id='button' onClick={verFacturas}>Ver Facturas</button></center>
                        <br/>
                        <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                        <br/>
                    </div>
                    <div className="imagen">
                        <center><img src={imagen} width="250" height="250" alt='some value' /></center>
                        <br/>
                    </div>
                  </form>
            </div>
          </div>
    );
}