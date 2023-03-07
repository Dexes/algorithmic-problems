SELECT DATE_FORMAT(`o`.`order_date`, '%Y-%m') `month`, COUNT(`o`.`order_id`) `order_count`, COUNT(DISTINCT `o`.`customer_id`) `customer_count`
FROM `orders` `o`
WHERE `o`.`invoice` > 20
GROUP BY `month`