<template>
  <v-container>
    <v-row justify="center" align="center" style="height: 70vh">
      <v-col cols="12" sm="6" lg="4">
        <v-card>
          <v-card-text>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field
                  label="Username"
                  v-model="username"
                  :rules="nameRules"
                  required
              ></v-text-field>
              <v-text-field
                  label="Password"
                  type="password"
                  v-model="password"
                  :rules="passwordRules"
                  required
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-divider class="mt-4"></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="login" :disabled="!valid">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    username: "",
    password: "",
    nameRules: [
      v => !!v || "Username is required",
    ],
    passwordRules: [
      v => !!v || "Password is required",
    ],
  }),
  created() {
    this.$store.dispatch("auth/logout")
  },
  methods: {
    login() {
      this.$refs.form.validate()
      const { username, password } = this
      const { dispatch } = this.$store
      if (username && password) {
        dispatch("auth/login", { username, password })
      }
    },
  },
}
</script>