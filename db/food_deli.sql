CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(50),
  `email` varchar(50),
  `password` varchar(50),
  `phone_number` varchar(10),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `user_ratings` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `restaurant_id` int UNIQUE NOT NULL,
  `food_id` int UNIQUE NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `user_addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `address_id` int UNIQUE NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `user_payment_methods` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `payment_method_id` int UNIQUE NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_address` varchar(255),
  `latitude` decimal(8,6),
  `longitude` decimal(9,6),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `payment_methods` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `bank_name` varchar(50),
  `user_name` varchar(50),
  `number` varchar(50),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `restaurants` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(50),
  `email` varchar(50),
  `password` varchar(50),
  `phone_number` varchar(10),
  `address` varchar(50),
  `price_range` varchar(50),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `restaurant_addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `restaurant_id` int UNIQUE NOT NULL,
  `address_id` int UNIQUE NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `foods` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `restaurant_id` int UNIQUE NOT NULL,
  `name` varchar(255),
  `description` text,
  `price` float,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `restaurant_id` int UNIQUE NOT NULL,
  `all_item_price` int,
  `voucher` int,
  `shipping_price` int,
  `total_price` float,
  `status` varchar(20),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `items` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int UNIQUE NOT NULL,
  `restaurant_id` int UNIQUE NOT NULL,
  `food_id` int UNIQUE NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `user_ratings` (`user_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `user_ratings` (`restaurant_id`);

ALTER TABLE `foods` ADD FOREIGN KEY (`id`) REFERENCES `user_ratings` (`food_id`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `user_addresses` (`user_id`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`id`) REFERENCES `user_addresses` (`address_id`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `user_payment_methods` (`user_id`);

ALTER TABLE `payment_methods` ADD FOREIGN KEY (`id`) REFERENCES `user_payment_methods` (`payment_method_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `restaurant_addresses` (`restaurant_id`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`id`) REFERENCES `restaurant_addresses` (`address_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `foods` (`restaurant_id`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `orders` (`user_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `orders` (`restaurant_id`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `items` (`user_id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `items` (`restaurant_id`);

ALTER TABLE `foods` ADD FOREIGN KEY (`id`) REFERENCES `items` (`food_id`);
