# godi

Golang  IOC(DI) Framework

## 特点

- 极简API
- 兼容性强，能支持绝大多数场景

## 功能list
- [ ] 定义Bean（Provider）
  - [x] 支持工具方法定义Bean，支持指定name、是否单例、原始对象、类型等
  - [x] 支持通过“继承”定义Bean
  - [ ] 支持使用放通过tag将某个依赖对象定义为Bean
  - [x] Bean自定义属性
    - [ ] 支持通过实现GetBeanProps方法定义Bean属性，如name、是否单例等
    - [ ] 支持将Bean定义为私有的（非单例）
    - [ ] 支持注入后自定义初始化
      - [ ] 支持bean实现Init方法，注入完成后，自动调用init方法
      - [ ] 定义Bean的工具方法，支持传入初始化回调，在注入完成调用回调初始化
- [ ] 依赖关系与注入（Consumer）
  - [x] 只要有Autowired tag的，就自动注入，默认匿名注入。
  - [x] 支持Autowired tag，支持指定名称，如果不指定名称，则默认按类型进行注入。
  - [x] 支持内嵌成员（匿名成员）的依赖注入
  - [ ] 支持可选注入（弱依赖）
  - [x] 支持私有注入
  - [ ] 支持interface类型成员默认按类型进行匿名注入，取第一个实现了interface的bean实例
  - [ ] 支持指定interface类型，将bean注入到slice、map
  - [ ] 支持Setter方法注入???
- [ ] 容器&运行时工具方法（Runtime）
  - [ ] 同类型只能有唯一的匿名Bean（非单例的除外），可以有多个不同名但同类型的Bean
  - [ ] 支持运行时获取Bean
    - [ ] 支持通过泛型类型（struct、interface等）获取单个Bean
    - [ ] 支持通过Bean名称获取单个Bean
    - [ ] 支持通过interface获取多个Bean
  - [ ] 支持获取全部的bean实例
- [ ] 支持并发的依赖注入
- [ ] 支持对Bean进行代理封装??? - go动态代理实现（非代码生成方案） https://github.com/cocotyty/dpig
- [ ] 