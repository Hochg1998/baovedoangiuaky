# onlineshop

## Project setup

```
npm install
```

### Compiles and hot-reloads for development

```
npm run serve
```

### Compiles and minifies for production

```
npm run build
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).

# Cấu trúc ứng dụng

## 1. Giới thiệu

Dự án sử dụng Vuejs 3.0

1. Cấu trúc thư mục có thêm những sub package như router, views, components, store
2. Bổ xung Vuex để quản lý global state
3. Sử dụng Router để định dạng route

## 2. Cấu trúc lại thư mục

.
├── public <-- Chứa file index để hiện thị toàn bộ trang
│ └── index.html
├── src <-- Chứa toàn bộ src code của dự án
│ ├── assets <--- lưu các logo , banner
│ ├── components <-- chứa các component sử dụng trong view
│ ├── router <-- cấu hình đường dẫn
│ ├── store <-- sử dụng vuex để lưu global state sử dụng trong nhiều component
│ └── views <-- ứng với mỗi path trong router là một component ở trong views
│
├── browserlistrc
├── .gitignore <-- Những file không được upload lên git sẽ được ghi vào đây
├── babel.config.js <-- Compile về javascript thuần
├── package-lock.json <-- Cầu hình babel
├── package.json <-- cấu hình vuejs
└── ReadMe.md <-- File này quan trọng nhất, làm gì thì làm, luôn phải viết document!

## 3. Cấu hình router sử dụng vue router

Không sử dụng lazyload
{
path: '/',
name: 'Home',
component: Home
},
và sử dụng lazyload
{
path: '/about',
name: 'About',
// route level code-splitting
// this generates a separate chunk (about.[hash].js) for this route
// which is lazy-loaded when the route is visited.
component: () => import(/_ webpackChunkName: "about" _/ '../views/About.vue')

},

## 4. Sử dụng Vuex để quản lý global state

const store = createStore({
strict: true,
state() {
//data

},
mutations: {
//method
//thêm object + quantity vao giỏ hàng

},
actions: {
//method async
//dùng để thay đổi category trong cart
},
getters: {
sử lý computed
},
})
