SELECT `q`.`id`, `q`.`year`, IFNULL(`n`.`npv`, 0) `npv`
FROM `queries` `q` LEFT JOIN `npv` `n` USING (`id`, `year`)