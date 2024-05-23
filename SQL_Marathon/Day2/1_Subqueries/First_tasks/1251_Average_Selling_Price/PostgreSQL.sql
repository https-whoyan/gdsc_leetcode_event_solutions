SELECT 
    tb1.product_id, 
    COALESCE(
        ROUND(
            SUM(tb2.units * tb1.price) * 1.0 / SUM(tb2.units), 
        2), 
    0) as average_price
FROM Prices tb1
LEFT JOIN UnitsSold tb2
ON tb2.product_id = tb1.product_id
AND tb2.purchase_date BETWEEN tb1.start_date 
AND tb1.end_date
GROUP BY tb1.product_id