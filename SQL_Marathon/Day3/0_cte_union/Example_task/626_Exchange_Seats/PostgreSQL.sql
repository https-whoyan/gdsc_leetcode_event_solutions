with mxSeat as (
    SELECT MAX(id) as mxValue
    FROM Seat 
)

SELECT *
FROM (
    SELECT 
        CASE 
            WHEN tb1.id % 2 = 1 THEN
                CASE
                    WHEN tb1.id = mxSeat.mxValue THEN tb1.id
                    ELSE tb1.id + 1
                END
            ELSE tb1.id - 1
        END as id,
        student 
    FROM Seat tb1, mxSeat 
) tmp
ORDER BY id