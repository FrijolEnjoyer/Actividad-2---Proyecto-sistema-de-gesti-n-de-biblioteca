const base = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api'

export async function get(path) {
  const res = await fetch(`${base}${path}`)
  if (!res.ok) throw new Error(await res.text())
  return res.json()
}

export async function post(path, body) {
  const res = await fetch(`${base}${path}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
  if (!res.ok) throw new Error(await res.text())
  return res.json()
}

export async function del(path) {
  const res = await fetch(`${base}${path}`, { method: 'DELETE' })
  if (!res.ok) throw new Error(await res.text())
  return res.json()
}
