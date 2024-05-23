SELECT tb1.class
FROM Courses tb1
JOIN (
    SELECT COUNT(DISTINCT student) as student_count, class
    FROM Courses
    GROUP BY class
) tb2
ON tb1.class = tb2.class
WHERE tb2.student_count >= 5
GROUP BY tb1.class