<template>
  <h2 class="text-center buffer-top">404</h2>
  <h3 class="text-center">Page Not Found</h3>
</template>

<script>
export default {
  name: "404Page",
  components: {},
};
</script>

<style scoped>
.buffer-top {
  margin-top: 180px;
}
@media only screen and (max-width: 768px) {
  .buffer-top {
    margin-top: 300px;
  }
}
</style>
