select id,src,dst from history order by id desc ;
select src from history where dst = "";
delete from history where dst = "";