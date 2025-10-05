import { useEffect, useState } from 'react'
import { get, del } from '../api'

export default function UserList({ refreshKey = 0 }) {
  const [users, setUsers] = useState([])
  useEffect(()=>{ (async()=>{
    try {
      const res = await get('/users')
      setUsers(Array.isArray(res) ? res : [])
    } catch (e) {
      console.error(e)
      setUsers([])
    }
  })() }, [refreshKey])
  return (
    <div>
      <h3>Usuarios</h3>
      <ul>
        {(Array.isArray(users) ? users : []).map(u => (
          <li key={u.id} style={{display:'flex', alignItems:'center', justifyContent:'space-between', gap:8}}>
            <span>{u.id} - {u.name}</span>
            <button className="btn secondary" onClick={async()=>{
              if (!confirm('¿Eliminar este usuario?')) return
              try { await del(`/users?id=${encodeURIComponent(u.id)}`); const res = await get('/users'); setUsers(Array.isArray(res)?res:[]) } catch(e){ alert(String(e)) }
            }}>Eliminar</button>
          </li>
        ))}
      </ul>
    </div>
  )
}
