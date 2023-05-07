SELECT DISTINCT `emp_id`, `firstname`, `lastname`, MAX(`salary`) OVER (PARTITION BY `emp_id`) `salary`, `department_id`
FROM `salary`