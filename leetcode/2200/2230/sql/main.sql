CREATE PROCEDURE getUserIDs(startDate DATE, endDate DATE, minAmount INT)
BEGIN
    SELECT DISTINCT `user_id`
    FROM `purchases`
    WHERE `amount` >= minAmount AND `time_stamp` BETWEEN startDate AND endDate
    ORDER BY `user_id`;
END