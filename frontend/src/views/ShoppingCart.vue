<template>
  <div class="container">
    <h5 class="text-start">Shopping Cart</h5>
    <HeaderSC />
    <br />
    <ProductsList />
    <br />
    <Checkout />
  </div>
</template>

<script>
import ProductsList from "@/components/shoppingcart/ProductsList";
import HeaderSC from "@/components/shoppingcart/HeaderSC";

import Checkout from "@/components/shoppingcart/Checkout";

export default {
  name: "App",
  components: {
    ProductsList,
    HeaderSC,
    Checkout,
  },
  data() {
    return {};
  },
  methods: {
    rollbackProduct() {
      console.log(this.deletedProducts);
      this.deletedProducts.forEach((product) => this.products.push(product));
      this.deletedProducts = [];
    },
    deleteProduct(id) {
      if (confirm("Are you sure about dat?")) {
        this.products = this.products.filter((product) => {
          if (product.id === id) {
            console.log(product);
            this.deletedProducts.push(product);
          }

          return product.id != id;
        });
      }
    },
  },
  computed: {
    totalProductInCart() {
      console.log(this.$store.state.cart);
      return this.$store.state.cart;
    },
    countTotal() {
      let totalProducts = 0;
      this.products.forEach((product) => {
        totalProducts += parseInt(product.quantity);
      });
      return parseInt(totalProducts);
    },
    preventEmpty() {
      this.products.forEach((product) => {
        if (
          product.quantity === undefined ||
          product.quantity === "" ||
          parseInt(product.quantity) > 99
        ) {
          product.quantity = 0;
        }
      });
      return 1;
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
h5 {
  color: green;
}
</style>
