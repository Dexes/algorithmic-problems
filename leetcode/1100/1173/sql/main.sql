SELECT ROUND(
   (SELECT COUNT(`delivery_id`) FROM `delivery` WHERE `order_date` = `customer_pref_delivery_date`)
   / (SELECT COUNT(`delivery_id`) FROM `delivery`)
   * 100,
   2
) AS `immediate_percentage`