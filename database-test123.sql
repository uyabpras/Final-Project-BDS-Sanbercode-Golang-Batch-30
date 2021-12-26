-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 26, 2021 at 02:43 PM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 8.0.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test123`
--

-- --------------------------------------------------------

--
-- Table structure for table `addres`
--

CREATE TABLE `addres` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL,
  `addresline` longtext DEFAULT NULL,
  `postalcode` bigint(20) DEFAULT NULL,
  `city` longtext DEFAULT NULL,
  `state` longtext DEFAULT NULL,
  `country` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `addres`
--

INSERT INTO `addres` (`id`, `user_id`, `created_at`, `update_at`, `addresline`, `postalcode`, `city`, `state`, `country`) VALUES
(1, 1, '2021-12-26 20:42:15.000', NULL, 'jogkarta', 65555, 'yogyakarta', 'yogyakarta', 'indonesia');

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `store_id` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `carts`
--

INSERT INTO `carts` (`id`, `user_id`, `product_id`, `store_id`) VALUES
(3, 1, 2, 1),
(4, 1, 3, 1);

-- --------------------------------------------------------

--
-- Table structure for table `checkouts`
--

CREATE TABLE `checkouts` (
  `id` bigint(20) NOT NULL,
  `store_id` bigint(20) DEFAULT NULL,
  `orderitem_id` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `checkouts`
--

INSERT INTO `checkouts` (`id`, `store_id`, `orderitem_id`, `created_at`) VALUES
(1, 1, 1, '2021-12-26 20:40:57.000'),
(2, 1, 2, '2021-12-26 20:40:57.000');

-- --------------------------------------------------------

--
-- Table structure for table `inventories`
--

CREATE TABLE `inventories` (
  `product_id` bigint(20) DEFAULT NULL,
  `store_id` bigint(20) DEFAULT NULL,
  `stock` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `inventories`
--

INSERT INTO `inventories` (`product_id`, `store_id`, `stock`, `created_at`, `update_at`) VALUES
(4, 1, 99, '2021-12-26 20:40:14.000', NULL),
(1, 1, 10, '2021-12-26 20:40:14.000', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `orderitems`
--

CREATE TABLE `orderitems` (
  `id` bigint(20) NOT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `addres_id` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orderitems`
--

INSERT INTO `orderitems` (`id`, `product_id`, `user_id`, `price`, `created_at`, `addres_id`) VALUES
(1, 1, 1, 1000000, '2021-12-26 20:39:21.000', NULL),
(2, 4, 1, 100000, '2021-12-26 20:39:21.000', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` bigint(20) NOT NULL,
  `name` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `price`, `created_at`, `update_at`) VALUES
(1, 'meja gaming', 1000000, '2021-12-26 20:36:57.000', NULL),
(2, 'speaker', 300000, '2021-12-26 20:36:57.000', NULL),
(3, 'baju polos', 25000, '2021-12-26 20:38:19.000', NULL),
(4, 'celana', 100000, '2021-12-26 20:38:19.000', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `stores`
--

CREATE TABLE `stores` (
  `id` bigint(20) NOT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `city` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `stores`
--

INSERT INTO `stores` (`id`, `username`, `password`, `city`, `created_at`, `update_at`) VALUES
(1, 'giant', '$2a$10$M4qN7Sgq9IRKL4VfRFoWTewyEJez2NXZrl4ePyMPHailX01WK1suO', 'bandung', '2021-12-26 19:56:39.514', '2021-12-26 20:17:15.496');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `username` varchar(191) NOT NULL,
  `password` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `created_at`, `update_at`) VALUES
(1, 'udin', '$2a$10$Hto5/nc3FsjjQt7mDN6iue0ltByFUx22S4UW/L.61FDUoqjN/Du8W', '2021-12-25 21:00:54.028', '2021-12-26 20:15:59.594');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `addres`
--
ALTER TABLE `addres`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_addres` (`user_id`);

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_cart` (`user_id`),
  ADD KEY `fk_products_cart` (`product_id`),
  ADD KEY `fk_stores_cart` (`store_id`);

--
-- Indexes for table `checkouts`
--
ALTER TABLE `checkouts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_orderitems_checkout` (`orderitem_id`),
  ADD KEY `fk_stores_checkout` (`store_id`);

--
-- Indexes for table `inventories`
--
ALTER TABLE `inventories`
  ADD KEY `fk_stores_inventory` (`store_id`),
  ADD KEY `fk_products_inventory` (`product_id`);

--
-- Indexes for table `orderitems`
--
ALTER TABLE `orderitems`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_addres_orderitem` (`addres_id`),
  ADD KEY `fk_products_orderitem` (`product_id`),
  ADD KEY `fk_users_orderitem` (`user_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stores`
--
ALTER TABLE `stores`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `password` (`password`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `addres`
--
ALTER TABLE `addres`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `checkouts`
--
ALTER TABLE `checkouts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `orderitems`
--
ALTER TABLE `orderitems`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `stores`
--
ALTER TABLE `stores`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `addres`
--
ALTER TABLE `addres`
  ADD CONSTRAINT `fk_addres_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `fk_users_addres` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `fk_carts_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_carts_store` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`),
  ADD CONSTRAINT `fk_products_cart` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_stores_cart` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`),
  ADD CONSTRAINT `fk_users_cart` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `checkouts`
--
ALTER TABLE `checkouts`
  ADD CONSTRAINT `fk_orderitems_checkout` FOREIGN KEY (`orderitem_id`) REFERENCES `orderitems` (`id`),
  ADD CONSTRAINT `fk_stores_checkout` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`);

--
-- Constraints for table `inventories`
--
ALTER TABLE `inventories`
  ADD CONSTRAINT `fk_inventories_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_products_inventory` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_stores_inventory` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`);

--
-- Constraints for table `orderitems`
--
ALTER TABLE `orderitems`
  ADD CONSTRAINT `fk_addres_orderitem` FOREIGN KEY (`addres_id`) REFERENCES `addres` (`id`),
  ADD CONSTRAINT `fk_products_orderitem` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_users_orderitem` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
