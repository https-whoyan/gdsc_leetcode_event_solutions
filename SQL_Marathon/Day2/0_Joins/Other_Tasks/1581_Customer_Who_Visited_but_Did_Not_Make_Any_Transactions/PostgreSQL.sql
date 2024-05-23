SELECT
    V.customer_id,
    COUNT(V.visit_id) as count_no_trans
FROM Visits V
LEFT JOIN Transactions T
ON V.visit_id=T.visit_id
WHERE T.transaction_id IS NULL
GROUP BY V.customer_id