SELECT `buyer_id`
FROM `sales`
WHERE `buyer_id` IN (SELECT `buyer_id` FROM `sales` JOIN `product` USING(`product_id`) WHERE `product_name` = 'S8')
  AND `buyer_id` NOT IN (SELECT `buyer_id` FROM `sales` JOIN `product` USING(`product_id`) WHERE `product_name` = 'iPhone')
GROUP BY `buyer_id`