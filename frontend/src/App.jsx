import React, { useState } from 'react'
import BookForm from './components/BookForm'
import UserForm from './components/UserForm'
import BookList from './components/BookList'
import UserList from './components/UserList'
import './styles.css'

export default function App() {
  const [booksVersion, setBooksVersion] = useState(0)
  const [usersVersion, setUsersVersion] = useState(0)

  const onBookCreated = () => setBooksVersion(v => v + 1)
  const onUserCreated = () => setUsersVersion(v => v + 1)

  return (
    <div className="container">
      <div className="header">
        <div>
          <div className="brand">📚 Biblioteca</div>
          <div className="subtitle">Demostración con estructuras de datos en Go + React</div>
        </div>
        <div className="row">
          <span className="kbd">GET</span>
          <span className="kbd">POST</span>
          <span className="kbd">/api/*</span>
        </div>
      </div>

      <div className="grid-2">
        <div className="card"><BookForm onCreated={onBookCreated} /></div>
        <div className="card"><UserForm onCreated={onUserCreated} /></div>
      </div>

      <div className="grid-1" style={{marginTop:16}}>
        <div className="card"><BookList refreshKey={booksVersion} /></div>
        <div className="card"><UserList refreshKey={usersVersion} /></div>
      </div>
    </div>
  )
}
