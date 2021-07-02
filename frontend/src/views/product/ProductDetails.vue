<template>
  <div class="container">
    <div class="row buffer-top ">
      <div class="row">
        <h2 class="text-start border-bottom">Thông tin chi tiết sản phẩm</h2>
      </div>
      <br />
      <div class="row ">
        <div class="col-md-4">
          <img :src="`http://localhost:3000/${product.Image}`" />
        </div>
        <div class="col-md-8">
          <div class="row">
            <h3>{{ product.Name }}</h3>
          </div>
          <div class="row">
            <h4 class="redtext">{{ formatPrice(product.Price) }}</h4>
          </div>
          <br />
          <br />
          <br />
          <br />
          <br />
          <br />

          <div class="row">
            <p class="col-md-3">Số lượng</p>
            <div class="quantity buttons_added col-md-4">
              <input
                type="button"
                @click="decreaseAmount()"
                value="-"
                class="minus"
              /><input
                type="number"
                step="1"
                min="1"
                max="100"
                readonly
                name="quantity"
                v-model.number="amount"
                title="Qty"
                class="input-text qty text"
                size="4"
                pattern=""
                inputmode=""
              /><input
                type="button"
                @click="increaseAmount()"
                value="+"
                class="plus"
              />
            </div>
          </div>
          <br />
          <div class="row">
            <button @click="addProductsByQuantity(product)">
              Thêm vào giỏ hàng
            </button>
          </div>
        </div>
      </div>
    </div>
    <br />
    <div class="row">
      <h2 class="border-bottom">Mô tả sản phẩm</h2>
      <p>{{ product.Description }}</p>
    </div>
    <!-- <br />
    <br />
    <div class="row">
      <h2 class="border-bottom">Sản phẩm liên quan</h2>
    </div> -->
  </div>
</template>

<script>
export default {
  name: "ProductDetails",
  components: {},
  props: ["id"],
  data() {
    return {
      product: [],
      amount: 1,
    };
  },
  methods: {
    formatPrice(price) {
      return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    addProductsByQuantity(product) {
      product["quantity"] = this.amount;
      this.$store.commit("addProductToCart", product);
    },
    decreaseAmount() {
      if (this.amount > 1) {
        this.amount--;
      }
    },
    increaseAmount() {
      if (this.amount < 50) this.amount++;
    },
  },

  async created() {
    const response = await fetch(
      `http://localhost:3000/api/product/${this.id}`
    );
    this.product = await response.json();
    console.log(this.product);
  },
};
</script>

<style scoped>
.buffer-top {
  margin-top: 180px;
  padding-top: 10px;
  padding-bottom: 10px;
}
@media only screen and (max-width: 768px) {
  .buffer-top {
    padding-top: 300px;
  }
}
img {
  display: block;
  max-width: 400px;
  max-height: 400px;
  width: auto;
  height: auto;
}
.bgcolor {
  background-color: rgba(14, 12, 12, 0.08);
}
.redtext {
  color: red;
}
.quantity {
  display: inline-block;
}

.quantity .input-text.qty {
  width: 35px;
  height: 39px;
  padding: 0 5px;
  text-align: center;
  background-color: transparent;
  border: 1px solid #efefef;
}

.quantity.buttons_added {
  text-align: left;
  position: relative;
  white-space: nowrap;
  vertical-align: top;
}

.quantity.buttons_added input {
  display: inline-block;
  margin: 0;
  vertical-align: top;
  box-shadow: none;
}

.quantity.buttons_added .minus,
.quantity.buttons_added .plus {
  padding: 7px 10px 8px;
  height: 41px;
  background-color: #ffffff;
  border: 1px solid #efefef;
  cursor: pointer;
}

.quantity.buttons_added .minus {
  border-right: 0;
}

.quantity.buttons_added .plus {
  border-left: 0;
}

.quantity.buttons_added .minus:hover,
.quantity.buttons_added .plus:hover {
  background: #eeeeee;
}

.quantity input::-webkit-outer-spin-button,
.quantity input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  -moz-appearance: none;
  margin: 0;
}

.quantity.buttons_added .minus:focus,
.quantity.buttons_added .plus:focus {
  outline: none;
}
</style>
