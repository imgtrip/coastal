-- +migrate Up
insert into `comment`(`id`,`album_id`,`parent_id`,`user_id`,`content`,`created_at`,`updated_at`) values
('5','24','0','1','<img src="/static/image/1f600.png">','2017-09-18 23:18:32','2017-09-18 23:18:32'),
('6','19','0','1','<img src="/static/image/1f914.png">  哈哈','2017-09-18 23:19:02','2017-09-18 23:19:02'),
('7','15','0','7','丑','2017-09-21 13:42:28','2017-09-21 13:42:28'),
('8','15','0','7','我的头像好丑，怎麽样换头像','2017-09-21 13:43:04','2017-09-21 13:43:04'),
('9','19','0','1','头像是自动设置为你最后一张喜欢的图,所以如果想换头像的话,再喜欢一张图就好啦
','2017-09-21 15:50:44','2017-09-21 15:50:44'),
('10','27','0','31','我是第一个哦 <img src="/static/image/1f600.png">！欢迎来访','2017-10-05 16:37:54','2017-10-05 16:37:54'),
('11','16','0','31','马克！','2017-10-05 16:38:58','2017-10-05 16:38:58'),
('12','27','0','1','<img src="/static/image/1f600.png">   <a href=\'/user/show/31\'>@Arvin</a>','2017-10-06 13:50:18','2017-10-06 13:50:18'),
('13','16','0','1','好想去旅行 太美了 <img src="/static/image/1f60d.png">','2017-10-15 11:17:17','2017-10-15 11:17:17'),
('14','33','0','1','我在坦克世界里面好像开过桥上那个坦克... <img src="/static/image/1f923.png">','2017-10-15 11:24:01','2017-10-15 11:24:01'),
('15','15','0','1','<a href=\'/user/show/7\'>@jessie</a> 目前头像是自动设置为你喜欢的最后一张图 , 以后再做主动设置头像功能吧  <img src="/static/image/1f60e.png">','2017-10-17 09:58:49','2017-10-17 09:58:49'),
('16','19','0','44','萌！','2017-10-20 15:15:47','2017-10-20 15:15:47'),
('17','37','0','1','<img src="/static/image/1f600.png"> 风景很美~','2017-10-20 22:03:41','2017-10-20 22:03:41'),
('18','37','0','41','<img src="/static/image/1f60d.png"> 没想到作者竟然回复了，很感谢作者制作的这个网站，也很感谢这些摄影家采集的如此瑰丽的大自然。','2017-10-22 19:06:27','2017-10-22 19:06:27'),
('19','37','0','1','你的支持,是旅图网继续优化前进的动力 !   <a href=\'/user/show/41\'>@文仓</a> 我也很感谢你主动创建图集 , 让我知道有人喜欢  <img src="/static/image/1f600.png">','2017-10-22 21:09:00','2017-10-22 21:09:00'),
('20','37','0','41','<img src="/static/image/1f606.png">','2017-10-23 11:10:12','2017-10-23 11:10:12'),
('21','40','0','1','是要过程序员节吗... <img src="/static/image/1f914.png">','2017-10-24 19:59:20','2017-10-24 19:59:20'),
('22','37','0','70','这个世界，~~~~美`~~~`只有自己才能寻找','2017-11-04 20:46:39','2017-11-04 20:46:39'),
('23','43','0','70','希望带着光芒，如同破晓带着闪电，

而我带着春暖花开来到你身旁

明信片载着你一同飞向远方','2017-11-04 21:18:02','2017-11-04 21:18:02'),
('24','43','0','1','<img src="/static/image/1f600.png">','2017-11-04 22:28:55','2017-11-04 22:28:55'),
('25','40','0','21','<a href=\'/user/show/1\'>@qskane</a> 233','2017-11-16 09:38:44','2017-11-16 09:38:44'),
('26','50','0','132','个非官方个','2017-12-17 05:37:30','2017-12-17 05:37:30'),
('34','39','0','41','<p><img src="/static/image/1f601.png"></p>','2017-12-23 15:26:07','2017-12-23 15:26:07'),
('35','112','0','41','<p><img src="/static/image/1f60a.png">加油</p>','2017-12-30 10:52:36','2017-12-30 10:52:36'),
('36','112','0','158','<p>哈哈，能搜图吗</p>','2018-01-06 16:27:09','2018-01-06 16:27:09'),
('37','36','0','165','<p>为何？究竟是为何？好久没有更新了。</p>','2018-01-10 17:29:24','2018-01-10 17:29:24'),
('38','112','0','168','<p>加油</p>','2018-01-11 17:51:52','2018-01-11 17:51:52'),
('39','40','0','168','<p>1024就比较厉害了</p>','2018-01-11 17:56:04','2018-01-11 17:56:04'),
('40','112','0','1','<h4><img src="/static/image/1f602.png"></h4>','2018-01-19 09:22:47','2018-01-19 09:22:47'),
('41','112','0','196','<p>相信你辛勤的付出会结出丰硕的果实！加油！</p>','2018-01-22 00:08:06','2018-01-22 00:08:06'),
('42','112','0','214','<p>豪哥</p>','2018-01-24 23:54:45','2018-01-24 23:54:45'),
('43','112','0','233','<p>试试看，图不错</p>','2018-01-27 22:19:36','2018-01-27 22:19:36'),
('44','112','0','289','<p>怎样搜斯嘉丽约翰逊的图？？？</p>','2018-02-11 15:11:16','2018-02-11 15:11:16'),
('45','112','0','347','<p><img src="/static/image/1f618.png">大大辛苦，资瓷~</p>','2018-02-23 10:39:38','2018-02-23 10:39:38'),
('46','441','0','388','<p>今天认识你，旅图网<img src="/static/image/1f609.png"></p>','2018-02-27 21:48:00','2018-02-27 21:48:00'),
('47','112','0','395','<p>做的很不错，在各方面来讲都是行业里面做的很好的</p>','2018-02-28 17:56:22','2018-02-28 17:56:22'),
('48','112','0','395','<p>很不错，加油</p>','2018-02-28 17:56:35','2018-02-28 17:56:35'),
('49','65','0','23','<p><img src="/static/image/1f61b.png"></p>','2018-03-02 15:38:55','2018-03-02 15:38:55'),
('50','36','0','414','<p>好喜欢啊，谢谢</p>','2018-03-04 16:37:42','2018-03-04 16:37:42'),
('51','39','0','466','<p><img src="/static/image/1f604.png"></p>','2018-03-11 14:26:34','2018-03-11 14:26:34'),
('52','37','0','466','<p>好看</p>','2018-03-11 14:49:03','2018-03-11 14:49:03'),
('53','19','0','466','<p>好想法</p>','2018-03-11 14:55:41','2018-03-11 14:55:41'),
('54','227','0','1','<p>此图辑，风格另类，很有特点，不错不错<img src="/static/image/1f923.png"><img src="/static/image/1f923.png"><img src="/static/image/1f923.png"></p>','2018-03-12 20:33:10','2018-03-12 20:33:10'),
('55','112','0','529','<p>想问下,那个刷新后的图片是随即展示还是根据用户喜欢的类型推荐展示的?</p>','2018-03-17 14:34:56','2018-03-17 14:34:56'),
('56','596','0','529','<p>简单写实</p><p>大胆创意</p><p>真实女人</p><p>会心一笑</p><p>刹那暴虐</p><p>恬静如画</p><p>原始诱惑</p><p>奇景怪色</p>','2018-03-17 14:50:54','2018-03-17 14:50:54'),
('57','112','0','1','<h4>@<a href="https://www.imgtrip.com/u/529" target="_blank" style="color: rgb(24, 144, 255);">仗剑天涯</a> 目前首页是随机展示的，根据喜好推荐功能正在开发中 <img src="/static/image/1f923.png"></h4>','2018-03-17 19:43:11','2018-03-17 19:43:11'),
('58','649','0','579','<p>121212121212121</p>','2018-03-23 12:27:32','2018-03-23 12:27:32'),
('59','649','0','579','<p>Please try again. Your captcha couldn\'t be validated. If this happens again please contact info@pexels.com.Please try again. Your captcha couldn\'t be validated. If this happens again please contact info@pexels.com.Please try again. Your captcha couldn\'','2018-03-23 12:28:57','2018-03-23 12:28:57'),
('60','112','0','627','<p>加油！！</p><p><br></p>','2018-03-29 18:13:01','2018-03-29 18:13:01'),
('61','112','0','557','<p>加油！！这么好的网站应该得到更多人的支持！！！！</p>','2018-03-30 13:10:39','2018-03-30 13:10:39'),
('62','64','0','737','<p>毛绒绒的猫咪爱玩毛线团</p>','2018-04-13 17:57:34','2018-04-13 17:57:34'),
('63','112','0','843','<p>大佬！怎么搜图啊</p><p><br></p>','2018-04-29 19:06:29','2018-04-29 19:06:29'),
('64','112','0','866','<p>加油！</p>','2018-05-02 17:20:27','2018-05-02 17:20:27'),
('65','112','0','866','<p>如果可以添加搜图的功能可能会更好</p>','2018-05-02 17:21:13','2018-05-02 17:21:13'),
('66','32','0','867','<p>飞得更高</p>','2018-05-11 17:25:24','2018-05-11 17:25:24'),
('67','112','0','942','<p>在“中国独立开发者项目列表”上看到这个项目，特来逛逛，感觉很棒；个人也有独立开发 &amp; 维护一个网站（ https://nicelinks.site/ ），用以收录各类优秀站点，以供更多人受用；美图方面，现在也收录了蛮，可参见： https://nicelinks.site/theme/picture ，深夜造访，特邀请您的这款 Web 应用入住，不知可有兴趣？</p>','2018-05-13 02:33:27','2018-05-13 02:33:27'),
('68','377','0','949','<p>好</p>','2018-05-14 20:14:06','2018-05-14 20:14:06'),
('69','112','0','182','<p>如果能搜索就更好了(●°u°●)​&nbsp;」</p>','2018-05-18 19:05:55','2018-05-18 19:05:55'),
('70','32','0','990','<p>简约</p>','2018-05-20 11:29:47','2018-05-20 11:29:47'),
('71','1119','0','1012','<p>.</p>','2018-05-23 20:57:39','2018-05-23 20:57:39'),
('72','1119','0','1012','<p>.</p><p><br></p>','2018-05-23 20:57:47','2018-05-23 20:57:47'),
('73','112','0','1047','<p>想知道怎么搜索 <img src="/static/image/1f60c.png"></p>','2018-05-28 14:35:05','2018-05-28 14:35:05'),
('74','21','0','1054','<p>超级震撼</p><p><br></p>','2018-05-29 18:41:10','2018-05-29 18:41:10'),
('75','112','0','1066','<p><img src="/static/image/1f600.png"></p>','2018-06-01 15:33:12','2018-06-01 15:33:12'),
('76','19','0','1135','<p>好萌</p>','2018-06-13 10:59:28','2018-06-13 10:59:28'),
('77','377','0','1136','<p>赞赞赞，求网盘打包地址<img src="/static/image/1f604.png">992649785@qq.com</p>','2018-06-13 20:42:27','2018-06-13 20:42:27'),
('78','112','0','1260','<p>加油加油</p>','2018-07-07 15:38:30','2018-07-07 15:38:30'),
('79','112','0','1264','<p>加个打赏功能</p>','2018-07-09 20:21:55','2018-07-09 20:21:55');