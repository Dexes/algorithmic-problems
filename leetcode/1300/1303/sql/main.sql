SELECT `employee_id`, COUNT(`team_id`) OVER (PARTITION BY `team_id`) `team_size`
FROM `employee`