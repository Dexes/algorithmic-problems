SELECT TRIM(LOWER(`s`.`product_name`)) `product_name`, DATE_FORMAT(`s`.`sale_date`, '%Y-%m') `sale_date`, COUNT(`s`.`sale_id`) `total`
FROM `sales` `s`
GROUP BY 1, 2
ORDER BY 1, 2