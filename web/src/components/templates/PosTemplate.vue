<template>
  <v-row>
    <v-col md="8">
      <item-list></item-list>
    </v-col>
    <v-col md="4">
      <order-list></order-list>
    </v-col>
  </v-row>
</template>

<script>
import ItemList from "@/components/organisms/ItemList.vue"
import OrderList from "@/components/organisms/OrderList"

export default {
  name: "PosTemplate",
  components: { ItemList, OrderList },
}
</script>

<style scoped>

</style>