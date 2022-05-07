SELECT `customer_id`, COUNT(`visits`.`visit_id`) AS `count_no_trans`
FROM `visits` LEFT JOIN `transactions` ON `visits`.`visit_id` = `transactions`.`visit_id`
WHERE `transactions`.`visit_id` IS NULL
GROUP BY `customer_id`