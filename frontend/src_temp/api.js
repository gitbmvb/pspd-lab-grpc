// src/api.js
export const loginUser = async (email, password) => {
    const res = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    })
  
    if (!res.ok) throw new Error('Login failed')
    return res.json()
  }
  
  export const registerUser = async (name, email, password) => {
    const res = await fetch('http://localhost:8080/api/users/create', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, name, password }),
    })
    console.log('Request body:', JSON.stringify({ email, name, password })) // Log the request body for debugging
    const data = await res.json()

    console.log('Response data:', data) // Log the response data for debugging
  
    if (!res.ok) {
      // You can log or show this detailed error message
      console.error('Registration error:', data)
      throw new Error(data.message || 'Registration failed')
    }
  
    return data
  }  
  