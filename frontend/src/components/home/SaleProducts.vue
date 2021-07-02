<template>
  <div class="container">
    <h2 class="text-start buffer-top border-bottom">On-sale Products</h2>
    <div class="dropdown col-md-2">
      <select
        @change="printText"
        class="form-select"
        aria-label="Default select example"
      >
        <option value="" selected>Sắp xếp</option>
        <option value="sort-desc">Price: high to low</option>
        <option value="sort-asc">Price: low to high</option>
      </select>
    </div>
    <br />
    <br />
    <div class="row row-cols-md-3 g-5">
      <div
        class="col sanpham product-item"
        v-for="product in products"
        :key="product.ID"
      >
        <router-link
          :to="{ name: 'ProductDetails', params: { id: product.ID } }"
        >
          <div class="card bg-light">
            <img :src="`http://localhost:3000/${product.Image}`" />
            <h3>{{ product.Name }}</h3>
            <h4
              v-bind:style="[
                product.IsSale == true ? { color: 'red' } : { color: 'black' },
              ]"
            >
              {{ formatPrice(product.Price) }}
            </h4>
            <button
              @click.prevent="addOneProductToCart(product)"
              class="btn themecolor"
            >
              Thêm vào giỏ
            </button>
          </div></router-link
        >
      </div>
    </div>
    <br />
    <br />
    <br />

    <div class="row">
      <div class="col-md-5"></div>

      <div class="col-md-3">
        <nav aria-label="Page navigation example">
          <ul class="pagination">
            <li class="page-item">
              <a class="page-link" href="#">Previous</a>
            </li>
            <li
              v-for="(page, index) in totalPage"
              :key="index"
              class="page-item"
            >
              <a class="page-link" @click="goToOtherPage(index)" href="#">{{
                index + 1
              }}</a>
            </li>

            <li class="page-item"><a class="page-link" href="#">Next</a></li>
          </ul>
        </nav>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "SaleProducts",
  data() {
    return {
      totalPage: 0,
      products: [],
    };
  },
  methods: {
    printText(e) {
      // document.getElementsByClassName("printText").innerText =
      if (e.target.value === "sort-asc") {
        this.products.sort((prev, next) => prev.Price - next.Price);
      } else if (e.target.value === "sort-desc") {
        this.products.sort((prev, next) => next.Price - prev.Price);
      } else if (e.target.value === "1") {
        this.products.sort((prev, next) => prev.ID - next.ID);
      }
    },
    async goToOtherPage(index) {
      const response = await fetch(
        `http://localhost:3000/api/product/issale/true?limit=6&offset=${index *
          6}`
      );
      let results = await response.json();
      this.products = results.listProduct;
    },
    formatPrice(price) {
      return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    addOneProductToCart(product) {
      product["quantity"] = 1;
      this.$store.commit("addProductToCart", product);
    },
  },
  async created() {
    const response = await fetch(
      "http://localhost:3000/api/product/issale/true?limit=6&offset=0"
    );
    let results = await response.json();
    this.products = results.listProduct;

    this.totalPage = results.totalPages;
    console.log(this.totalPage);
  },
};
</script>

<style scoped>
a {
  text-decoration: none;
  color: rgb(91, 92, 92);
}
a:hover {
  color: red;
}
.buffer-top {
  margin-top: 50px;
  margin-bottom: 50px;
}
.product-item:hover {
  transform: scale(1.1);
}
.product-item {
  transition: transform 0.6s ease;
}
</style>
