<template>
  <v-card
      class="mx-auto"
      max-width="180"
      @click="addBook(id)"
  >
    <v-img
        :src="cover"
        width="100%"
        contain
    ></v-img>
    <v-card-title class="text-subtitle-2">
      {{ title }}
    </v-card-title>
    <v-card-text>
      ฿ {{ price }}
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    id: {
      type: Number,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    cover: {
      type: String,
      required: true,
    },
    price: {
      type: String,
      required: true,
    },
  },
  data: () => ({
    show: false,
  }),
  methods: {
    addBook(bookID) {
      this.$store.dispatch("cart/addBook", { bookID })
    },
  },
}
</script>

<style scoped>

</style>