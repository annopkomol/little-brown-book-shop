import { authService } from "@/services"
import router from "@/router"

const token = localStorage.getItem("token")
const defaultLoginState = !!token


export const auth = {
  namespaced: true,
  state: { isLogin: defaultLoginState },
  actions: {
    login({ commit }, { username, password }) {
      authService.login(username, password)
        .then(
          () => {
            commit("loginSuccess")
            router.push("/")
          },
          error => {
            commit("loginFailure")
            alert(error)
          },
        )
    },
    logout({ commit }) {
      authService.logout()
      commit("logout")
    },
  },
  mutations: {
    loginSuccess(state) {
      state.isLogin = true
    },
    loginFailure(state) {
      state.isLogin = false
    },
    logout(state) {
      state.isLogin = false
    },
  },
}