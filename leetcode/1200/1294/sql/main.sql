SELECT `c`.`country_name`, CASE WHEN AVG(`w`.`weather_state`) <= 15 THEN 'Cold' WHEN AVG(`w`.`weather_state`) >= 25 THEN 'Hot' ELSE 'Warm' END `weather_type`
FROM `countries` `c` JOIN `weather` `w` ON `c`.`country_id` = `w`.`country_id`
WHERE `w`.`day` BETWEEN '2019-11-01' AND '2019-11-30'
GROUP BY `c`.`country_id`