golang 易错点

1. 结构体比较规则：

   - 只有相同类型的结构体才能比较，结构体是否相同不仅与属性类型个数有关，还和属性顺序有关
   - 结构体相同，但结构体中属性有不可比较的类型，如map ，slice，则不能用==比较

2. 函数返回值类型：

   nil可用作interface、function、pointer、map、slice、channel的空值，但如果不特别指定的话，go 语言不能识别类型会报错,如下代码将报错。

   ```go
   func GetValue(m map[int]string, id int) (string, bool) {
   	if _, exist := m[id]; exist {
   		return "存在数据", true
   	}
   	return nil, false
   }
   ```

3. slice拼接问题

   两个slice进行append时，需将第二个slice进行 ... 打散再拼接

   ```go
   s1:=[]int{1,2,3}
   s2:=[]int{4,5}
   s1=append(s1,s2...)
   ```

4. interface是golang类型的父类，func  f(x interface{})的interface{}可传入golang的任何类型，包括指针，但是函数func g(x *interface{})只能接收  *interface。