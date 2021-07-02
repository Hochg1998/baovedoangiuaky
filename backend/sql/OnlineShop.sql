CREATE TABLE `products` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `price` int,
  `category_id` int,
  `image` varchar(255) NOT NULL,
  `is_sale` TINYINT
);
CREATE TABLE `categories` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL
);
ALTER TABLE
  `products`
ADD
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`ID`);