import { cartService } from "@/services"

export const cart = {
  namespaced: true,
  state: {
    cart: {},
    discounts: [],
    totalDiscount: 0,
    netAmount: 0,
  },
  actions: {
    async getCart({ commit }) {
      const res = await cartService.getCart()
      commit("setCart", res)

    },
    async addBook({ dispatch }, { bookID }) {
      try {
        await cartService.addBookToCart(bookID)
        dispatch("getCart")
      } catch (err) {
        alert("couldn't add book to cart")
      }
    },
    async removeBook({ dispatch }, { bookID }) {
      try {
        await cartService.removeBookFromCart(bookID)
        dispatch("getCart")
      } catch (err) {
        alert("couldn't remove book from cart")
      }
    },
    async checkout({ dispatch }, { cash }) {
      try {
        let res = await cartService.cartCheckout(cash)
        dispatch("getCart")
        return res.change
      } catch (err) {
        alert("couldn't checkout cart")
        throw err
      }
    },
  },
  mutations: {
    setCart(state, data) {
      state.cart = data.cart
      state.discounts = data.discounts
      state.totalDiscount = data.total_discount
      state.netAmount = data.net_amount
    },
  },
}