# 第一题
Container接口当中多了Open和Close以及Remains方法
如果没有调用Open，那么调用以下方法时:
* PutIn返回error，内容是"Container Closed"
* PourOut返回nil

如果没有调用Close，那么调用Blender的Blend方法，返回error，内容是"Container Not Closed"

任何时候调用Remains方法都应该返回容器当前的剩余空间（Capacity减去已存在的物质的体积）

**注意！**

由于目前container这个struct的Open、Close和Remains方法直接调用会panic，请给出恰当的实现

# 第二题
新增PowerSource接口，有三个方法：
* Voltage()返回电源电压
* Ampere()返回电源电流
* Shape()返回电源插口形状，有"cn"、"uk"、"us"三种

Blender新增Voltage、Ampere、PlugShape三个Field。要求：
* Blender的Voltage、Ampere、PlugShape必须与PowerSource的Voltage()、Ampere()、Shape()完全相同，否则应用合适的方式返回error，内容是"Incompatible PowerSource"
* 如果Blender的PowerSource是nil，那么调用Blend方法应返回error，内容是"No Power"

# 第三题
新增Bubble接口，有TurnOn、TurnOff、SetColor三个方法，要求：
* SetColor只能在TurnOn以后执行，否则返回error，内容是"Bubble Is TurnOff"
* 每一个Blender创建时都**应该**有一个Bubble
* Blend时先TurnOn Bubble，然后SetColor("red")，搅拌完成后，在Blend函数返回前，Bubble要SetColor("green")，再TurnOff