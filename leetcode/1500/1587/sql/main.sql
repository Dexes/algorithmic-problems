SELECT `name`, SUM(`amount`) AS `balance`
FROM `users` INNER JOIN `transactions` ON `users`.`account` = `transactions`.`account`
GROUP BY `users`.`account`
HAVING SUM(`amount`) > 10000