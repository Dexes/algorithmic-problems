SELECT `product_name`, `year`, `price`
FROM `sales` INNER JOIN `product` ON `sales`.`product_id` = `product`.`product_id`