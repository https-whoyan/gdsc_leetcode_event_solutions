SELECT (
    SELECT salary
    FROM (
        SELECT DISTINCT salary
        FROM Employee 
        ORDER BY salary DESC
    )
    LIMIT 1
    OFFSET 1
) as SecondHighestSalary