很多时候都有一个定时任务的需求，但是go貌似还没有带持久化任务管理的库
配合robfig/cron简单实现了下，逻辑挺简单，就不打算发布成一个库了，只是方便以后copy
日后要使用的时候再拓展，很多东西都还未完善，保存日志列表
设计初衷是灵活性，所以这里需要手动建表，维护任务的状态（当任务执行完成时执行回调函数改变状态什么的）
任务执行什么的，一般我们项目之中都有MySQL相关操作了，所以这里不封装了