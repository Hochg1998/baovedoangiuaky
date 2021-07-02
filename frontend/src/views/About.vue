<template>
  <div class="about buffer-top">
    <h1>This is an about page</h1>
  </div>
</template>

<script>
export default {
  name: "About",
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
