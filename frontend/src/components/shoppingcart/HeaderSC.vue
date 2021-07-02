<template>
  <div class="row buffer-top">
    <p class="text-start col-md-6">
      <router-link to="/">Home</router-link> > Shopping Cart
    </p>
    <p class="text-end col-md-6">{{ quantityTotal }} items in the bag</p>
  </div>
</template>

<script>
export default {
  name: "Header",
  computed: {
    quantityTotal() {
      return this.$store.state.cart.length;
    },
  },
};
</script>

<style scoped>
h5 {
  color: red;
}
.buffer-top {
  margin-top: 110px;
}
@media only screen and (max-width: 768px) {
  .buffer-top {
    margin-top: 210px;
  }
}
</style>
