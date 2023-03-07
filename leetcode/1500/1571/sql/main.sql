SELECT `w`.`name` `warehouse_name`, SUM(`p`.`width` * `p`.`height` * `p`.`length` * `w`.`units`) `volume`
FROM `warehouse` `w` JOIN `products` `p` USING (`product_id`)
GROUP BY `w`.`name`