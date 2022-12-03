WITH `KeepRows` AS (
    SELECT MIN(`id`) as `id`
    FROM `Person`
    GROUP BY `email`
)
DELETE
FROM `Person`
WHERE `id` NOT IN (SELECT `id` FROM `KeepRows`)
