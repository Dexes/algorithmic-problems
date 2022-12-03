SELECT `actor_id`, `director_id`
FROM `actorDirector`
GROUP BY `actor_id`, `director_id`
HAVING COUNT(`actor_id`) >= 3