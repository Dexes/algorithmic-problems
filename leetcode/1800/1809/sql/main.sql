SELECT `p`.`session_id`
FROM `playback` `p` LEFT JOIN `ads` `a`
    ON `p`.`customer_id` = `a`.`customer_id` AND `a`.`timestamp` BETWEEN `p`.`start_time` AND `p`.`end_time`
WHERE `a`.`ad_id` IS NULL
GROUP BY `p`.`session_id`