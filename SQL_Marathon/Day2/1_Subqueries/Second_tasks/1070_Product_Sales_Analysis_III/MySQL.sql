SELECT 
    tb1.product_id,
    tb2.min_year as first_year,
    tb1.quantity, 
    tb1.price
FROM Sales tb1
JOIN (
    SELECT 
        product_id,
        MIN(year) as min_year
    FROM Sales
    GROUP BY product_id
) tb2
ON tb1.product_id = tb2.product_id 
WHERE tb1.year = tb2.min_year