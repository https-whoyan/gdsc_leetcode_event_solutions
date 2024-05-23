SELECT 
    ROUND(
        COUNT(tb2.games_played) * 1.0 / 
        COUNT(tb1.games_played) 
    , 2) as fraction  
FROM Activity tb1
LEFT JOIN Activity tb2 
ON tb1.player_id = tb2.player_id
AND tb2.event_date - tb1.event_date = 1
WHERE tb1.event_date IN (
    SELECT MIN(event_date)
    FROM Activity
    WHERE player_id = tb1.player_id 
)