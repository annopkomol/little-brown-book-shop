<template>
  <v-container>
    <v-row justify="center">
      <v-card width="90%" color="primary white--text">
        <v-card-title>
          <table style="width:100%;">
            <tbody>
            <tr>
              <td class="text-left">
                Discount
              </td>
              <td class="text-right">
                ฿ {{ totalDiscount }}
              </td>
            </tr>
            </tbody>
          </table>
        </v-card-title>
        <div v-if="discounts">
          <v-card-subtitle v-for="(d,i) in discounts" :key="i">
            {{ d.message }}
          </v-card-subtitle>
        </div>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "DiscountDialog",
  props: {
    totalDiscount: {
      required: true,
    },
    discounts: {
      required: true,
    },
  },
}
</script>

<style scoped>

</style>