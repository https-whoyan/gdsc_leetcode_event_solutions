SELECT 
    ROUND(
        SUM(tiv_2016)::numeric * 1.0, 2
    ) as tiv_2016
FROM (
    SELECT tb1.tiv_2016
    FROM Insurance tb1
    JOIN Insurance tb2
    ON tb1.tiv_2015 = tb2.tiv_2015
    JOIN (
        SELECT lat, lon
        FROM Insurance
        GROUP BY lat, lon
        HAVING COUNT(lat) = 1 AND COUNT(lon) = 1
    ) tb3
    ON tb3.lat = tb1.lat
    AND tb3.lon = tb1.lon
    WHERE tb2.pid != tb1.pid
    GROUP BY tb1.pid, tb1.tiv_2016
) tmp