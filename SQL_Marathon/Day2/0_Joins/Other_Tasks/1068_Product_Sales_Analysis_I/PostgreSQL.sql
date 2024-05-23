SELECT 
    tb2.product_name, 
    tb1.year, 
    tb1.price
FROM Sales tb1
JOIN Product tb2
ON tb1.product_id = tb2.product_id