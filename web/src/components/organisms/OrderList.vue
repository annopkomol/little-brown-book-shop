<template>
  <div>
    <div class="d-flex">
      <div>
        <v-icon x-large>mdi-cart</v-icon>
      </div>
      <div class="text-h4 font-weight-bold text--secondary">Cart Ref: {{ cart.id }}</div>
      <v-spacer></v-spacer>
      <div>
        <v-icon x-large>mdi-cash-register</v-icon>
      </div>
      <div class="text-h4 font-weight-bold text--secondary">POS ID: {{ cart.pos_terminal_id }}</div>
    </div>
    <v-card style="margin: 20px 0">
      <order-list-table :orders="cart.orders"></order-list-table>
    </v-card>
    <discount-dialog :discounts="discounts" :total-discount="totalDiscount"></discount-dialog>
    <net-amount-dialog :netAmount="netAmount"></net-amount-dialog>
  </div>
</template>

<script>
import NetAmountDialog from "@/components/molecules/NetAmountDialog"
import DiscountDialog from "@/components/molecules/DiscountDialog"
import OrderListTable from "@/components/molecules/OrderListTable"

export default {
  name: "OrderList",
  components: { OrderListTable, DiscountDialog, NetAmountDialog },
  async created() {
    await this.$store.dispatch("cart/getCart")
  },
  computed: {
    cart() {
      return this.$store.state.cart.cart
    },
    discounts() {
      return this.$store.state.cart.discounts
    },
    totalDiscount() {
      return parseFloat(this.$store.state.cart.totalDiscount).toFixed(2)
    },
    netAmount() {
      return parseFloat(this.$store.state.cart.netAmount).toFixed(2)
    },
  },
  methods: {},
}
</script>