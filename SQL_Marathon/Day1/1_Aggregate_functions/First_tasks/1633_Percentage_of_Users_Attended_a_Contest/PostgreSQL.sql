SELECT 
    contest_id, 
    ROUND(
        COUNT(distinct user_id) * 100.0 /
        (
            SELECT COUNT(user_id) 
            FROM users
        ), 2
    ) as percentage
FROM Register
GROUP BY contest_id
ORDER BY percentage DESC, contest_id