SELECT `users`.`user_id` AS `buyer_id`, `users`.`join_date`, IFNULL(COUNT(`order_date`), 0) AS `orders_in_2019`
FROM `users` LEFT JOIN `orders` ON `users`.`user_id` = `orders`.`buyer_id` AND YEAR(`order_date`) = '2019'
GROUP BY `users`.`user_id`