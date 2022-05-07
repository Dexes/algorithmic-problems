WITH `secondHighest` AS (
    SELECT DISTINCT `salary`
    FROM `employee`
    ORDER BY `salary` DESC
    LIMIT 1 OFFSET 1
)

SELECT IFNULL((SELECT `salary` FROM `secondHighest`), NULL) AS `SecondHighestSalary`