WITH tmp AS (
    SELECT
        num,
        id,
        LAG(id) OVER(
            ORDER BY id
        ) previd,
        LAG(num) OVER(
            ORDER BY id
        ) prevnum,
        LEAD(id) OVER(
            ORDER BY id
        ) nextid,
        LEAD(num) OVER(
            ORDER BY id
        ) nextnum
    FROM logs
)

SELECT
    DISTINCT num as "ConsecutiveNums" 
FROM tmp
WHERE
    # ids
    previd + 1 = id AND id + 1 = nextid
    AND
    # nums
    prevnum = num AND num = nextnum