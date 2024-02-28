select src,dst from histories order by created_at desc ;
DELETE FROM histories WHERE dst LIKE '%%';
DELETE FROM histories WHERE dst LIKE '';
select histories.src,histories.dst from histories order by id desc;