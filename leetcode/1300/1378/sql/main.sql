SELECT `u`.`unique_id`, `e`.`name`
FROM `employees` `e` LEFT JOIN `employeeUNI` `u` USING (`id`)