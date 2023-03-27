SELECT `a`.`machine_id`, ROUND(AVG(`b`.`timestamp` - `a`.`timestamp`), 3) AS `processing_time`
FROM `activity` `a` JOIN `activity` `b` USING(`machine_id`, `process_id`)
WHERE `a`.`activity_type` = 'start' AND `b`.`activity_type` = 'end'
GROUP BY `a`.`machine_id`