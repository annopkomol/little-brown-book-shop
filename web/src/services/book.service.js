import { authHeader } from "@/utils"
import { authService as auth } from "./auth.service"

const API_URL = process.env.VUE_APP_API_URL

export const bookService = {
  getAll,
}

async function getAll() {
  const requestOptions = {
    method: "GET",
    headers: authHeader(),
  }

  const res = await fetch(`${API_URL}/books`, requestOptions).then(auth.handleResponse)
  return await res.json()
}

