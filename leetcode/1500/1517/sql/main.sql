SELECT `u`.`user_id`, `u`.`name`, `u`.`mail`
FROM `users` `u`
WHERE `u`.`mail` REGEXP '^[a-z][a-z0-9\\-_.]*@leetcode\\.com$'