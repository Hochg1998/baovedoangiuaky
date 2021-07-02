<template>
  <Promotions />
  <Navigation />
  <SaleProducts />
</template>

<script>
import Promotions from "@/components/home/Promotions";
import Navigation from "@/components/home/Navigation";
import SaleProducts from "@/components/home/SaleProducts";

export default {
  name: "Home",
  components: {
    Promotions,
    Navigation,
    SaleProducts,
  },
};
</script>

<style></style>
