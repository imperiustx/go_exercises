CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(50),
  `email` varchar(50),
  `password` varchar(50),
  `phone_number` varchar(10),
  `deliver_addresses` json,
  `avatar` json,
  `payment_method` json,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `user_ratings` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `restaurant_id` int UNIQUE NOT NULL,
  `food_id` int UNIQUE NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `restaurants` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(50),
  `email` varchar(50),
  `password` varchar(50),
  `phone_number` varchar(10),
  `address` varchar(50),
  `price_range` varchar(50),
  `images` json,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `foods` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `restaurant_id` int UNIQUE NOT NULL,
  `name` varchar(255),
  `description` text,
  `price` float,
  `images` json,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `restaurant_id` int UNIQUE NOT NULL,
  `items` json,
  `all_item_price` int,
  `voucher` int,
  `shipping_price` int,
  `total_price` float,
  `status` varchar(10),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `user_ratings` (`user_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `user_ratings` (`restaurant_id`);

ALTER TABLE `foods` ADD FOREIGN KEY (`id`) REFERENCES `user_ratings` (`food_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `foods` (`restaurant_id`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `orders` (`user_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `orders` (`restaurant_id`);
