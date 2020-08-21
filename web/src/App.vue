<template>
  <v-app>
    <v-app-bar
        app
        color="primary"
        dark
    >
      <div class="text-h3 font-weight-black white--text">LITTLE BROWN BOOK SHOP</div>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn to="/login" color="accent">
          {{ authText }}
        </v-btn>
      </v-toolbar-items>
    </v-app-bar>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",
  computed: {
    authText() {
      return this.$store.state.auth.isLogin ? "Logout" : "Login"
    },
  },
}
</script>
