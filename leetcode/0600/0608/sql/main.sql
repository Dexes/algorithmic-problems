SELECT `id`, 'Root' AS `Type`
FROM `tree` WHERE `p_id` IS NULL

UNION

SELECT `tree`.`id`, 'Inner' AS `Type`
FROM `tree` INNER JOIN `tree` AS `children` ON `tree`.`id` = `children`.`p_id`
WHERE `tree`.`p_id` IS NOT NULL

UNION

SELECT `tree`.`id`, 'Leaf' AS `Type`
FROM `tree` LEFT JOIN `tree` AS `children` ON `tree`.`id` = `children`.`p_id`
WHERE `tree`.`p_id` IS NOT NULL AND `children`.`p_id` IS NULL