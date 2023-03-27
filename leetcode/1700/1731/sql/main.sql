SELECT `a`.`employee_id`, `a`.`name`, COUNT(`b`.`employee_id`) `reports_count`, ROUND(AVG(`b`.`age`)) `average_age`
FROM `employees` `a` INNER JOIN `employees` `b` ON `a`.`employee_id` = `b`.`reports_to`
GROUP BY `a`.`employee_id`
ORDER BY `a`.`employee_id`