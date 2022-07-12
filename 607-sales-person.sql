# Write your MySQL query statement below

select name
from SalesPerson
where not exists(
        select 1
        from (select distinct sales_id as id
              from Orders
              where exists(
                            select 1
                            from Company
                            where Company.name = 'RED'
                              and Company.com_id = Orders.com_id
                        )) t
        where t.id = sales_id
    )
