import { useState } from 'react'
import { post } from '../api'

export default function UserForm({ onCreated }) {
  const [form, setForm] = useState({ id:'', name:'' })
  const submit = async (e) => {
    e.preventDefault()
    await post('/users', form)
    setForm({ id:'', name:'' })
    onCreated && onCreated()
  }
  return (
    <form onSubmit={submit} className="grid-1" style={{gap:10}}>
      <h3>Registrar Usuario</h3>
      <input className="input" placeholder="ID" value={form.id} onChange={e=>setForm({...form,id:e.target.value})} required />
      <input className="input" placeholder="Nombre" value={form.name} onChange={e=>setForm({...form,name:e.target.value})} required />
      <div className="row" style={{justifyContent:'flex-end'}}>
        <button className="btn" type="submit">Agregar usuario</button>
      </div>
    </form>
  )
}
