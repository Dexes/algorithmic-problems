CREATE FUNCTION getUserIDs(startDate DATE, endDate DATE, minAmount INT) RETURNS INT
BEGIN
    RETURN (
        SELECT COUNT(DISTINCT `user_id`) `user_cnt`
        FROM `purchases`
        WHERE `amount` >= minAmount AND `time_stamp` BETWEEN startDate AND endDate
    );
END