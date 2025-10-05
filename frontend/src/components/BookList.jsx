import { useEffect, useState } from 'react'
import { get, post, del } from '../api'

export default function BookList({ refreshKey = 0 }) {
  const [books, setBooks] = useState([])
  const [q, setQ] = useState('')
  const [userId, setUserId] = useState('')

  const load = async () => {
    try {
      const res = await get('/books')
      setBooks(Array.isArray(res) ? res : [])
    } catch (e) {
      console.error(e)
      setBooks([])
    }
  }
  useEffect(()=>{ load() }, [refreshKey])

  const search = async () => {
    try {
      const res = await get(`/books/search?q=${encodeURIComponent(q)}`)
      setBooks(Array.isArray(res) ? res : [])
    } catch (e) {
      console.error(e)
      setBooks([])
    }
  }
  const borrow = async (bookId) => {
    try {
      await post('/loans/borrow', { userId, bookId })
      load()
    } catch (e) {
      alert(String(e))
    }
  }
  const ret = async (bookId) => {
    try {
      await post('/loans/return', { userId, bookId })
      load()
    } catch (e) {
      alert(String(e))
    }
  }

  const removeBook = async (bookId) => {
    if (!confirm('¿Eliminar este libro?')) return
    try {
      await del(`/books?id=${encodeURIComponent(bookId)}`)
      load()
    } catch (e) {
      alert(String(e))
    }
  }

  return (
    <div>
      <h3>Libros</h3>
      <div className="row" style={{marginBottom:8}}>
        <input className="input" placeholder="Buscar..." value={q} onChange={e=>setQ(e.target.value)} />
        <button className="btn" onClick={search}>Buscar</button>
        <input className="input" placeholder="Usuario para préstamo" value={userId} onChange={e=>setUserId(e.target.value)} />
      </div>
      <table className="table">
        <thead><tr><th>ID</th><th>Título</th><th>Autor</th><th>Disponible</th><th>Acciones</th></tr></thead>
        <tbody>
          {(Array.isArray(books) ? books : []).length === 0 ? (
            <tr><td colSpan="5" style={{color:'#9ca3af'}}>No hay libros cargados todavía.</td></tr>
          ) : (
            (Array.isArray(books) ? books : []).map(b => (
              <tr key={b.id}>
                <td>{b.id}</td>
                <td>{b.title}</td>
                <td>{b.author}</td>
                <td>{b.available? 'Sí':'No'}</td>
                <td className="row" style={{gap:6}}>
                  <button className="btn success" onClick={()=>borrow(b.id)} disabled={!b.available || !userId}>Prestar</button>
                  <button className="btn danger" onClick={()=>ret(b.id)} disabled={b.available || !userId}>Devolver</button>
                  <button className="btn secondary" onClick={()=>removeBook(b.id)}>Eliminar</button>
                </td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  )
}
