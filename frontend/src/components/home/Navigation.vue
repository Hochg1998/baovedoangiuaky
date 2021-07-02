<template>
  <div class="container">
    <div class="row">
      <div class="col-md-2 text-center">
        <router-link
          :to="{ name: 'ProductByCategory', params: { category: 'dogs' } }"
        >
          <img
            class="smaller-image thick-green-border"
            src="../../assets/images/icons/trade-mark-icon.jpg"
            alt="A cute orange cat lying on its back."
          />
          <p>Thú cưng chó</p>
        </router-link>
      </div>

      <div class="col-md-2 text-center">
        <router-link
          :to="{ name: 'ProductByCategory', params: { category: 'cats' } }"
        >
          <img
            class="smaller-image thick-green-border"
            src="../../assets/images/icons/trade-mark-icon.jpg"
            alt="A cute orange cat lying on its back."
          />
          <p>Thú cưng mèo</p></router-link
        >
      </div>
      <div class="col-md-2 text-center ">
        <router-link
          :to="{ name: 'ProductByCategory', params: { category: 'dogsfood' } }"
        >
          <img
            class="smaller-image thick-green-border"
            src="../../assets/images/icons/trade-mark-icon.jpg"
            alt="A cute orange cat lying on its back."
          />
          <p>Thức ăn cho chó</p></router-link
        >
      </div>
      <div class="col-md-2 text-center">
        <router-link
          :to="{ name: 'ProductByCategory', params: { category: 'catsfood' } }"
        >
          <img
            class="smaller-image thick-green-border"
            src="../../assets/images/icons/trade-mark-icon.jpg"
            alt="A cute orange cat lying on its back."
          />
          <p>Thức ăn cho mèo</p></router-link
        >
      </div>
      <div class="col-md-2 text-center">
        <router-link
          :to="{ name: 'ProductByCategory', params: { category: 'dogsitem' } }"
        >
          <img
            class="smaller-image thick-green-border"
            src="../../assets/images/icons/trade-mark-icon.jpg"
            alt="A cute orange cat lying on its back."
          />
          <p>Đồ dùng cho chó</p></router-link
        >
      </div>
      <div class="col-md-2 text-center">
        <router-link
          :to="{ name: 'ProductByCategory', params: { category: 'catsitem' } }"
        >
          <img
            class="smaller-image thick-green-border"
            src="../../assets/images/icons/trade-mark-icon.jpg"
            alt="A cute orange cat lying on its back."
          />
          <p>Đồ dùng cho mèo</p></router-link
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Navigation",
};
</script>

<style scoped>
.thick-green-border {
  border-color: rgb(132, 202, 219);
  border-width: 10px;
  border-style: solid;
  border-radius: 10px;
  border-radius: 50%;
}

.smaller-image {
  margin-top: 20px;
  margin-bottom: 20px;
  width: 80px;
}
a {
  text-decoration: none;
  color: rgb(91, 92, 92);
}
a:hover {
  color: red;
}
</style>
