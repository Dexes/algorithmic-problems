SELECT `project_id`
FROM `project`
GROUP BY `project_id`
HAVING COUNT(`project_id`) = (SELECT MAX(`count`) FROM (SELECT COUNT(`project_id`) AS `count` FROM `project` GROUP BY `project_id`) AS `x`)