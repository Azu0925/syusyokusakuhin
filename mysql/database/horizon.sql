CREATE DATABASE IF NOT EXISTS `horizon` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE `horizon`;

CREATE TABLE `users` (
  `user_id` int(32) NOT NULL,
  `name` text NOT NULL
)