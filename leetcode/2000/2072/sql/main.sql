WITH `excellent` AS (
    SELECT
        (SELECT COUNT(*) FROM `NewYork` WHERE score > 89) `ny`,
        (SELECT COUNT(*) FROM `California` WHERE score > 89) `ca`
)

SELECT (
           CASE
               WHEN `excellent`.`ny` > `excellent`.`ca` THEN 'New York University'
               WHEN `excellent`.`ny` < `excellent`.`ca` THEN 'California University'
               ELSE 'No Winner'
               END
           ) `winner`
FROM `excellent`