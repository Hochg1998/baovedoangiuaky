
import { createStore } from 'vuex'

// Create a new store instance.
const store = createStore({
    strict: true,
    state() {
        //data
        return {
            discountCode: "",
            discountPercentage: 0,
            deletedProducts: [],
            cart: [],
            promotions: [
                {
                    code: "SUMMER",
                    discount: 50,
                },
                {
                    code: "AUTUMN",
                    discount: 40,
                },
                {
                    code: "WINTER",
                    discount: 30,
                },
            ],
        }
    },
    mutations: {
        //method
        //thêm object + quantity vao giỏ hàng
        applyDiscount(state) {
            state.promotions.reduce((inputCode, curr) => {
                if (inputCode === "") {
                    state.discountPercentage = 0;
                }
                else if (inputCode === curr.code) {
                    console.log(curr.discount);
                    state.discountPercentage = curr.discount;
                    return curr.discount;
                }
            }, state.discountCode);

            if (state.discountPercentage === 0) {
                alert("Mã giảm giá không tồn tại");
            }
        },
        updateDiscountCode(state, discountCode) {
            state.discountCode = discountCode
        },
        deleteProduct(state, id) {
            console.log(id);
            if (confirm("Are you sure about dat?")) {
                state.cart = state.cart.filter((product) => {
                    if (product.ID === id) {

                        state.deletedProducts.push(product);
                    }

                    return product.ID != id;
                });
            }
            console.log(state.cart);
        },
        rollbackProduct(state) {
            console.log(state.deletedProducts);
            state.deletedProducts.forEach((product) => state.cart.push(product));
            state.deletedProducts = [];
        },
        addProductToCart(state, product) {
            if (state.cart.length === 0) {
                state.cart.push(product);
            } else {
                let count = 0;

                state.cart.find(productInCart => {
                    if (productInCart.ID === product.ID) {
                        let first = parseInt(productInCart.quantity);
                        console.log(first);

                        let second = parseInt(product.quantity);
                        console.log(second);
                        let total = first + second
                        console.log(total);
                        productInCart.quantity = total
                        count++;

                    }

                })
                if (count === 0) {
                    console.log(count);
                    state.cart.push(product);
                }
            }


        },
        increaseAmount(state, id) {
            state.cart.find(product => {
                if (product.ID === id) {
                    if (product.quantity < 50) {
                        product.quantity++;
                    }
                }
            })
        },
        decreaseAmount(state, id) {
            state.cart.find(product => {
                if (product.ID === id) {
                    if (product.quantity > 1) {
                        product.quantity--;
                    }
                }
            })
        },
    },
    actions: {
        //method async
        //dùng để thay đổi category trong cart
    },
    getters: {
        subTotal(state) {
            let subTotal = 0;
            state.cart.forEach((element) => {
                subTotal += element.Price * element.quantity;
            });
            return subTotal;
        },
        calculatedTax(state) {
            let subTotal = 0;
            state.cart.forEach((element) => {
                subTotal += element.Price * element.quantity;
            });

            return (subTotal / 100) * 10;
        },
        // //computed
        // totalProductInCart() {
        //     console.log(this.$store.state.cart);
        //     return this.$store.state.cart;
        // },
        // countTotal() {
        //     let totalProducts = 0;
        //     this.products.forEach((product) => {
        //         totalProducts += parseInt(product.quantity);
        //     });
        //     return parseInt(totalProducts);
        // },
    },
})



// Install the store instance as a plugin
export default store