SELECT 
    ROUND(
        COUNT(DISTINCT player_id) / 
        (
            SELECT 
                COUNT(DISTINCT player_id) 
            FROM Activity
        )
    , 2) as fraction
FROM Activity
WHERE (player_id, DATE_SUB(event_date, INTERVAL 1 DAY)) IN (
    SELECT 
        player_id,
        MIN(event_date) as first_login 
    FROM Activity
    GROUP BY player_id
)