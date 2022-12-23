SELECT MIN(ABS(`a`.`x` - `b`.`x`)) AS `shortest`
FROM `point` `a` INNER JOIN `point` `b` ON `a`.`x` != `b`.`x`