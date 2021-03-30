# 写在开头
- 2013 年 5 月，入行互联网领域，从事UI设计（原生 PHP 撸了一个简单的博客）
- 2014 年转做切图仔，二开诸多开源CMS（当时流行），累计了前端布局与二开经验
- 2015 年继续切图工作，开始涉及运维，接触 ThinkPHP 3.2.3（撸了个 CRM 系统）
- 2016 年 8 月，入职新公司，基于公司现有系统（ThinkPHP），继续堆功能
- 2017 年继续后端工作，开始接触 React 前端开发（全新的领域）
- 2018 年 6 月，老大离职，接手他的岗位，转技术管理岗（加点工资，一把梭，偶尔兼任 UI 角色）
- 2019 年带领前端小伙伴，开始新的前端项目开发（React + Typescript）
- 2020 年疫情来袭
  - 10 天时间，一个人基于现有的系统，开发视频学习平台（后台 + 微信端 + 后端）
  - 上半年主要忙项目迭代优化
  - 下半年偏后台开发的事多点（开始接触 Go 语言）
- 2021 年开始用 Go，为自己开发一个博客系统
  - 说起来真丢人，从业 8 年，还没有给自己做过一个项目
  - 决心 2021，用新的语言为自己开发一个博客

# 代码托管
- 采用大仓库模式

# 架构
- 前后端分离开发

# 技术栈
## 前端
- React
- Redux
- Antd
- ReduxForm / Formik

## 后端
- Go

# Todo
- [x] 配置
- [x] 缓存
- [x] Redis
- [x] 图片验证码
- [x] 邮箱验证码
- [x] 手机验证码
- [x] 多语言支持
- [x] 响应格式化
- [ ] 参数验证
- [ ] 账号平台
  - [ ] 登录
    - [ ] 手机号登录
    - [ ] 邮箱登录
    - [ ] 账号密码登录
  - [ ] 注册
    - [ ] 手机号注册
    - [ ] 邮箱注册
    - [ ] 第三方注册
  - [ ] 个人资料
    - [ ] 设置账号
    - [ ] 设置密码
    - [ ] 修改密码
    - [ ] 新增手机号绑定
    - [ ] 解除手机号绑定
    - [ ] 换绑手机号
    - [ ] 新增邮箱绑定
    - [ ] 解除邮箱绑定
    - [ ] 换绑邮箱
    - [ ] 新增第三方绑定
    - [ ] 解除第三方绑定
  - [ ] 找回密码
    - [ ] 通过手机号找回
    - [ ] 通过邮箱找回
- [ ] 后台

# 响应码

值 | 描述
---|---
0 | 成功
10001 | 系统错误
10002 | 参数错误
10003 | 图片验证码错误