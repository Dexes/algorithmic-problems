SELECT `contest_id`, ROUND(COUNT(`user_id`) / (SELECT count(`user_id`) FROM `users`) * 100, 2) `percentage`
FROM `register` `r`
GROUP BY `contest_id`
ORDER BY `percentage` DESC, `contest_id`