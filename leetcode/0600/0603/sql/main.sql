SELECT DISTINCT `a`.`seat_id`
FROM `cinema` `a` INNER JOIN `cinema` `b` ON ABS(`a`.`seat_id` - `b`.`seat_id`) = 1
WHERE `a`.`free` AND `b`.`free`
ORDER BY `a`.`seat_id`