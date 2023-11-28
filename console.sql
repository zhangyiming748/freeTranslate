select * from avif join task on task.id = avif.task_id;
select * from task order by create_time desc ;
select * from avif;
drop table av,custom,err,file,image,save,telegraph,text,video,audio,avif,task;
drop table video;
show tables;
drop table file;

select * from video order by id desc ;
select * from task order by id desc ;
select * from err order by create_time desc ;