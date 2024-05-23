SELECT 
    tb1.user_id, 
    ROUND(
        avg(CASE
            WHEN tb2.action = 'confirmed' THEN 1.0
            ELSE 0
        END)
    ,2) AS confirmation_rate
FROM Signups tb1
LEFT JOIN Confirmations tb2
ON tb2.user_id = tb1.user_id
GROUP BY tb1.user_id