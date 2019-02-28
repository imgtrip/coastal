## bugs
- album-show componentDidMount时加载评论，当评论加载成功后，padding bottom 会一直false导致images load more失败
- go:image.go 如果已展示数据过多，会导致where in sql长度过长异常
- 

## 开发中

## 缺失
- images，albums 首次加载检测是否可以scroll，不可用时需要立即请求page 2
- 浏览器兼容测试 （根据官方库解释只需要支持http2即可,但是在chrome正常使用的情况下,最新版firfox及其他ie系未通过测试,尝试改进为http2后测试）
- seo 数据配置 
- 默认album cover
- download log √×
- 尺寸可选download  √× (剩余最终开始下载)
- GA统计
- 评论字数限制显示

## 已完成
>第一个√表示开发完成，第二个√表示测试完成
- register √√
    - 创建favorite album √√
    - 绑定token √√
    - 注册后添加/取消favorite √√
    - 注册后新建album √√
    - 注册后新建album并添加当前image到新建album √√  
    - 注册后添加到新建album √√
    - 添加喜欢/删除喜欢图集图片数量计数 √√
    - 添加到已有图集完成后关闭 modal √√
- login √
    - 已喜欢的图点亮红心 √√
    - 获取已创建的图集 √√
    - 用户信息 √√
    - 持久化登陆 √√
    - token更新 √√
- home images 
    - init √√
    - load more √√
    - 每次加载更多需要更换返回数据，当前为相同数据;**由于测试数据较少，未将已显示数据加入session** √× 
    - 加载更多会同时加载多个结果返回 √√
- collect 
    - to new √√
    - to exist √√
- favorite √√
- download ××
- albums 
    - init  √√
    - load more √√
- album-show
    - init album √√
    - update token 过期时间 √√
    - images init √√
    - images load more √√
- update album cover  √√
- remove album image √√
- update album title √√
- amounts √√
  views comments √√
- comments load more √√
- comment create  √√
- 键盘左右键切换zoom图 √√
- user-home
    - albums √√
    - logout √√
    - 用户头像 √√
-  zoom  log √√
- js / css 版本号更新机制 √√ 

### production

browser  - nginx:80 - nodeServer:3000 ---
browser  - nginx:80