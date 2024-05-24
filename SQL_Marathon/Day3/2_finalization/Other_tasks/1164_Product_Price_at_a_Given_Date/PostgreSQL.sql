SELECT
    DISTINCT tmp1.product_id,
    COALESCE(tmp2.new_price, 10) as price
FROM Products tmp1
LEFT JOIN (
    SELECT product_id, new_price
    FROM Products tb1
    WHERE change_date = (
        SELECT MAX(change_date) 
        FROM Products tb2
        WHERE change_date <= '2019-08-16'
        AND tb2.product_id = tb1.product_id
    )
) tmp2
ON tmp1.product_id = tmp2.product_id