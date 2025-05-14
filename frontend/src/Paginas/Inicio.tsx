import './Estiloinicio.css';

const Inicio = () => (
    <div className="contenedor">
        <div className="contenedorTexto">
            <div className="info_texto">
                <h1 className="titulo">Bienvenido, nuevo portador de la Death Note</h1>
                <p className="subtitulo">
                    Registra nombres, causa y detalles para administrar el destino de los humanos.
                </p>
                <a href="/registro" className="btn">
                    Empezar a escribir
                </a>
            </div>
            <div className="imagenContainer">
                <img
                    src="https://i.pinimg.com/736x/c3/a4/56/c3a456e65739f6bf8a97b38189678b01.jpg"
                    alt="Death Note"
                />
            </div>
        </div> 
    </div>
  );
  

export default Inicio;
