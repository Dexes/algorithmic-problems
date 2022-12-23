SELECT `sales`.`product_id`, SUM(`quantity`) AS `total_quantity`
FROM `sales` INNER JOIN `product` ON `sales`.`product_id` = `product`.`product_id`
GROUP BY `sales`.`product_id`