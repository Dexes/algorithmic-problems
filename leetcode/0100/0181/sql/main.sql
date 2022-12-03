SELECT `employee`.`name` AS `Employee`
FROM `Employee` `employee`
         LEFT JOIN `Employee` `manager` ON `employee`.`managerId` = `manager`.`id`
WHERE `employee`.`salary` > `manager`.`salary`