import { bookService } from "@/services"

export const book = {
  namespaced: true,
  state: {
    books: [],
  },
  actions: {
    async getBooks({ commit }) {
      try {
        const books = await bookService.getAll()
        commit("setBooks", books)
      } catch (err) {
        alert(err)
      }
    },
  },
  mutations: {
    setBooks(state, books) {
      state.books = books
    },
  },
}