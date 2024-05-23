SELECT
    tb1.user_id,
    ROUND(
        AVG(
            IF(
                tb2.action="confirmed", 1, 0
            )
        )
    , 2) AS confirmation_rate
FROM Signups tb1
LEFT JOIN Confirmations tb2
ON tb2.user_id = tb1.user_id
GROUP BY user_id