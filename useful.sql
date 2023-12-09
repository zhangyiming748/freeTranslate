select src,dst from history order by create_time desc ;
DELETE FROM history WHERE history.dst = "";
DELETE FROM history WHERE dst LIKE '%%';