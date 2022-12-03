SELECT `First`.`id`
FROM `Weather` AS `First`,
     `Weather` AS `Second`
WHERE `First`.`temperature` > `Second`.`temperature`
  AND `First`.`recordDate` = `Second`.`recordDate` + INTERVAL 1 DAY;