SELECT `salesPerson`.`name`
FROM `salesPerson`
WHERE `salesPerson`.`sales_id` NOT IN (
    SELECT `orders`.`sales_id`
    FROM `company` INNER JOIN `orders` ON `company`.`com_id` = `orders`.`com_id`
    WHERE `name` = 'RED'
)