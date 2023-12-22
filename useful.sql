select src,dst from history order by create_time desc ;
DELETE FROM history WHERE dst LIKE '%%';
DELETE FROM history WHERE dst LIKE '';
select histories.src,histories.dst from histories order by id desc;