<template>
  <div class="header fixed-top">
    <div class="container">
      <div class="row">
        <div class="col-md-6 text-start">
          Kết nối &emsp;
          <a href="" class="me-4 text-reset">
            <i class="fab fa-facebook-f"></i>
          </a>
          <a href="" class="me-4 text-reset">
            <i class="fab fa-instagram"></i>
          </a>
        </div>
        <div class="col-md-6 text-end">Thông báo - Đăng nhập / đăng ký</div>
      </div>
      <div class="row top-buffer">
        <div class="col-md-3">
          <router-link to="/"
            ><h1 class="text-start icontohome">
              <span class="textRED">Stupid</span> Pet
            </h1></router-link
          >
        </div>
        <div class="col-md-7 buffer-top">
          <form class="d-flex flex-row" v-on:submit.prevent="">
            <input
              class="form-control"
              placeholder="Search"
              aria-label="Search"
              v-model="inputSearch"
            />

            <router-link
              :to="{
                name: 'SearchResults',
                params: { input_name: inputSearch },
              }"
              ><button class="btn btn-outline-light" type="submit">
                Search
              </button></router-link
            >
          </form>
        </div>
        <div class="col-md-2 text-center buffer-top">
          <router-link to="/shoppingcart"
            ><i class="fas fa-shopping-cart"
              ><sup class="textRED" v-if="totalProductInCart > 0">{{
                totalProductInCart
              }}</sup></i
            ></router-link
          >
        </div>
      </div>
    </div>
  </div>
  <div></div>
</template>

<script>
export default {
  name: "Header",
  components: {},
  data() {
    return {
      inputSearch: " ",
    };
  },
  computed: {
    totalProductInCart() {
      return this.$store.state.cart.length;
    },
  },
};
</script>

<style scoped>
.top-buffer {
  padding-top: 20px;
  padding-bottom: 20px;
  margin-bottom: 20px;
}

.header {
  background-color: rgb(132, 202, 219);
}
a {
  text-decoration: none;
  color: rgb(91, 92, 92);
}
a:hover {
  color: blue;
}
.fas {
  font-size: 250%;
}
.buffer-top {
  padding-top: 8px;
}
i:hover {
  color: red;
  cursor: pointer;
}
.icontohome {
  border-radius: 10px;
  background: white;
  padding-left: 7px;
  width: 200px;
  height: 50px;
}
.textRED {
  color: red;
}

/* .numberofproductincart {
    position: absolute;
    top: -13  px;
    font-size: 20px;
    <span class="numberofproductincart">0</span>
  } */
</style>
