# godi
Golang  IOC(DI) Framework

## 特点

- 极简API
- 兼容性强，能支持绝大多数场景

## 功能list

- [x] 支持快速将struct定义为bean
- [x] 依赖自动注入，无需提前创建bean和显式声明注入
- [ ] 容器暴露方法可以在运行时获取bean
  - [x] 支持通过泛型类型获取bean
  - [ ] 支持通过bean名称获取bean
  - [ ] 支持通过接口获取bean
- [ ] 支持非单例bean
- [ ] 支持接口注入
- [x] 支持注入后自定义初始化
- [ ] 支持并发注入
- [ ] 支持命名注入，允许重名bean
- [ ] 支持注入bean list、bean map