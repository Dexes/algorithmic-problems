SELECT COALESCE(ROUND(SUM(`sessions`) / COUNT(`user_id`), 2), 0) AS `average_sessions_per_user`
FROM (SELECT `user_id`, COUNT(DISTINCT `session_id`) AS `sessions`
      FROM `activity`
      WHERE `activity_date` BETWEEN '2019-06-28' AND '2019-07-27'
      GROUP BY `user_id`) AS `x`