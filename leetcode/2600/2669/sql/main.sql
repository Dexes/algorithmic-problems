SELECT `artist`, COUNT(`id`) `occurrences`
FROM `spotify`
GROUP BY `artist`
ORDER BY 2, 1