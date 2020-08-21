<template>
  <div class="text-center">
    <v-dialog
        v-model="propDialog"
        width="500"
    >
      <v-card>
        <v-card-title class="headline grey lighten-2">
          Checkout Success
        </v-card-title>
        <v-card-text class="text-center font-weight-bold text-h5">
          <div>Change: ฿{{ change }}</div>
          <div>Thank you. See you again!</div>
        </v-card-text>
        <v-divider></v-divider>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  name: "ChangeDialog",
  props: {
    dialog: {
      default: false,
    },
    ifClosed: {},
    change: {
      required: true,
    },
  }, watch: {}, computed: {
    propDialog: {
      get() {
        return this.dialog
      },
      set() {
          this.$emit("resetForm")
      },
    },
  },
}
</script>

<style scoped>

</style>