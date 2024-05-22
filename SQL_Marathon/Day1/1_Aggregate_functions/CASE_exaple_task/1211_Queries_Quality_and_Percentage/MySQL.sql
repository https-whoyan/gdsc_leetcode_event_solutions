SELECT 
    query_name,
    ROUND(
        AVG(
            rating * 1.0 
            / position 
        ) * 1.0
    , 2) as quality,
    ROUND(
        SUM(
            CASE
                WHEN rating < 3 THEN 1 
                ELSE 0
            END
        ) * 1.0 
        / COUNT(*) * 100
    , 2) as poor_query_percentage 
FROM Queries
WHERE query_name IS NOT NULL
GROUP BY query_name