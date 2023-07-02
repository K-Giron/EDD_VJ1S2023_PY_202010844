import React, {useState, useEffect} from 'react';
import '../css/administrador.css'
import 'bootstrap/dist/css/bootstrap.min.css'

import moment from 'moment'

// 18:6:5
export const GenerarFactura = () => {
    //(DD-MM-YY-::HH:MM:SS)
    const fecha = moment().format("DD-MM-yyyy-::hh:mm:ss")
    const idEmpleado = localStorage.getItem("empleado")
    const [cliente, setCliente] = useState('')
    const [pago, setPago] = useState('')

    useEffect(() => {
        peticion();
    },[])

    const peticion = () => {
        fetch("http://localhost:3001/obtenerPedido",{})
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        var idcliente = data.datos.Id_Cliente
        setCliente(idcliente.toString())
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/generarfactura',{
            method: 'POST',
            body: JSON.stringify({
                Timestamp: fecha,
                Biller:    idEmpleado,
                Customer:  cliente,
                Payment:   pago
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => console.log(data))
    }

    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleado","_self")
    }

    return(
        <div className="form-signin1">
            <div className="text-center">
                  <form onSubmit={handleSubmit} className="card card-body">
                  <h1 className="heading">Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                    <label htmlFor="fechaInput">Fecha y tiempo:</label>
                    <label htmlFor="inputEmail" className="visually-hidden">Fecha</label>
                    <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" disabled
                    value={fecha}  
                    autoFocus/>
                    <br/>
                    <label htmlFor="fechaInput">Empleado Logeado:</label>
                    <label htmlFor="inputEmail" className="visually-hidden">Empleado Cobrador</label>
                    <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" disabled
                    value={idEmpleado}  
                    autoFocus/>
                    <br/>
                    <label htmlFor="fechaInput">Id de Cliente:</label>
                    <label htmlFor="inputEmail" className="visually-hidden">Usuario</label>
                    <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" disabled
                    value={cliente}  
                    autoFocus/>
                    <br/>
                    <label htmlFor="fechaInput">Pago:</label>
                    <label htmlFor="inputEmail" className="visually-hidden">Pago</label>
                    <input type="text" id="userI" className="form-control" placeholder="Q0.0" required
                    onChange={e => setPago(e.target.value)} 
                    value={pago}  
                    autoFocus/>
                    <br/>
                    <center><button id='button' type="submit">Generar Pago</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br/>
                    
                  </form>
            </div>
        </div>
    )
}