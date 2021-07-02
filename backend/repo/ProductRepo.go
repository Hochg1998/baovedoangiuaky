package repo

import (
	"fmt"

	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initProduct() {
	// var sony, xiaomi model.Manufacturer
	// Db.Where("name = ?", "Sony").First(&sony)
	// Db.Where("name = ?", "Xiaomi").First(&xiaomi)

	var cats, dogs, catsfood, dogsfood, catsitem, dogsitem model.Category
	Db.Where("name LIKE ?", "cats%").First(&cats)
	Db.Where("name LIKE ?", "dogs%").First(&dogs)
	Db.Where("name LIKE ?", "catsfood%").First(&catsfood)
	Db.Where("name LIKE ?", "dogsfood%").First(&dogsfood)
	Db.Where("name LIKE ?", "catsitem%").First(&catsitem)
	Db.Where("name LIKE ?", "dogsitem%").First(&dogsitem)

	catsfood_1 := model.Product{
		Name: "Xúc xích cho chó mèo ,bò gà vịt , phô mai, cá, dê",
		Description: `"Xúc xích dinh dưỡng chứa hàm lượng protein cao và các chất thiếu yếu khác giúp boss lớn khỏe 
		Phù hợp cho việc huấn luyện thú cưng hay ăn bữa phụ cho boss
		Loại xúc xích hạn sử dụng chỉ 1 năm ít chất bảo quản hơn các dòng xúc xích khác
		HSD: 4/2022"`,
		Price:    1500,
		Category: &catsfood,
		Image:    "public/img/xxbanle.jpg",
		IsSale:   true,
	}
	Db.Create(&catsfood_1)

	catsitem_1 := model.Product{
		Name: "Đệm lông nệm ổ nằm cho thú cưng , chó mèo ,êm ái",
		Description: `"#STUPIDCAT1312
		SHOP CÓ BÁN KÈM CHIẾU LÓT MÙA HÈ, KHÁCH MUỐN MUA KÈM CHIẾU VUI LÒNG NHẮN TIN CHO SHOP.
		Cam kết hàng loại 1, shop đã bán hơn 500 chiếc trong nửa năm và không nhận bất kỳ phản hồi nào tiêu cực. Hiện tại trên thị trường xuất hiện rất nhiều loại hàng trôi nổi, lượng bông nhồi không ổn định và chất vải nilon nằm rất bí. Hàng của shop chất vải cotton lông, riêng đệm pet là vải nhung cao cấp, bông nhồi nhiều nên đệm rất êm và thấm hút rất tốt. Dưới đệm là lớp chống trơn chống thấm cao cấp.
		-Đệm 45x35 XS - cún 3kg - miu 4kg
		-Đệm 50x40 S - cún 5kg - miu 6kg
		-Đệm 60x45 M - cún 10kg 
		-Đệm 70x55 L - cún 20kg 
		-Đệm 80x65 XL - cún 40kg 
		Lưu ý: kích thước trên là kích thước đo ở phần lớn nhất của đệm (bao gồm cả viền đệm) có thể sai số tối đa 2-3cm.
		#Demchocho #odemchomeo #demchomeo #demchopet #pet #forpet #chomeo #cho #meo #odemchocho #odem #dembongchomeo #pets #petshop 
		Mọi thắc mắc xin liên hệ shop :
		0817208868 - 0943214741  "`,
		Price:    199000,
		Category: &catsitem,
		Image:    "public/img/demmeo.jpg",
		IsSale:   true,
	}
	Db.Create(&catsitem_1)

	catsitem_2 := model.Product{
		Name: "Hạt khử mùi cho cát vệ sinh mèo",
		Description: `"
		Các hạt khử mùi đóng vai trò hấp phụ mùi, hấp thụ một phần nước tiểu, khóa nước tiểu và do đó làm giảm mùi hôi. 
		
		Cách sử dụng đơn giản, bạn chỉ cần rắc hạt lên chậu cát của mèo, nó sẽ giúp khử mùi hôi sau khi mèo đi vệ sinh. 
		ngoài ra, bạn có thể để túi hạt khu vực nơi mèo ở, đi vệ sinh để hấp thụ mùi , loại bỏ mùi hôi của thú cưng trong nhà
		
		Công nghệ sợi tiên tiến giúp loại bỏ mùi hôi nhanh chóng, giữ không khí trong lành lâu dài.
		
		Được làm từ chất liệu tự nhiên, không độc hại, an toàn khi sử dụng.
		
		Chứa các lợi khuẩn giúp chống vi khuẩn và khử mùi hôi hiệu quả.
		
		#hạt-khử-mùi
		#giảm-mùi-hôi
		#cát-vệ-sinh-mèo
		#phụ-kiện-thú-cưng"`,
		Price:    85000,
		Category: &catsitem,
		Image:    "public/img/catkhumui.jpg",
		IsSale:   false,
	}
	Db.Create(&catsitem_2)

	catsitem_3 := model.Product{
		Name: "Bàn cào móng hình tròn - ổ nằm tròn cho mèo",
		Description: `"🔯MÔ TẢ SẢN PHẨM
		I, Loại viền nhựa cao cấp
		✅ Kích thước: 41 x 41 x 7 cm, cho bé 5-7kg
		✅ Có 3 mẫu lựa chọn
		✅ Thiết kế chắc chắn, sử dụng lâu dài.
		✅ Đáy CÓ lớp nhựa hứng vụn giấy
		✅ Lõi dùng được hai mặt, có thể tháo rời lắp lõi mới
		✅ Mỗi lõi có thể sử dụng từ 2-6 tháng tùy tần suất cào móng của bé mèo
		
		II, Loại nhựa tai mèo cao cấp
		✅ Kích thước: 39x39x5cm cho bé 4-7kg
		✅ Có 2 màu lựa chọn: trắng, xanh
		✅ Đáy CÓ lớp nhựa hứng vụn giấy
		✅ Lõi dùng được hai mặt, có thể tháo rời lắp lõi mới
		✅ Mỗi lõi có thể sử dụng từ 2-6 tháng tùy tần suất cào móng của bé mèo
		
		III, Loại viền nhựa có  thêm que đồ chơi
		✅ Kích thước: 40x40x10cm cho bé 5-7kg
		✅ Có 3 màu
		✅ Thiết kế có thêm que đồ chơi giúp bé x2 độ giải trí
		✅ Đáy KHÔNG có lớp nhựa hứng vụn giấy
		✅ Lõi dùng được hai mặt, có thể tháo rời lắp lõi mới
		✅ Mỗi lõi có thể sử dụng từ 2-6 tháng tùy tần suất cào móng của bé mèo
		
		IV, loại đĩa bay viền nhựa
		✅ Kích thước: 38x38x5cm cho bé 3-6kg
		✅ Màu trắng
		✅ Thiết kế có thêm lỗ cắm que đồ chơi
		✅ Đáy có lớp nhựa hứng vụn giấy
		✅ Lõi dùng được hai mặt, có thể tháo rời lắp lõi mới
		✅ Mỗi lõi có thể sử dụng từ 2-6 tháng tùy tần suất cào móng của bé mèo
		
		V, Loại xạnh cào giấy xanh lục
		✅ Kích thước 32x32x5cm cho bé 1-4kg
		✅ Thiết kế viền nhựa cứng
		✅ Giá thành rẻ
		!!! Tuy nhiên, loại này không thể thay thế lõi cào :((
		
		 VI, Loại bàn càu ngầu lòi
		✅ Kích thước: 44x44x10cm cho bé  4-7kg
		✅ Màu hồng, xanh lá, vàng, xanh dương
		✅ Thiết kế có thêm lỗ cắm que đồ chơi
		✅ Đáy có lớp nhựa hứng vụn giấy
		✅ Lõi dùng được hai mặt, có thể tháo rời lắp lõi mới
		✅ Mỗi lõi có thể sử dụng từ 2-6 tháng tùy tần suất cào móng của bé mèo
		
		VII, Loại bàn cào bánh kem
		✅ Kích thước: 39x39x10cm cho bé 3-5kg và 50x50x10 cho bé 5-9kg (mua size to inb riêng cho shop)
		✅ Màu hồng và trắng
		✅ Size 50cm thiết kế có thêm lỗ cắm que đồ chơi
		✅ Đáy KHÔNG có lớp nhựa hứng vụn giấy
		✅ Lõi dùng được hai mặt, có thể tháo rời lắp lõi mới
		✅ Mỗi lõi có thể sử dụng từ 2-6 tháng tùy tần suất cào móng của bé mèo
		
		Liên hệ qua boxchat hoặc sđt 0817 208 868 để tư vấn thêm"`,
		Price:    240000,
		Category: &catsitem,
		Image:    "public/img/bancaomong.jpg",
		IsSale:   true,
	}
	Db.Create(&catsitem_3)

	catsitem_4 := model.Product{
		Name: "BALO Phi hành gia cho mèo trong suốt",
		Description: `"MÔ TẢ SẢN PHẨM
		🚚Các bạn ở Hà Nội có thể chọn Nowship hoặc Grab để nhận hàng hỏa tốc trong vòng 1 giờ nhé.
		
		🚚 Bạn cần giao hàng HỎA TỐC hoặc cần tư vấn thêm về sản phẩm. 📞 Hãy Chat, Gọi, iMessage hoặc Zalo cho Mình theo số: 0817.208.868 để được hỗ trợ nhé.
		
		♥️♥️ Yêu các bạn và các boss nhiều ạ ♥️♥️. 
		
		----------------------
		🎯 Balo (Ba lô) Phi Hành Gia Trong Suốt Vận Chuyển Chó Mèo (Kèm Thảm Lót) Giá Rẻ  🎯 
		
		Balo (Ba lô) Phi Hành Gia Trong Suốt Vận Chuyển Chó Mèo là phụ kiện không thể thiếu nếu các bạn đang nuôi boss mèo hoặc boss cún nhỏ. Boss nhà mình xinh thì phải cho ra đường khoe với mọi người phải không ạ. 
		
		📌Thông tin chung
		✔️ Kích thước (dài x rộng x cao ): 33 x 30 x 40cm . 
		✔️ Màu sắc : đỏ, vàng, xanh dương, xanh lá cây, đen, hồng, xanh dương nhạt
		
		📌Đặc điểm
		✔️ Balo phi hình gia trong suốt vận chuyển thú cưng được làm bằng vật liệu nhựa PVC và acrylic thân thiện với môi trường, chất lượng cao, không thấm nước và dễ dàng vệ sinh lau chùi.
		✔️ Chất liệu vải oxford sáng màu không thấm nước. Các lỗ phía trước và bên được thiết kế để thông gió cho thú cưng thở.
		✔️ Dây đeo vai thoải mái, thiết kế có đai ngực, vì vậy người đeo sẽ cảm thấy thoải mái khi vận chuyển thú cưng
		✔️Một bên lưới có chỗ để vừa một chai nước nhỏ hoặc đồ ăn vặt, phía bên kia là lỗ có khóa kéo để thú cưng có thể chui ra
		✔️ Dây đeo tay phù hợp với đường cong của bàn tay, có thể mang lại cảm giác thoải mái cho bàn tay để tránh mệt mỏi.
		✔️ Phù hợp với mèo và chó cún cỡ nhỏ với trọng lượng dưới 8kg 
		
		#balophihanhgia #balotrongsuot #balochochomeo #balochomeo
		
		----------------------
		
		🏠   Địa chỉ: B15 Đại Kim, Hoàng Mai , HN
		☎️ ĐT: 0817.208.868"`,
		Price:    179000,
		Category: &catsitem,
		Image:    "public/img/balochomeo.jpg",
		IsSale:   false,
	}
	Db.Create(&catsitem_4)

	dogsfood_1 := model.Product{
		Name: "Xúc xích thưởng - dinh dưỡng cho chó mèo 50cc",
		Description: `"XÚC XÍCH DINH DƯỠNG 


		Hạn sử dụng: 12 tháng kể từ ngày sản xuất
		
		THÔNG TIN SẢN PHẨM
		- Xúc Xích chứa Vitamin & khoáng chất cho cơ thể những thú cưng luôn khỏe mạnh, linh hoạt. Chất xơ cho hệ thống tiêu hóa tốt, tăng cường khả năng hấp thu chất dinh dưỡng, góp phần vào sự phát triển đồng đều cho thú cưng.
		- Chọn nguyên liệu tươi tự nhiên chất lượng cao, tất cả các nguyên liệu được kiểm tra nghiêm ngặt.
		- Giúp ổn định hệ thống vi khuẩn đường ruột, tiêu hóa khỏe, bé ăn ngon
		- Xúc Xích với đầy đủ dưỡng chất giúp bé ăn ngon.
		- Sản phẩm chứa đầy đủ các dưỡng chất thiết yếu, hỗ trợ cho quá trình hoạt động và phát triển của thú cưng nhà bạn, bao gồm:
		- Chất đạm: được chọn từ thịt có chất lượng, cung cấp acid amin cần thiết để xây dựng cơ bắp khỏe mạnh.
		- Chất béo: cung cấp năng lượng và nguồn acid béo thiết yếu để duy trì một làn da khỏe mạnh và một bộ lông bóng mượt.
		- Vitamin: giúp cơ thể phát triển khỏe mạnh và tăng cường hệ thống miễn dịch.
		- Khoáng chất: giúp cơ thể phát triển khỏe mạnh, xương chắc và răng khỏe.
		Không thêm chất bảo quản Không thêm chất màu Không thêm chất hấp dẫn
		#katholic - 0817208868 Hân hạnh phục vụ !!!"`,
		Price:    65000,
		Category: &dogsfood,
		Image:    "public/img/xx50cc.jpg",
		IsSale:   false,
	}
	Db.Create(&dogsfood_1)

	catsitem_5 := model.Product{
		Name: "Cattree cho mèo Cat tree nhà cây thư giãn",
		Description: `"Cat tree giúp mèo xả stress đùa nghịch giảm cân khi chủ vắng nhà hoặc cô đơn không có bạn đời bên cạnh
		Còn có bạn đời xả stress cùng thì càng vui 
		NHÀ CÂY CHO MÈO CÓ THANG LEO TRÈO 
		Nhà cây cho mèo  là một trong những dòng sản phẩm cho mèo được ưa chuộng nhất hiện nay. Sở hữu một  sẽ giúp bạn hạn chế sự hư hỏng của rất nhiều vật dụng trong gia đình trước sự tinh nghịch của những chú mèo.“Bản năng cào” của loài mèo nhằm mục đích đánh dấu và cảnh báo những chú mèo khác về lãnh thổ của chúng. Đồng thời loại bỏ lớp móng cũ và làm bộ vuốt của mình sắc bén hơn.
		Đặc Điểm Sản Phẩm Cây Nhà Cho Mèo 
		Chất Liệu
		Nhà cây cho mèo  được sử dụng chất liệu gỗ Mdf phủ 2 mặt có độ bền cao kết hợp cùng dây thừng sisal. Đây là loại chất liệu được ưa chuộng nhất trên thị trường hiện nay bởi độ bền, an toàn cũng như tính thẩm mỹ cao.
		
		Ưu Điểm Sản Phẩm Cây Nhà Cho Mèo
		
		Thiết kế hiện đại, thích hợp cho mèo ở và decor nội thất.
		Độ bền cao so với các sản phẩm cùng loại được thiết kế bằng chất liệu khác.
		Tính thẩm mỹ cao, phù hợp với không gian hiện đại.
		Dễ dàng vệ sinh và vận chuyển.
		Giá cả phù hợp với mọi đối tượng.
		
		>>>> HỖ TRỢ LẮP ĐẶT NẾU KHÁCH HÀNG CÓ YÊU CẦU
		>>>> Sản phẩm đóng gói kèm bộ lắp đặt và bản hướng dẫn lắp đặt ( nữ cũng dễ dàng lắp đặt)"`,
		Price:    990000,
		Category: &catsitem,
		Image:    "public/img/cattreecui.jpg",
		IsSale:   true,
	}
	Db.Create(&catsitem_5)

	catsitem_6 := model.Product{
		Name: "Cattree Cat tree Nhà cây cho mèo thư giãn",
		Description: `"Cây cào móng giúp mèo mài móng, giảm stress, chạy nhẩy leo trèo và ngủ nướng được luôn,  giúp boss đỡ buồn khi ở nhà 1 mình
		Size : 116cm
		130cm
		136cm
		150cm
		Chất liệu : gỗ bọc dây thừng sợi gai, đệm mút bọc nỉ nhung
		Công dụng: để các bé mèo cào móng, giảm stress, hạn chế phá đồ, cào hư đồ trong nhà.
		Là một hình thức thể dục, tận hưởng khoảng thời gian thoải mái, thư giãn cho các bé mèo.
		
		
		----------------------------------------------
		
		#Katholic - PHUKIENTHUCUNG ❤️ cung cấp phụ kiện, đồ dùng thú cưng - Chất lượng sản phẩm tốt - Dịch vụ chăm sóc khách hàng trực tuyến ❤️
		
		SĐT: 0817208868
		#chó #mèo #petshop #quanao #quanyemchomeo #đochomeo #aochomeo #dobochomeo #thoitrangchoboss
		 #DOCHOICHOMEO #diabay #dochoichochomeo #phukienchomeo #phukienthucung #trucaomong"`,
		Price:    1250000,
		Category: &catsitem,
		Image:    "public/img/cattreexin.jpg",
		IsSale:   false,
	}
	Db.Create(&catsitem_6)

	catsitem_7 := model.Product{
		Name:        "Bát ăn sứ cho mèo bát sứ ăn chống gù lưng",
		Description: "Sofa Miami 2 chỗ là một thiết kế tối giản cho không gian phòng khách hiện đại, chất liệu sofa vải cao cấp trên tông màu xám nhạt rất dễ dàng phối hợp cùng các sản phẩm khác. Kích thước nhỏ gọn 2 chỗ sẽ phù hợp cho các căn hộ, hoặc những góc nhỏ trong ngôi nhà của bạn.",
		Price:       135000,
		Category:    &catsitem,
		Image:       "public/img/batsuunc.jpg",
		IsSale:      true,
	}
	Db.Create(&catsitem_7)

	dogsitem_1 := model.Product{
		Name: "Máy uống nước tự động cho chó , mèo thú cưng ",
		Description: `"Máy uống nước tự động cho chó mèo
		Thương hiệu : Jumping dog
		Xuất xứ: nội địa Trung
		————————————
		Máy uống nước tự động vừa tiện lợi cho Sen lẫn Boss lại còn vừa sạch sẽ
		✨Ưu điểm:
		- Lọc các tạp chất trong nước
		- Cải thiện Ion kim loại trong nước
		- Cải thiện hương vị, Clo dư và vi sinh vật
		- Dễ dàng vệ sinh
		———————————————
		Một bộ sẽ bao gồm: máy, bộ lọc, vòi con vịt, dây cắm USB
		Kích thước: 18.5x18,5x18cm
		Dung tích: 2,5L
		Màu: Trắng , Hồng , Xanh
		———————————————
		Khách nên mua thêm miếng lọc để thay
		1 miếng lọc dùng được 2-6 tuần"`,
		Price:    340000,
		Category: &dogsitem,
		Image:    "public/img/mayuongnuoc.jpg",
		IsSale:   false,
	}
	Db.Create(&dogsitem_1)

	catsfood_2 := model.Product{
		Name: "đơn hàng custom tùy chỉnh thùng xúc xích 200cc",
		Description: `"xúc xích thơm ngon giá hạt rẻ phù hợp với mọi giai đoạn phát triển của bé
		1 cây nặng 15g có 5 vị bò gà vịt phô mai cừu
		KHÔNG TỰ Ý ĐẶT NẾU KHÔNG INB SHOP"`,
		Price:    380000,
		Category: &catsfood,
		Image:    "public/img/xx200cc.jpg",
		IsSale:   true,
	}
	Db.Create(&catsfood_2)

	dogsitem_2 := model.Product{
		Name: "Lõi bàn cào móng thay thế tiện lợi",
		Description: `Đồ tặng kèm cho các đơn hàng lớn deal trực tiếp với shop để xin quà \
		dây thừng hình xương , mài răng cho cún`,
		Price:    105000,
		Category: &dogsitem,
		Image:    "public/img/loibancao.jpg",
		IsSale:   true,
	}
	Db.Create(&dogsitem_2)

	sofa_5 := model.Product{
		Name:        "Dây thừng hình xương , thư giãn mài răng cho cún",
		Description: "Sofa Dubai 2.5 chỗ với đường nét mỏng đảm bảo cái nhìn nhẹ nhàng và thanh thoát. Thiết kế sofa 2 chỗ nhưng vẫn mang lại cảm giác chỗ ngồi rộng hơn. Sofa Dubai có 2 màu nâu và kem để chọn lựa phù hợp cho không gian phòng khách hiện đại của gia đình bạn.",
		Price:       16500000,
		Category:    &dogsitem,
		Image:       "public/img/recan.jpg",
		IsSale:      false,
	}
	Db.Create(&sofa_5)

	catsitem_8 := model.Product{
		Name: "Chậu cát vệ sinh cho mèo hình chữ nhật",
		Description: `chậu cát có 2 size dành cho bé mèo
		size S: 30x38cm cho mèo dưới 4kg
		size M: 30x43cm cho mèo dưới 6kg
		Chậu cát dáng chữ nhật tăng diện tích sử dụng màu trắng tao nhã dễ dàng bài trí trong căn phòng của ban`,
		Price:    189000,
		Category: &catsitem,
		Image:    "public/img/chaucat.jpg",
		IsSale:   true,
	}
	Db.Create(&catsitem_8)

	dogsitem_3 := model.Product{
		Name: "Cây lăn lông mèo trên quần áo, chăn ga gối",
		Description: `cây lăn lông giúp loại bỏ lông bám trên giường chiếu quần áo và các loại vải khác
		có 3 size: to và bé và loại cao cấp
		Đặc điểm:
		1. Sản phẩm này có thể loại bỏ lông thú cưng khỏi đồ nội thất nhưng không thể loại bỏ bóng lông.
		2. Phạm vi sử dụng rộng rãi, như dùng cho ghế sofa, thảm, giường hoặc đồ nội thất khác.
		3. Cầm 1 tay để sử dụng, tiết kiệm nhiều thời gian và công sức, dễ dàng và thuận tiện.
		4. Thiết kế 2 chiều có hiệu quả cao hơn.
		5. Được làm từ chất liệu chất lượng cao, sản phẩm này có tuổi thọ dài và dễ dàng làm sạch.
		
		Chi tiết sản phẩm loại cao cấp - hàng sẵn
		1. Sản phẩm mới 100% và chất lượng cao
		2. Chất liệu: ABS
		3. Màu sắc: Như hình ảnh hiển thị
		4. Kích thước: 19.5cm * 6.5cm * 12.5cm/19.5*19.0*5cm
		
		Gói hàng bao gồm: 1 x Dụng cụ loại bỏ lông thú cưng
		
		Lưu ý:
		Vui lòng cho phép sai số kích thước nhỏ 2-3% do đo lường thủ công. 
		Vui lòng kiểm tra Bảng kích thước cẩn thận trước khi mua sản phẩm. 
		Xin lưu ý do ánh sáng và màn hình, màu sắc sản phẩm có thể có chênh lệch nhẹ, điều này là chấp nhận được.`,
		Price:    69000,
		Category: &dogsitem,
		Image:    "public/img/cuonlong.jpg",
		IsSale:   false,
	}
	Db.Create(&dogsitem_3)

	catsitem_9 := model.Product{
		Name: "Que đồ chơi con chuột kèm chuông",
		Description: `Que đồ chơi cho mèo dài 45cm kèm chuông xinh xắn cho bé mèo vui chơi thỏa thích.
		Chất liệu gỗ chắc chắn kèm chuông nghe vui tai`,
		Price:    22000,
		Category: &catsitem,
		Image:    "public/img/thanhtreumeo.jpg",
		IsSale:   true,
	}
	Db.Create(&catsitem_9)

	catsitem_10 := model.Product{
		Name: "Bát ăn , khay ăn , gấp gọn tiên lợi ",
		Description: `
		Thông tin chi tiết:
		Nhẹ, bền và có thể tái sử dụng - chống trượt, bát ăn uống đi du lịch cho thú cưng. Lý tưởng cho những chú cún nhỏ và vừa.
		Bát uống nước dành cho cún cưng.
		Kích thước cầm tay, bát này có thiết kế gấp lại. Đơn giản chỉ cần bật lên và sau đó gấp lại - vừa túi, ví, balo hoặc hộp găng tay, dễ dàng để làm sạch.
		An toàn trong máy rửa chén.
		
		Loại sản phẩm: Bát ăn uống cho thú cưng
		Chất liệu: Silicone
		Phù hợp dùng cho: Cún, mèo, thú cưng, v.v.
		Phù hợp dùng vào dịp: Du lịch, ở nhà, ngoài trời, vv
		Các tính năng: Khung đen, có thể gập lại, kích thước cầm tay, đựng nước thức ăn
		Đường kính miệng bát: 13cm / 5.12" (xấp xỉ)
		Đường kính đáy bát: 9cm / 3,54" (xấp xỉ)
		Chiều cao: 1cm - 5.0cm / 0.39" - 1.97" (xấp xỉ)
		
		Ghi chú:
		Do chênh lệch cài đặt ánh sáng và màn hình, màu sắc của sản phẩm có thể hơi khác so với hình ảnh.
		Vui lòng cho phép chênh lệch kích thước nhỏ do đo lường thủ công khác nhau.
		
		Gói hàng bao gồm:
		1 x Bát ăn uống cho thú cưng`,
		Price:    19000,
		Category: &catsitem,
		Image:    "public/img/batuongnuoc.jpg",
		IsSale:   true,
	}
	Db.Create(&catsitem_10)

}

func FindProductById(Id int) *model.Product {
	var product *model.Product
	Db.Preload("Category").First(&product, Id)
	fmt.Println(product)
	return product
}

func FindProductByCategory(categoryname string) (result []*model.Product) {
	switch categoryname {
	case "cats":
		{
			Db.Where("category_id=?", "1").Find(&result)
			return result
		}
	case "dogs":
		{
			Db.Where("category_id=?", "2").Find(&result)
			return result
		}
	case "catsfood":
		{
			Db.Where("category_id=?", "3").Find(&result)
			return result
		}
	case "dogsfood":
		{
			Db.Where("category_id=?", "4").Find(&result)
			return result
		}
	case "catsitem":
		{
			Db.Where("category_id=?", "5").Find(&result)
			return result
		}
	case "dogsitem":
		{
			Db.Where("category_id=?", "6").Find(&result)
			return result
		}
	default:
		{
			Db.Find(&result)
			return result
		}
	}
}

func FindProductSale(issale string) (result []*model.Product) {
	switch issale {
	case "true":
		{
			Db.Where("is_sale=?", true).Find(&result)
			return result
		}
	case "false":
		{
			Db.Where("is_sale=?", false).Find(&result)
			return result
		}
	default:
		{
			Db.Find(&result)
			return result
		}
	}
}

func FindProductByName(name string) (result []*model.Product) {
	Db.Where("name LIKE ?", "%"+name+"%").Find(&result)
	return result
}
