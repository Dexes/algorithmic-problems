SELECT `p`.`name`,
       IFNULL(SUM(`i`.`rest`), 0)     `rest`,
       IFNULL(SUM(`i`.`paid`), 0)     `paid`,
       IFNULL(SUM(`i`.`canceled`), 0) `canceled`,
       IFNULL(SUM(`i`.`refunded`), 0) `refunded`
FROM `product` `p` LEFT JOIN `invoice` `i` USING (`product_id`)
GROUP BY `p`.`product_id`
ORDER BY `p`.`name`