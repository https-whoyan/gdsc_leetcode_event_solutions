SELECT 
    tb1.student_id, 
    tb1.student_name, 
    tb2.subject_name, 
    COUNT(tb3.student_id) AS attended_exams
FROM Students tb1
CROSS JOIN Subjects tb2
LEFT JOIN Examinations tb3
ON 
    tb1.student_id = tb3.student_id 
    AND
    tb3.subject_name = tb2.subject_name
GROUP BY 
    tb1.student_id,
    tb1.student_name, 
    tb2.subject_name
ORDER BY 
    tb1.student_id,
    tb2.subject_name