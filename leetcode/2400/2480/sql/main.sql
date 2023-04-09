SELECT `m`.`symbol` `metal`, `n`.`symbol` `nonmetal`
FROM `elements` `m` JOIN `elements` `n` ON `m`.`type` = 'Metal' AND `n`.`type` = 'Nonmetal'