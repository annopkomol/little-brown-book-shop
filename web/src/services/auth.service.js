export const authService = {
  login,
  handleResponse,
  logout,
}

const API_URL = process.env.VUE_APP_API_URL

async function login(username, password) {
  const formData = new URLSearchParams()
  formData.append("username", username)
  formData.append("password", password)
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: formData,
  }

  const res = await fetch(`${API_URL}/auth/login`, requestOptions)
  if (res.ok) {
    let obj = await res.json()
    const token = obj.access_token
    localStorage.setItem("token", token)
  } else {
    throw Error(await res.text())
  }
  return res
}

function logout() {
  localStorage.removeItem("token")
}

async function handleResponse(resp) {
  if (!resp.ok) {
    if (resp.status === 401) {
      logout()
      location.reload()
    }
    const err = await resp.text()
    throw Error(err)
  }
  return resp
}