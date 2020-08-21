<template>
  <v-container>
    <v-row justify="center">
      <v-card width="90%" color="accent white--text">
        <v-card-title>
          <div class="d-flex spacer">
            <div>Net Amount ฿ {{ netAmount }}</div>
            <v-spacer></v-spacer>
          </div>
          <checkout-dialog></checkout-dialog>
        </v-card-title>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
import CheckoutDialog from "@/components/molecules/CheckoutDialog"
export default {
  name: "NetAmountDialog",
  components: { CheckoutDialog },
  props: {
    netAmount: {
      required: true,
    },
  },
}
</script>

<style scoped>

</style>