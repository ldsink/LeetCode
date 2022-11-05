select s.name   as Department,
       e.name   as Employee,
       e.salary as Salary
from Employee e
         left join (select departmentId, d.Name as name, min(salary) as salary
                    from (select departmentId, salary, RANK() OVER(PARTITION BY departmentId ORDER BY salary desc) AS ranked
                          from (select distinct departmentId, salary
                                from Employee) e) e
                             left join Department d on e.departmentId = d.id
                    where ranked < 4
                    group by departmentId) s on e.departmentId = s.departmentId
where e.salary >= s.salary
