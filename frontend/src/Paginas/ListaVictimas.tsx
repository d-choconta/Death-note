import './Estilovictimas.css';
import { useEffect, useState } from 'react'
import axios from 'axios'

interface Victim {
  id: number
  full_name: string
  cause_of_death: string
  details: string
  image_url: string
  is_dead: boolean
}

const ListaVictimas = () => {
  const [victims, setVictims] = useState<Victim[]>([])

  const getVictims = async () => {
    try {
      const res = await axios.get<Victim[]>(`${import.meta.env.VITE_API_URL}/victims`)
      console.log("Víctimas recibidas:", res.data)
      setVictims(res.data)
    } catch (error) {
      console.error('Error al cargar víctimas', error)
    }
  }

  useEffect(() => {
    console.log("Ejecutando getVictims") 
    getVictims()
  }, [])

  return (
    <div className="list">
      <h2>Lista de víctimas</h2>
      {victims.map(v => (
        <div className="victim-card" key={v.id}>
          <img src={v.image_url} alt={v.full_name} />
          <div>
            <h3>{v.full_name}</h3>
            <p><strong>Causa:</strong> {v.cause_of_death}</p>
            {v.details && <p><strong>Detalles:</strong> {v.details}</p>}
            <p className="estado">{v.is_dead ? 'Muerto' : 'Vivo'}</p>
          </div>
        </div>
      ))}
    </div>
  )
}

export default ListaVictimas
