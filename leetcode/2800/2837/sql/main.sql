SELECT `users`.`user_id`, `users`.`name`, IFNULL(SUM(`rides`.`distance`), 0) 'traveled distance'
FROM `users` LEFT JOIN `rides` USING (`user_id`)
GROUP BY `users`.`user_id`
ORDER BY `users`.`user_id`