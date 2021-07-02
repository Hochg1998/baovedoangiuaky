<template>
  <div class="row">
    <div class="col-md-6">
      <p class="text-start">Have A Promo Code?</p>

      <div class="input-group mb-3">
        <input
          :value="discountCode"
          @input="updateDiscountCode"
          type="text"
          class="form-control"
          placeholder="Mã giảm giá"
          aria-label="Recipient's username"
          aria-describedby="button-addon2"
        />
        <button
          class="btn btn-success"
          type="button"
          id="button-addon2"
          @click="applyDiscount"
        >
          Áp dụng
        </button>
      </div>
    </div>
    <div class="col-md-6">
      <h5 class="text-end">
        Subtotal &emsp;&emsp; <span>{{ formatPrice(subTotal) }}</span>
      </h5>
      <h5 class="text-end">
        Tax &emsp;&emsp;&emsp;<span>{{ formatPrice(calculatedTax) }}</span>
      </h5>
      <h5 class="text-end">
        Giảm &emsp;&emsp;&emsp;{{
          formatPrice((subTotal / 100) * discountPercentage)
        }}
      </h5>
      <h5 class="text-end">
        Total &emsp;&emsp;<span>{{
          formatPrice(
            subTotal + calculatedTax - (subTotal / 100) * discountPercentage
          )
        }}</span>
      </h5>
      <div class="text-end">
        <button class="btn btn-success">Checkout ></button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Checkout",
  components: {},

  props: {
    products: Object,
    promotions: Object,
  },
  methods: {
    updateDiscountCode(e) {
      this.$store.commit("updateDiscountCode", e.target.value);
    },
    applyDiscount() {
      this.$store.commit("applyDiscount");
    },
    formatPrice(price) {
      return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
  },

  computed: {
    discountPercentage() {
      return this.$store.state.discountPercentage;
    },
    discountCode() {
      return this.$store.state.discountCode;
    },
    subTotal() {
      return this.$store.getters.subTotal;
    },
    calculatedTax() {
      return this.$store.getters.calculatedTax;
    },
  },
};
</script>

<style scoped>
h5 {
  color: blue;
}
</style>
