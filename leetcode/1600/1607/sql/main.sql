WITH `orders2020` AS (
    SELECT DISTINCT `o`.`seller_id`
    FROM `orders` `o`
    WHERE YEAR(`o`.`sale_date`) = 2020
)

SELECT `s`.`seller_name`
FROM `seller` `s` LEFT JOIN `orders2020` `o` USING(`seller_id`)
WHERE `o`.`seller_id` IS NULL
ORDER BY `s`.`seller_name`