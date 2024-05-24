SELECT * FROM (
    SELECT tb2.name as results
    FROM (
        SELECT user_id, COUNT(tb1.user_id) as tb1_count
        FROM MovieRating tb1
        GROUP BY user_id
    ) tb1
    JOIN Users tb2
    ON tb1.user_id = tb2.user_id
    ORDER BY tb1_count DESC, tb2.name
    LIMIT 1
) tmp1

UNION ALL

SELECT * FROM (
    SELECT tb4.title as results
    FROM (
        SELECT movie_id, AVG(rating) as tb3_avg_rating
        FROM MovieRating
        WHERE created_at <= "2020-02-29" 
        AND created_at >= "2020-02-01"
        GROUP BY movie_id
    ) tb3 
    JOIN Movies tb4
    ON tb3.movie_id = tb4.movie_id
    ORDER BY tb3_avg_rating DESC, tb4.title
    LIMIT 1
) tmp2