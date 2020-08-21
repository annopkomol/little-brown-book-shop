import Vue from "vue"
import VueRouter from "vue-router"
import Home from "../views/Home.vue"
import Login from "@/views/Login.vue"

Vue.use(VueRouter)

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  { path: "*", redirect: "/" },
]

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
})

router.beforeEach((to, from, next) => {

  const publicPages = ["/login"]
  const authRequired = !publicPages.includes(to.path)
  const loggedIn = localStorage.getItem("token")

  if (authRequired && !loggedIn) {
    return next("/login")
  }

  next()
})

export default router
