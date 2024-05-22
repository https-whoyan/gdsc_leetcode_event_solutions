SELECT 
    sell_date, 
    COUNT(DISTINCT product) as num_sold,
    string_agg(
        DISTINCT product, ','
    ) as products
FROM Activities
GROUP BY sell_date
ORDER BY sell_date