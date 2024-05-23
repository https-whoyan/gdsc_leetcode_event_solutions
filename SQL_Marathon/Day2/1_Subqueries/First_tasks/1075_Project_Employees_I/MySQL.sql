SELECT
    tb1.project_id,
    ROUND(
        AVG(tb2.experience_years)
    , 2) as average_years
FROM Project tb1
LEFT JOIN Employee tb2
ON tb1.employee_id = tb2.employee_id
GROUP BY tb1.project_id