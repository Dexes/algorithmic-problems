WITH
    MAX_MARKETING as ( SELECT MAX(`salary`) `max_marketing` FROM `salaries` WHERE `department` = 'Marketing' ),
    MAX_ENGINEERING as ( SELECT MAX(`salary`) `max_engineering` FROM `salaries` WHERE `department` = 'Engineering' )

SELECT ABS(`max_marketing` - `max_engineering`) `salary_difference`
FROM MAX_MARKETING, MAX_ENGINEERING