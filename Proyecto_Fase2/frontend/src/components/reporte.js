import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Reportes = () => {
    const [imagen, setImagen] = useState('https://w0.peakpx.com/wallpaper/851/500/HD-wallpaper-violet-nebula-starry-sky-universe-space.jpg')
    
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/admin","_self")
    }

    const validar = (data) =>{
        console.log(data)
        setImagen(data.imagen.Imagenbase64)
    }

    const reporteGrafo = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporteGrafo',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const reporteArbol = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const reporteBlockchain = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-bloque',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    return(
        <div className="form-signin1">
            <div className="text-center">
                  <form className="card card-body">
                    <h1 className="cabecera">Reportes Administrador</h1>
                    <br/>
                    <div className="botones">
                        <center><button id='button' onClick={reporteGrafo}>Grafo</button></center>
                        <center><button id='button' onClick={reporteArbol}>Arbol AVL</button></center>
                        <center><button id='button' onClick={reporteBlockchain}>Facturas</button></center>
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