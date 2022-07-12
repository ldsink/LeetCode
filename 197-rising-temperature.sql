-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | recordDate    | date    |
-- | temperature   | int     |
-- +---------------+---------+

select a.id
from Weather a
         left join Weather b on DATE_SUB(a.recordDate, INTERVAL 1 DAY) = b.recordDate
where a.temperature > b.temperature