
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link
} from 'react-router-dom'

import Inicio from './Paginas/Inicio'
import Registro from './Paginas/Registro'
import ListaVictimas from './Paginas/ListaVictimas'

import './App.css'

export default function App() {
  return (
    <Router>
      <header>
        <nav className="navbar">
          <ul>
            <li>
              <Link to="/" className="navItem">Inicio</Link>
            </li>
            <li>
              <Link to="/registro" className="navItem">Registrar Persona</Link>
            </li>
            <li>
              <Link to="/victimas" className="navItem">Lista de Víctimas</Link>
            </li>
          </ul>
        </nav>
      </header>

      <main>
        <Routes>
          <Route path="/" element={<Inicio />} />
          <Route path="/registro" element={<Registro />} />
          <Route path="/victimas" element={<ListaVictimas />} />
        </Routes>
      </main>
    </Router>
  )
}
