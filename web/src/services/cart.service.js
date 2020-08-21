import { authHeader } from "@/utils"
import { authService as auth } from "./auth.service"

export const cartService = {
  getCart,
  addBookToCart,
  removeBookFromCart,
  cartCheckout,
}

const API_URL = process.env.VUE_APP_API_URL

async function getCart() {
  const requestOptions = {
    method: "GET",
    headers: authHeader(),
  }
  const res = await fetch(`${API_URL}/cart`, requestOptions).then(auth.handleResponse)
  return await res.json()
}

async function addBookToCart(bookID) {
  const requestOptions = {
    method: "POST",
    headers: authHeader(),
  }
  const res = await fetch(`${API_URL}/cart/${bookID}`, requestOptions).then(auth.handleResponse)
  return await res.json()
}

async function removeBookFromCart(bookID) {
  const requestOptions = {
    method: "DELETE",
    headers: authHeader(),
  }
  const res = await fetch(`${API_URL}/cart/${bookID}`, requestOptions).then(auth.handleResponse)
  return await res.json()
}

async function cartCheckout(cash) {
  const formData = new URLSearchParams()
  formData.append("cash", cash)
  let headers = authHeader()
  headers["content-type"] = "application/x-www-form-urlencoded"
  const requestOptions = {
    method: "POST",
    headers: headers,
    body: formData,
  }
  const res = await fetch(`${API_URL}/cart/checkout`, requestOptions).then(auth.handleResponse)
  return await res.json()
}
