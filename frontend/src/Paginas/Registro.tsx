import './Estiloregistro.css'
import { useState, useEffect, useRef } from 'react'
import axios from 'axios'

interface Victim {
  full_name: string
  image_url: string
  cause_of_death: string
  details: string
  is_dead: boolean
}

const Registro = () => {
  const [form, setForm] = useState<Victim>({
    full_name: '',
    image_url: '',
    cause_of_death: '',
    details: '',
    is_dead: true
  })

  const [errorMessage, setErrorMessage] = useState('')
  const [causaTimer, setCausaTimer] = useState(0)
  const [detalleWindowTimer, setDetalleWindowTimer] = useState(0)
  const [finalTimer, setFinalTimer] = useState(0)

  const causaInterval = useRef<NodeJS.Timeout | null>(null)
  const detalleWindowInterval = useRef<NodeJS.Timeout | null>(null)
  const finalInterval = useRef<NodeJS.Timeout | null>(null)

  const enviado = useRef(false)

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target
    setForm(prev => ({ ...prev, [name]: value }))

    if (name === 'full_name') {
      const palabras = value.trim().split(/\s+/)
      if (palabras.length < 2) {
        setErrorMessage('Debes ingresar el nombre completo (mínimo dos palabras).')
      } else {
        setErrorMessage('')
      }
    }

    if (
      name === 'image_url' &&
      value.trim() !== '' &&
      form.full_name.trim().split(/\s+/).length >= 2 &&
      causaTimer === 0
    ) {
      setCausaTimer(40)
    }

    if (name === 'cause_of_death' && value.trim() !== '' && detalleWindowTimer === 0) {
      clearTimer(causaInterval)
      setCausaTimer(0)
      setDetalleWindowTimer(400)
    }

    if (name === 'details' && value.trim() !== '' && finalTimer === 0) {
      setFinalTimer(10)
    }
  }

  useEffect(() => {
    if (causaTimer > 0 && !causaInterval.current) {
      causaInterval.current = setInterval(() => {
        setCausaTimer(prev => {
          if (prev <= 1) {
            clearTimer(causaInterval)
            if (!form.cause_of_death.trim()) {
              enviarAuto({ ...form, cause_of_death: 'Ataque al corazón' })
            }
            return 0
          }
          return prev - 1
        })
      }, 1000)
    }
  }, [causaTimer])

  useEffect(() => {
    if (detalleWindowTimer > 0 && !detalleWindowInterval.current) {
      detalleWindowInterval.current = setInterval(() => {
        setDetalleWindowTimer(prev => {
          if (prev <= 1) {
            clearTimer(detalleWindowInterval)
            return 0
          }
          return prev - 1
        })
      }, 1000)
    }
  }, [detalleWindowTimer])

  useEffect(() => {
    if (finalTimer > 0 && !finalInterval.current) {
      finalInterval.current = setInterval(() => {
        setFinalTimer(prev => {
          if (prev <= 1) {
            clearTimer(finalInterval)
            setForm(currentForm => {
              enviarAuto(currentForm)
              return currentForm
            })
            return 0
          }
          return prev - 1
        })
      }, 1000)
    }
  }, [finalTimer])

  const clearTimer = (ref: React.MutableRefObject<NodeJS.Timeout | null>) => {
    if (ref.current) {
      clearInterval(ref.current)
      ref.current = null
    }
  }

  const enviarAuto = async (victimData: Victim) => {
    if (enviado.current) return
    enviado.current = true

    const data: Victim = {
      ...victimData,
      cause_of_death: victimData.cause_of_death.trim() || 'Ataque al corazón',
      details: victimData.details.trim(),
      is_dead: true
    }

    try {
      await axios.post('http://192.168.66.117:8080/victims', data)
      alert('Víctima registrada exitosamente')
      resetForm()
    } catch (error) {
      console.error(error)
      alert('Error al registrar víctima')
    }
  }

  const resetForm = () => {
    setForm({
      full_name: '',
      image_url: '',
      cause_of_death: '',
      details: '',
      is_dead: true
    })
    setCausaTimer(0)
    setDetalleWindowTimer(0)
    setFinalTimer(0)
    enviado.current = false
    clearTimer(causaInterval)
    clearTimer(detalleWindowInterval)
    clearTimer(finalInterval)
    setErrorMessage('')
  }

  return (
    <div className="registro-container">
      <div className="registro-page">
        <h2 className="form-title">Registrar Persona</h2>
        <form className="form" onSubmit={(e) => e.preventDefault()}>
          <input
            name="full_name"
            placeholder="Nombre completo"
            value={form.full_name}
            onChange={handleChange}
            required
          />
          {errorMessage && <p className="error-message">{errorMessage}</p>}

          <input
            name="image_url"
            placeholder="URL de la imagen"
            value={form.image_url}
            onChange={handleChange}
            required
          />

          {causaTimer > 0 && (
            <p className="timer-text">Tienes {causaTimer}s para escribir la causa de muerte...</p>
          )}
          <input
            name="cause_of_death"
            placeholder="Causa de muerte"
            value={form.cause_of_death}
            onChange={handleChange}
            disabled={!form.image_url.trim() || form.full_name.trim().split(/\s+/).length < 2}
          />

          {detalleWindowTimer > 0 && (
            <p className="timer-text">Tienes {detalleWindowTimer}s para escribir detalles...</p>
          )}
          <textarea
            name="details"
            placeholder="Detalles (opcional)"
            value={form.details}
            onChange={handleChange}
            disabled={!form.cause_of_death.trim()}
          />

          {finalTimer > 0 && (
            <p className="timer-text">Muerte registrada en: {finalTimer}s</p>
          )}
        </form>
      </div>

      <div className="reglas-box">
        <h3>Reglas de la Death Note</h3>
        <ul>
          <li>La persona cuyo nombre sea escrito en este cuaderno morirá.</li>
          <li>Si la causa de la muerte no es especificada, la persona morirá de un ataque al corazón.</li>
          <li>Después de escribir el nombre, el usuario tiene 40 segundos para escribir la causa de muerte.</li>
          <li>Si se especifica la causa de la muerte, el usuario tiene 6 minutos y 40 segundos para escribir los detalles.</li>
          <li>Si el usuario no tiene en mente el rostro de la persona que está escribiendo, la Death Note no funcionará.</li>
          <li>El usuario puede renunciar a la Death Note, perdiendo sus recuerdos sobre ella.</li>
          <li>Un humano que usa la Death Note no podrá ir ni al cielo ni al infierno.</li>
        </ul>
      </div>
    </div>
  )
}

export default Registro