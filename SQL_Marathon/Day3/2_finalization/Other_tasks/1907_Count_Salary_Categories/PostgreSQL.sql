SELECT 
    'Low Salary' as category, 
    COUNT(*) as accounts_count 
FROM (
    SELECT account_id 
    FROM Accounts 
    WHERE income < 20000
) lowTbl
UNION
SELECT 
    'Average Salary' as category, 
    COUNT(*) as accounts_count 
FROM (
    SELECT account_id 
    FROM Accounts 
    WHERE income >= 20000
    AND income <= 50000
) avgTbl
UNION
SELECT 
    'High Salary' as category, 
    COUNT(*) as accounts_count
FROM (
    SELECT account_id 
    FROM Accounts 
    WHERE income > 50000
) highTbl