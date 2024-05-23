SELECT
    MAX(single_num) as num
FROM (
    SELECT num as single_num
    FROM MyNumbers 
    GROUP BY num
    HAVING COUNT(num) = 1
) tb1