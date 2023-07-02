import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'


export const Administrador = () => {
    const reportes = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/reportes","_self")
    }
    const onChange = (e) => {
        e.preventDefault();
        var reader = new FileReader();
        reader.onload = (e) => {
            var obj = JSON.parse(e.target.result);
            console.log(obj.pedidos)
            fetch('http://localhost:3001/cargarpedidos',{
            method: 'POST',
            body: JSON.stringify({
                Pedidos: obj.pedidos
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))


        };
        reader.readAsText(e.target.files[0]);
    }

    const onChange1 = (e, file1) => {
        var file = file1 || e.target.files[0];
        console.log(file.name);
        fetch('http://localhost:3001/cargarempleados',{
            method: 'POST',
            body: JSON.stringify({
                Nombre: file.name
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) => {
        console.log(data)
        
    }

    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/","_self")
    }

    return(
        <div className="form-signin1">
            <div className="text-center">
                    <form className="card card-body">
                    <h1 className="heading">Administrador</h1>
                    <p className="cabecera" >Cargar Archivos</p>
                    <br/>
                    <div className="input-group mb-3">
                        <label className="input-group-text">Cargar Pedidos</label>
                        <input class="form-control form-control-lg" id="formFileLg"
                        type="file"
                        accept="application/json"
                        onChange={onChange}/>
                    </div>
                    <br/>
                    <div className="input-group mb-3">
                        <label className="input-group-text">Cargar Empleados</label>
                        <input class="form-control form-control-lg" id="formFileLg"
                        type="file"
                        accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
                        onChange={onChange1}/>
                    </div>
                    <br/>
                    <button id='button' onClick={reportes}>Reportes</button>
                    <button id='button' onClick={salir}>Cerrar Sesión</button>
                    </form>
            </div>
        </div>
    );
}