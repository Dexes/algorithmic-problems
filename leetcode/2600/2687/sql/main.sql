SELECT `bike_number`, MAX(`end_time`) `end_time`
FROM `bikes`
GROUP BY `bike_number`