<template>
  <v-row justify="center" style="margin: 0">
    <v-dialog v-model="dialog" scrollable max-width="800px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
            color="primary"
            dark
            v-bind="attrs"
            v-on="on"
            width="100%"
            :disabled="parseInt(netAmount) === 0"
        >
          Checkout
        </v-btn>
      </template>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-card>
          <v-card-title>
            <v-container class="d-flex justify-center text-h4 font-weight-bold ">
              <v-icon color="black" x-large>mdi-cash-register</v-icon>
              Summary
            </v-container>
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <order-list-table :orders="cart.orders" :read-only="true"></order-list-table>
          </v-card-text>
          <v-container class="text-center text-h5 font-weight-bold">
            <div class="blue-grey--text">Discount: ฿{{ totalDiscount }}</div>
            <div class="light-green--text">Net Amount: ฿{{ netAmount }}</div>
          </v-container>
          <div class="d-flex justify-center">
            <div style="margin-bottom: 8px">
              <v-text-field class="text-h5 font-weight-bold text-center" v-model="cash" outlined rounded solo
                            placeholder="Cash Amount" :hide-details="false" required :rules="cashRules"></v-text-field>
            </div>
          </div>
          <v-card-actions class="justify-center">
            <v-btn shaped color="primary" large @click="checkout(cash)" width="230px" style="margin-bottom: 8px"
                   :disabled="!valid">Checkout
            </v-btn>
            <change-dialog :dialog="changeDialog" :change="change" @resetForm="resetForm"
            ></change-dialog>
          </v-card-actions>
        </v-card>
      </v-form>
    </v-dialog>
  </v-row>
</template>

<script>
import OrderListTable from "@/components/molecules/OrderListTable"
import ChangeDialog from "@/components/molecules/ChangeDialog"

export default {
  name: "CheckoutDialog",
  components: { ChangeDialog, OrderListTable },

  data() {
    return {
      dialog: false,
      cash: "",
      change: "",
      changeDialog: false,
      valid: false,
      cashRules: [
        v => !!v || "cash amount is required",
        v => this.isInt(v) || "cash amount must contain only digit numbers",
        v => parseInt(v) >= parseFloat(this.netAmount) || "cash amount must be greater than net amount",
      ],
    }
  },
  methods: {
    isInt(value) {
      let er = /^-?[0-9]+$/
      return er.test(value)
    },
    async checkout(cash) {
      try {
        this.change = await this.$store.dispatch("cart/checkout", { cash })
        this.dialog = false
        this.changeDialog = true
      } catch (err) {
        console.log(err)
      }
    },
    resetForm() {
      this.dialog = false
      this.changeDialog = false
      this.cash = ""
      this.change = ""
      this.valid = false
    },
  },
  computed: {
    cart() {
      return this.$store.state.cart.cart
    },
    totalDiscount() {
      return parseFloat(this.$store.state.cart.totalDiscount).toFixed(2)
    },
    netAmount() {
      return parseFloat(this.$store.state.cart.netAmount).toFixed(2)
    },
    propDialog: {
      get() {
        return this.dialog
      },
      set(v) {
        if (v === false) {
          this.$emit("close", v)
        }
      },
    },
  }, watch: {
    changeDialog(d) {
      if (d === false) {
        this.resetForm()
      }
    },
  },
}
</script>

<style scoped>

</style>