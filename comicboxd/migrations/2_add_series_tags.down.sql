
DROP VIEW IF EXISTS series;

create view series as
select 
    s.*, 
    user_series.list 
from 
    (
        SELECT 
            series as name, 
            count(series) as total, 
            sum(read) as read, 
            user_id 
        FROM 
            "book_user_book" group by series, user_id
    ) as s left join user_series on name=series;