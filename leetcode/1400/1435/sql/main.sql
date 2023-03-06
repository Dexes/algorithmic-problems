WITH `conditions` AS (
    SELECT 1 `id`, '[0-5>' `bin`, 0 `min`, 299 `max`
    UNION ALL
    SELECT 2 `id`, '[5-10>' `bin`, 300 `min`, 599 `max`
    UNION ALL
    SELECT 3 `id`, '[10-15>' `bin`, 600 `min`, 899 `max`
    UNION ALL
    SELECT 4 `id`, '15 or more' `bin`, 900 `min`, 1000000 `max`
)

SELECT `c`.`bin`, COUNT(`s`.`session_id`) `total`
FROM `conditions` `c` LEFT JOIN `sessions` `s` ON `s`.`duration` BETWEEN `c`.`min` AND `c`.`max`
GROUP BY `c`.`id`