WITH `times` AS (
    SELECT `user_id`, `time_stamp` `current`, LEAD(`time_stamp`) OVER (PARTITION BY `user_id` ORDER BY `time_stamp`) `next`
    FROM `confirmations`
)

SELECT DISTINCT `user_id`
FROM `times`
WHERE TIMESTAMPDIFF(SECOND, `current`, `next`) <= 86400