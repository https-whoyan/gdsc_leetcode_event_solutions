SELECT
    tb2.product_name,
    SUM(tb1.unit) as unit
FROM Orders tb1
JOIN Products tb2
ON tb1.product_id = tb2.product_id
WHERE tb1.order_date BETWEEN "2020-02-01" AND "2020-02-29"
GROUP BY tb1.product_id
HAVING SUM(tb1.unit) >= 100