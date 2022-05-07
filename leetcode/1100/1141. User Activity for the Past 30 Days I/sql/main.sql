SELECT `activity_date` AS `day`, COUNT(DISTINCT `user_id`) AS `active_users`
FROM `activity`
WHERE SUBDATE(`activity_date`, 30) >= '2019-05-29'
GROUP BY `activity_date`