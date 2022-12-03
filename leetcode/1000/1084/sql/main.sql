SELECT `product`.`product_id`, `product`.`product_name`
FROM `product` LEFT JOIN `sales` ON `product`.`product_id` = `sales`.`product_id`
GROUP BY `sales`.`product_id`
HAVING MIN(`sale_date`) >= '2019-01-01' AND MAX(`sale_date`) <= '2019-03-31'