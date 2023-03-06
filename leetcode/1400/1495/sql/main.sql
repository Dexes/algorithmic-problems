SELECT `c`.`title`
FROM `content` `c` JOIN `tvProgram` `t` USING (`content_id`)
WHERE `c`.`kids_content` = 'Y' AND `c`.`content_type` = 'Movies' AND `t`.`program_date` BETWEEN '2020-06-01' AND '2020-06-30'
GROUP BY `c`.`content_id`