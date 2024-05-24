with tbl as (
    SELECT 
        visited_on, 
        SUM(amount) as day_sum 
    FROM Customer
    GROUP BY visited_on
)

SELECT 
    tb1.visited_on as visited_on, 
    SUM(tb2.day_sum) as amount,
    ROUND(AVG(tb2.day_sum * 1.0) * 1.0, 2) as average_amount
FROM
    tbl tb1, 
    tbl tb2
WHERE DATEDIFF(tb1.visited_on, tb2.visited_on) BETWEEN 0 AND 6
GROUP BY tb1.visited_on
HAVING COUNT(tb2.visited_on) = 7
ORDER BY tb1.visited_on 