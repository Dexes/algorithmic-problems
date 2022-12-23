SELECT `seller_id`
FROM `sales`
GROUP BY `seller_id`
HAVING SUM(`price`) = (SELECT MAX(`sum`) AS `sum` FROM (SELECT SUM(`price`) AS `sum` FROM `sales` GROUP BY `seller_id`) AS `x`)