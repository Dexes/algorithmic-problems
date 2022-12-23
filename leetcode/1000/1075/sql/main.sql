SELECT `project_id`, ROUND(AVG(`experience_years`), 2) AS `average_years`
FROM `project` INNER JOIN `employee` USING (`employee_id`)
GROUP BY `project_id`