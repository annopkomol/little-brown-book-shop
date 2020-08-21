<template>
  <div>
    <v-row style="margin: 4px">
      <v-col md="8">
        <div class="d-flex">
          <div>
            <v-icon x-large>mdi-format-list-bulleted-type</v-icon>
          </div>
          <div class="text-h4 font-weight-bold text--secondary">Item List</div>
          <v-spacer></v-spacer>
          <v-text-field
              rounded
              solo
              placeholder="search for a book..."
              v-model="searchText"
          >
          </v-text-field>
        </div>
      </v-col>
    </v-row>
    <v-card>
      <v-row justify="start" align="start" style="overflow: auto; height: 720px">
        <v-col md="2" v-for="b in filteredBooks" :key="b.id">
          <item-card
              :price="b.price"
              :cover="b.cover"
              :title="b.title"
              :id="b.id"
          ></item-card>
        </v-col>
      </v-row>
    </v-card>

  </div>
</template>

<script>
import ItemCard from "@/components/molecules/ItemCard"

export default {
  name: "ItemList",
  components: { ItemCard },
  data: () => ({
    searchText: "",
  }),
  async created() {
    await this.$store.dispatch("book/getBooks")
  },
  computed: {
    books() {
      return this.$store.state.book.books
    },
    filteredBooks() {
      return this.books.filter(book => {
        return book.title.toLowerCase().includes(this.searchText.toLowerCase())
      })
    },
  },
}
</script>