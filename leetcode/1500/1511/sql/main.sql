SELECT `c`.`customer_id`, `c`.`name`
FROM `customers` `c`
    JOIN `orders` `o` USING (`customer_id`)
    JOIN `product` `p` USING (`product_id`)
WHERE `o`.`order_date` BETWEEN '2020-06-01' AND '2020-07-31'
GROUP BY `c`.`customer_id`
HAVING SUM(IF(MONTH(`o`.`order_date`) = 6, `o`.`quantity` * `p`.`price`, 0)) >= 100
    AND SUM(IF(MONTH(`o`.`order_date`) = 7, `o`.`quantity` * `p`.`price`, 0)) >= 100