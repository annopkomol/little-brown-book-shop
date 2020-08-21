<template>

    <v-simple-table
        :fixed-header="true"
        :height="500"
        style="table-layout: fixed;"
    >
      <template v-slot:default>
        <thead>
        <tr>
          <th style="width:50%" class="text-left">Tittle</th>
          <th style="width:30%" class="text-center">Qty</th>
          <th style="width:20%" class="text-right">Price</th>
        </tr>
        </thead>
        <tbody>
        <tr style="padding: 0 0" v-for="o in orders" :key="o.id">
          <td style="padding: 0 4px; width: 50%">{{ o.book.title }}</td>
          <td class="text-center" style="padding: 0 4px; width: 30%">
            <v-btn v-if="!readOnly" icon color="accent" @click="removeBook(o.book.id)">
              <v-icon dense>mdi-minus</v-icon>
            </v-btn>
            {{ o.qty }}
            <v-btn v-if="!readOnly" icon color="primary" @click="addBook(o.book.id)">
              <v-icon dense>mdi-plus</v-icon>
            </v-btn>
          </td>
          <td class="text-right" style="padding: 0 4px; width: 20%">฿ {{ o.book.price * o.qty }}</td>
        </tr>
        </tbody>
      </template>
    </v-simple-table>
</template>

<script>
export default {
  name: "OrderListTable",
  props: {
    orders: {
      required: true,
    },
    readOnly: {
      default: false,
    }
  },
  methods: {
    addBook(bookID) {
      this.$store.dispatch("cart/addBook", { bookID })
    },
    removeBook(bookID) {
      this.$store.dispatch("cart/removeBook", { bookID })
    },
  },
}
</script>

<style scoped>

</style>