with sumByTurn as (
    SELECT
        person_id,
        person_name,
        SUM(weight) OVER (
            ORDER BY turn
        ) as dpSum
    FROM Queue
    ORDER BY turn DESC
)

SELECT 
    person_name
FROM sumByTurn
WHERE dpSum <= 1000
LIMIT 1