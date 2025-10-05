import { useState } from 'react'
import { post } from '../api'

export default function BookForm({ onCreated }) {
  const [form, setForm] = useState({ id:'', title:'', author:'', isbn:'' })
  const submit = async (e) => {
    e.preventDefault()
    await post('/books', form)
    setForm({ id:'', title:'', author:'', isbn:'' })
    onCreated && onCreated()
  }
  return (
    <form onSubmit={submit} className="grid-1" style={{gap:10}}>
      <h3>Registrar Libro</h3>
      <input className="input" placeholder="ID" value={form.id} onChange={e=>setForm({...form,id:e.target.value})} required />
      <input className="input" placeholder="Título" value={form.title} onChange={e=>setForm({...form,title:e.target.value})} required />
      <input className="input" placeholder="Autor" value={form.author} onChange={e=>setForm({...form,author:e.target.value})} required />
      <input className="input" placeholder="ISBN" value={form.isbn} onChange={e=>setForm({...form,isbn:e.target.value})} />
      <div className="row" style={{justifyContent:'flex-end'}}>
        <button className="btn" type="submit">Agregar libro</button>
      </div>
    </form>
  )
}
