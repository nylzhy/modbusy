# 需实现目标

- [ ] PLC/RAW API 目标
- [ ] 基于通讯信道的通讯队列实现，队列需要分高优先/低优先，临时生成写数据为高优先，配置读取为低优先级
- [ ] 实现写入 `Restful` 监听
- [ ] 配置文件导入，配置信息定义
- [ ] 独立的基于命令解析模块，是否可能其他地方复用？
- [ ] 建立实时标签map，反馈数据
- [ ] 建立基于 `restful` 的数据反馈通道？/ MQTT 订阅模式 / OPCUA 通道

## 第一优先

- [ ] 规划导入信息，JSON 文件，考虑通讯信号、通讯命令、通讯标签
- [ ] 规划优先队列信息，按照每接口
- [ ] 建立标签库，实时数据库
- [ ] 命令解析

- [ ] 命令标签是绑定设备，还是绑定命令？

长久来看，看到设备是更合理的，但是绑定设备的实质就是绑定到命令上，因此标签需要写清楚基于哪一条命令？，必须给每条命令起名字
