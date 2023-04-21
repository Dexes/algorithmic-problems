SELECT `s`.`user_id`, SUM(`p`.`price` * `s`.`quantity`) `spending`
FROM `sales` `s` JOIN `product` `p` USING (`product_id`)
GROUP BY `s`.`user_id`
ORDER BY 2 DESC