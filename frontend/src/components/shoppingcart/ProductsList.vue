<template>
  <div v-if="products.length > 0">
    <div
      v-for="product in products"
      :key="product.ID"
      class="d-flex border-top border-bottom"
    >
      <div class="align-self-start">
        <div class="flex-shrink-0">
          <img :src="`http://localhost:3000/${product.Image}`" />
        </div>
      </div>
      <div class="align-self-baseline">
        <router-link
          :to="{ name: 'ProductDetails', params: { id: product.ID } }"
        >
          <h5 class="text-start">{{ product.Name }}</h5>
        </router-link>
        <p class="text-start">{{ product.description }}</p>
        <h4 class="text-start">{{ formatPrice(product.Price) }}</h4>
      </div>
      <div class="ms-auto align-self-center">
        <input
          type="button"
          @click="decreaseAmount(product.ID)"
          value="-"
          class="minus"
        /><input
          type="number"
          step="1"
          min="1"
          max="100"
          readonly
          name="quantity"
          v-model="product.quantity"
          title="Qty"
          class="input-text qty text"
          size="4"
          pattern=""
          inputmode=""
        /><input
          type="button"
          @click="increaseAmount(product.ID)"
          value="+"
          class="plus"
        />
      </div>
      <div class="ms-auto align-self-center">
        <i @click="deleteProductInCart(product.ID)" class="fa fa-close"></i>
      </div>
    </div>
  </div>

  <div class="buffer-top" v-else>
    <h1>Nothing to see here</h1>
  </div>
  <button
    v-show="deletedProducts.length > 0"
    @click="rollbackProduct()"
    class="btn btn-success"
  >
    Rollback
  </button>
</template>

<script>
export default {
  name: "ProductsList",

  computed: {
    products() {
      return this.$store.state.cart;
    },
    deletedProducts() {
      return this.$store.state.deletedProducts;
    },
  },
  methods: {
    formatPrice(price) {
      return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    decreaseAmount(id) {
      this.$store.commit("decreaseAmount", id);
    },
    increaseAmount(id) {
      this.$store.commit("increaseAmount", id);
    },
    deleteProductInCart(id) {
      this.$store.commit("deleteProduct", id);
    },
    rollbackProduct() {
      this.$store.commit("rollbackProduct");
    },
  },
};
</script>

<style scoped>
img {
  display: block;
  max-width: 150px;
  max-height: 200px;
  width: auto;
  height: auto;
}
h5 {
  color: purple;
}
a {
  text-decoration: none;
  color: rgb(91, 92, 92);
}
a:hover {
  color: red;
}
</style>
