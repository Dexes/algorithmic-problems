SELECT `person_id`, CONCAT(`name`, '(', SUBSTR(`profession`, 1, 1), ')') AS `name`
FROM `person`
ORDER BY `person_id` DESC