
# Project


## Backend
	Code backend nằm trong thư mục backend
	Chạy code backend trên port 3000
	Chạy code bằng câu lệnh go run app.go
	
	API get all products: http://localhost:3000/api/product?limit={limit}&offset={offset}
		VD: http://localhost:3000/api/product?limit=17&offset=0
	API get product by id: http://localhost:3000/api/product/:id
		VD: http://localhost:3000/api/product/1
	API get product isSale: http://localhost:3000/api/product/issale/true?limit=6&offset=${index *
          6}
		VD: http://localhost:3000/api/product/issale/true?limit=6&offset=0}
		
## Frontend
	Chạy lệnh npm install để cài đặt các modules liên quan
	Code frontend nằm trong thư mục frontend
	Code frontend chạy trên cổng 8080
	v
