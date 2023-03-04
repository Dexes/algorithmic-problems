SELECT `ad_id`,
       IFNULL(ROUND(SUM(`action` = 'Clicked') / (sum(`action` = 'Clicked') + sum(`action` = 'Viewed')) * 100, 2), 0) as `ctr`
FROM `ads`
GROUP BY `ad_id`
ORDER BY `ctr` DESC, `ad_id` ASC