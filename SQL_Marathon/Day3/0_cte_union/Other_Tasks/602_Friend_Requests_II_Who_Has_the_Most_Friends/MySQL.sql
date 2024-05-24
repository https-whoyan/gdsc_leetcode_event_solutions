with idTable as (
    SELECT DISTINCT requester_id as id
    FROM RequestAccepted
    UNION 
    SELECT DISTINCT accepter_id as id
    FROM RequestAccepted
)

SELECT
    idTable.id,
    COUNT(DISTINCT tb1.accepter_id) + COUNT(DISTINCT tb2.requester_id) as num
FROM idTable
LEFT JOIN RequestAccepted tb1
ON idTable.id = tb1.requester_id
LEFT JOIN RequestAccepted tb2
ON idTable.id = tb2.accepter_id
GROUP BY idTable.id
ORDER BY num DESC
LIMIT 1