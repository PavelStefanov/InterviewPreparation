-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | employee_id   | int     |
-- | team_id       | int     |
-- +---------------+---------+
-- employee_id is the primary key for this table.
-- Each row of this table contains the ID of each employee and their respective team.

-- Employee Table:
-- +-------------+------------+
-- | employee_id | team_id    |
-- +-------------+------------+
-- |     1       |     8      |
-- |     2       |     8      |
-- |     3       |     8      |
-- |     4       |     7      |
-- |     5       |     9      |
-- |     6       |     9      |
-- +-------------+------------+
-- Result table:
-- +-------------+------------+
-- | employee_id | team_size  |
-- +-------------+------------+
-- |     1       |     3      |
-- |     2       |     3      |
-- |     3       |     3      |
-- |     4       |     1      |
-- |     5       |     2      |
-- |     6       |     2      |
-- +-------------+------------+
-- Employees with Id 1,2,3 are part of a team with team_id = 8.
-- Employees with Id 4 is part of a team with team_id = 7.
-- Employees with Id 5,6 are part of a team with team_id = 9.


# Best solution is using window function 
SELECT employee_id, 
		COUNT(employee_id) OVER (PARTITION BY team_id) AS team_size
FROM Employee
;


# Calculate in the Select statement
SELECT employee_id, 
		(SELECT COUNT(1)
		FROM Employee e2 WHERE e2.employee_id = e.employee_id) AS team_size 
FROM Employee e
;


# Use SELF-JOIN; Don't forgot to use GROUP BY 
# Why need 2 tables - It's basically used where there is any relationship between rows stored in the same table (correlate pairs of rows from the same table)
select e1.employee_id,count(*) team_size
		from employee e1 left join employee e2on e1.team_id = e2.team_id
group by e1.employee_id;