####路由组件
####需求：
#####1.请求被解析成一个路由和关联的参数
#####2.路由相关的一个控制器动作被创建出来处理这个请求
#####3.如果传入请求并没有提供一个具体的路由，（一般这种情况多为于对首页的请求）此时就会启用由 Application::defaultRoute 属性所指定的缺省路由
#####4.catchAll 路由（全拦截路由）有时候，你会想要将你的 Web应用临时调整到维护模式，所有的请求下都会显示相同的信息页。 当然，要实现这一点有很多种方法。这里面最简单快捷的方法就是在应用配置中设置 Application::catchAll 属性