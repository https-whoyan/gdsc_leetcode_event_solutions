<style>
.square1 {
    margin: 0 auto;
    height: 20px;
    width: 20px;
    background-color: #D0FFE4FF;
}
.square2 {
    margin: 0 auto;
    height: 20px;
    width: 20px;
    background-color: #ABA9CCFF;
}
.square3 {
    margin: 0 auto;
    height: 20px;
    width: 20px;
    background-color: #A17980FF;
}

.PostgreSQLText {
    color: #d0ffe4;
    display: inline-block;
    font-weight: bold;
}
.MySQLText {
    color: #A17980FF;
    display: inline-block;
    font-weight: bold;
}
.SameDBText {
    color: #aba9cc;
    display: inline-block;
    font-weight: bold;
}
.AdditionallyInfoText {
    color:#f8c6b6;
    display: inline-block;
    font-weight: bold;
}
.LeetCodeText {
    align-self: center;
    alignment-baseline: center;
    text-align-all: center;
    text-align: center;
    align-content: center;
    align-items: center;
    font-size: 18px;
    white-space: nowrap;
}
.Easy {
    color: #81ff6a;
    white-space: nowrap;
}
.Medium {
    color: #faff6a;
    white-space: nowrap;
}
.Hard {
    color: #ff6a6a;
    white-space: nowrap;
}
.Solution {
    color: #efefef;
    display: inline-block;
    white-space: nowrap;
}
</style>

<h2 align="center">
Welcome to SQL LeetCode Marathon!
<img src="https://res.cloudinary.com/startup-grind/image/upload/c_fill,w_500,h_500,g_center/c_fill,dpr_2.0,f_auto,g_center,q_auto:good/v1/gcs/platform-data-dsc/events/2023-10-10_13-10-44_NLy2I3k.png" 
width="22px" alt="GDSC AITU">
</h2>

<hr>
<h3 align="center">
Below is information regarding the <u>1st</u> day of the marathon 
</h3>
<hr>


<table style="border: none; border-collapse: collapse; padding-left: 10%; padding-right: 10%">
    <tr style="border: none;">
        <td width="26%" style="border: none; text-align: center; vertical-align: center; align-items: center; justify-content: center; align-content: center">
            <div class="square1"></div>
            <span>PostgreSQL</span>
        </td>
        <td width="26%" style="border: none; text-align: center; vertical-align: middle; align-items: center; justify-content: center;">
            <div class="square2"></div>
            <span>Same</span>
        </td>
        <td width="26%" style="border: none; text-align: center; vertical-align: middle; align-items: center; justify-content: center;">
            <div class="square3"></div>
            <span>MySQL</span>
        </td>
    </tr>
</table>


<hr> <b>
Today on the marathon you will learn: <br>
1) Basic SQL </b> <i>(Select, From, Where, Order By) </i> <br>
<b> 2) Aggregate functions </b> <i> (AVG, SUM, MAX, MIN, etc) </i>
<br>
<b> 3) What is </b>
<div class="PostgreSQLText" >COALESCE</div>(<div class="MySQLText"> <i>IFNULL</i> </div>),
how to 
<div class="SameDBText"> GROUP BY</div>,
<div class="SameDBText"> CASE</div>,
<div class="SameDBText"> ROUND</div>,
<div class="PostgreSQLText">TO_CHAR</div>(<div class="MySQLText">DATE_FORMAT</div>),
<br>
<b> 4) How to work to string type in SQL:</b>
<div class="SameDBText"> CONCAT</div>,
<div class="SameDBText"> UPPER/LOWER</div>,
<div class="SameDBText"> LIKE/ILIKE</div>,
<div class="SameDBText"> SUBSTR</div>, <b> and also </b> 
<div class="PostgreSQLText">STRING_AGG</div>(<div class="MySQLText">GROUP_CONCAT</div>),
<div class="PostgreSQLText">REGEXP_LIKE</div>(<div class="MySQLText">REGEXP</div>).

<hr>
<h4 align="center">
<ul>
    <li> Solve 15 tasks </li>
    <li> Marathon time: 2 hours, 40 minutes.</li>
</ul>
</h4>

<hr>
<hr>
<table style="border: none; border-collapse: collapse; margin-left: 10%">
    <tr style="border: none;">
        <td width="40%" style="border: none; text-align: center; vertical-align: center; align-items: center; justify-content: center; align-content: center">
            |>
            <span>Sample tasks</span>
        </td>
        <td width="50%" style="border: none; text-align: center; vertical-align: middle; align-items: center; justify-content: center;">
            |
            <span>Tasks to solve.</span>
        </td>
    </tr>
</table>
<hr>
<table style="border: none; border-collapse: collapse; padding-left: 10%">
    <tr style="border: none;">
        <td width="16%" style="border: none; text-align: center; vertical-align: center; align-items: center; justify-content: center; align-content: center">
            <text class="Easy">Easy tasks</text>
        </td>
        <td width="26%" style="border: none; text-align: center; vertical-align: middle; align-items: center; justify-content: center;">
            <text class="Medium">Medium tasks</text>
        </td>
        <td width="26%" style="border: none; text-align: center; vertical-align: middle; align-items: center; justify-content: center;">
            <text class="Hard">Hard tasks</text>
        </td>
    </tr>
</table>
<hr>

<h3 align="center"> <u>
Basic SQL </u>
</h3>
<hr>
<div align="center" class="AdditionallyInfoText">
    Additionally material: <a href="https://shultais.education/blog/sql-for-beginners/first-sql-queries"> click</a>, <a href="http://2sql.ru/novosti/sql-order-by/">click.</a>
</div>

<hr>
<div class="LeetCodeText">
    <text style="color: #ffda40">Leet</text><text style="color: #626262">Code</text>
    Tasks
</div>

> |>  <text class="Easy"> 1757</text>. Recyclable and Low Fat Products. <a href="https://leetcode.com/problems/recyclable-and-low-fat-products/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/0_Basic_SQL/Exaple_task/1757_Recyclable_and_Low_Fat_Products"> click.</a> <br>

> | <text class="Easy"> 584. </text>Find Customer Referee<a href="https://leetcode.com/problems/find-customer-referee/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/0_Basic_SQL/Other_tasks/584_Find_Customer_Referee"> click.</a> <br>

> | <text class="Easy"> 595. </text>Big Countries<a href="https://leetcode.com/problems/big-countries/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/0_Basic_SQL/Other_tasks/595_Big_Countries"> click.</a> <br>

> | <text class="Easy"> 1148. </text>Article Views I<a href="https://leetcode.com/problems/article-views-i/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/0_Basic_SQL/Other_tasks/1148_Article_Views_I"> click.</a> <br>

> | <text class="Easy"> 1683. </text>Invalid Tweets<a href="https://leetcode.com/problems/invalid-tweets/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/0_Basic_SQL/Other_tasks/1683_Invalid_Tweets"> click.</a> <br>

<hr>
<h3 align="center"> <u>
Aggregate functions </u>
</h3> 
<hr>
<div align="center" class="AdditionallyInfoText">
    Additionally material: <a href="https://sky.pro/media/agregatnye-funkczii-v-sql-sut-ponyatiya-i-primery/"> click</a>
</div>

<hr>
<div class="LeetCodeText">
    <text style="color: #ffda40">Leet</text><text style="color: #626262">Code</text>
    Tasks
</div>

> | <text class="Easy"> 620.</text> Not Boring Movies. <a href="https://leetcode.com/problems/not-boring-movies/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/1_Aggregate_functions/First_tasks/620_Not_Boring_Movies"> click.</a> <br>

> | <text class="Easy"> 1633.</text> Percentage of Users Attended a Contest <a href="https://leetcode.com/problems/percentage-of-users-attended-a-contest/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/1_Aggregate_functions/First_tasks/1633_Percentage_of_Users_Attended_a_Contest"> click.</a> <br>

<div align="center" class="AdditionallyInfoText">
    Additionally material, 
    <text class="SameDBText">CASE</text>( 
    <text class="MySQLText">IF</text> ):
    <a href="https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-case/"> click</a>,
    <a href="https://www.w3schools.com/MYSQL/func_mysql_if.asp"> click</a>
</div> <br>

> |> <text class="Easy"> 1633.</text> Queries Quality and Percentage <a href="https://leetcode.com/problems/queries-quality-and-percentage/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/1_Aggregate_functions/CASE_exaple_task/1211_Queries_Quality_and_Percentage"> click.</a> <br>

> | <text class="Easy"> 1141.</text> User Activity for the Past 30 Days I <a href="https://leetcode.com/problems/user-activity-for-the-past-30-days-i/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/1_Aggregate_functions/Second_tasks/1141_User_Activity_for_the_Past_30_Days_I"> click.</a> <br>

<div class="AdditionallyInfoText"> Additionally material, 
    <text class="PostgreSQLText">TO_CHAR</text>( 
    <text class="SameDBText">DATE_FORMAT</text> ):
    <a href="https://code.mu/ru/sql/manual/date_format/"> click</a>,
    <a href="https://www.postgresqltutorial.com/postgresql-string-functions/postgresql-to_char/"> click</a>
</div> <br>

> | <text class="Medium"> 1193. </text> Monthly Transactions I <a href="https://leetcode.com/problems/monthly-transactions-i/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/1_Aggregate_functions/Second_tasks/1193_Monthly_Transactions_I"> click.</a> <br>

> | <text class="Easy"> 2356. </text> Number of Unique Subjects Taught by Each Teacher <a href="https://leetcode.com/problems/find-followers-count/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/1_Aggregate_functions/Second_tasks/2356_Number_of_Unique_Subjects_Taught_by_Each_Teacher"> click.</a> <br>

<hr>
<h3 align="center"> <u>
String functions </u>
</h3> 

<hr>
<div align="center" class="AdditionallyInfoText"> Additionally material, 
    <text class="SameDBText">LIKE</text>( 
    <text class="PostgreSQLText">ILIKE</text> ):
    <a href="https://www.w3schools.com/postgresql/postgresql_like.php"> click</a>
</div> <br>

> | <text class="Easy"> 1527. </text> Patients With a Condition <a href="https://leetcode.com/problems/patients-with-a-condition/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/2_string_functions/concat_upper_lower_like_substr/1527_Patients_With_a_Condition"> click.</a> <br>
> 
<div class="AdditionallyInfoText"> Additionally material, 
    <text class="SameDBText">UPPER</text>,  
    <text class="SameDBText">LOWER</text>, 
    <text class="PostgreSQLText">INITCAP</text>:
    <a href="https://commandprompt.com/education/postgresql-letter-case-functions-with-practical-examples/#:~:text=PostgreSQL%20provides%20various%20functions%2C%20such,uppercase%2C%20lowercase%2C%20proper%20case."> click</a>
</div> <br>

<div class="AdditionallyInfoText"> 
    Additionally material, 
    <text class="SameDBText">CONCAT</text>,
    <a href="https://www.geeksforgeeks.org/concat-function-in-mysql/"> click</a>
</div> <br>
<div class="AdditionallyInfoText"> Additionally material, 
    <text class="SameDBText">SUBSTR</text>,
    <a href="https://oracleplsql.ru/substr-function.html"> click</a>
</div> <br>

> | <text class="Easy"> 1667. </text> Fix Names in a Table <a href="https://leetcode.com/problems/fix-names-in-a-table/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/2_string_functions/concat_upper_lower_like_substr/1667_Fix_Names_in_a_Table"> click.</a> <br>

<div class="AdditionallyInfoText"> Additionally material, 
    <text class="PostgreSQLText">STRING_AGG</text>,  
    <text class="MySQLText">GROUP_CONCAT</text>,
    <a href="https://code.mu/ru/sql/manual/group_concat/"> click</a>
    <a href="https://popsql.com/learn-sql/postgresql/how-to-use-stringagg-in-postgresql"> click</a>
</div> <br>

> | <text class="Easy"> 1484. </text> Group Sold Products By The Date <a href="https://leetcode.com/problems/group-sold-products-by-the-date/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/2_string_functions/group_concat/1484_Group_Sold_Products_By_The_Date"> click.</a> <br>

<div class="AdditionallyInfoText"> Additionally material, 
    <text class="Solution">What is Regexp?</text>,
    <a href="https://tproger.ru/articles/regexp-for-beginners"> click</a>,
    <a href="https://habr.com/ru/articles/567106/"> click</a>
</div> <br>

<div class="AdditionallyInfoText"> Additionally material, 
    <text class="MySQLText">REGEXP</text>,
    <text class="PostgreSQLText">REGEXP_LIKE</text>,
    <a href="https://a-billing.com/docs/avantys_billing/7/abilling/mysql_regexp.html"> click</a>,
    <a href="https://pgpedia.info/r/regexp_like.html"> click</a>
</div> <br>

> | <text class="Easy"> 1484. </text> Group Sold Products By The Date <a href="https://leetcode.com/problems/find-users-with-valid-e-mails/description/?envType=study-plan-v2&envId=top-sql-50"> click. </a>
> <text class="Solution">Solution: </text> <a href="https://github.com/https-whoyan/gdsc_leetcode_event_solutions/tree/main//SQL_Marathon/Day1/2_string_functions/regexp/1517_Find_Users_With_Valid_E-Mails"> click.</a> <br>

<br>
<hr>
<h2 align="center">
Thank you for attention! 🖤
<img src="https://res.cloudinary.com/startup-grind/image/upload/c_fill,w_500,h_500,g_center/c_fill,dpr_2.0,f_auto,g_center,q_auto:good/v1/gcs/platform-data-dsc/events/2023-10-10_13-10-44_NLy2I3k.png" 
width="22px" alt="GDSC AITU">
</h2>