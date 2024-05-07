package main
import (
	"fmt"
	"strconv"
	"strings"
)
	

func main(){
	//将数值转换为字符,两种：1.Sprintf;2.strconv
	var n1 int =10
	var s1 string=fmt.Sprintf("%d",n1)
	fmt.Printf("s1的数据类型是%T,s1=%v \n",s1,s1)
	var s2 string =strconv.FormatInt(int64(n1),10)//参数：第一个必须转为int,第二个参数指定字面值的进制形式
	fmt.Printf("s2的数据类型是%T,s2=%v \n",s2,s2)
	var n3 float64=4.25
	var s3 string =strconv.FormatFloat(n3,'f',5,64)//第一个参数是需转换的值，第二个是格式，第三个是小数点后位数，第四个是float的位数
	fmt.Printf("s3的数据类型是%T,s3=%v \n",s3,s3)
	//将字符转换为数值,s使用strconv
	var z1 int64
	z1,_ =strconv.ParseInt(s1,10,64)//第一个参数是需转换的值，第二个是参数指定字面值的进制形式，第三个是转换为多少位的int
	fmt.Printf("z1的类型为%T,z1=%d",z1,z1)

	//法1：for range可遍历数组、切片、字符串、map及通道，其格式为：
	// for key,val:=range coll{
		
	// }
	var str string="hello golang你好"
	for i,value:=range str{
		fmt.Printf("索引为%d,具体值为%c\n",i,value)
	}
	//法2：利用r:=[]rune(str)
	str1:=[]rune(str)
	for i:=0;i<len(str1);i++{
		fmt.Printf("%c\n",str1[i])
	}

	//strcov.Atoi
	num1,_:=strconv.Atoi("123")
	fmt.Println(num1)
	//strcov.Itoa
	num2:=strconv.Itoa(58)
	fmt.Println(num2)

	//统计字符串中出现的字串数目
	count:=strings.Count("hellogolang","go")
	fmt.Println(count)
	//不区分大小写的字符串比较
	falg:=strings.EqualFold("go","GO")
	fmt.Println(falg)
	//区分大小写的比较
	fmt.Println("go"=="GO")
	//返回字串在字符串第一次出现的索引值，如果没有返回-1
	fmt.Println(strings.Index("hellogolang","gaq"))
	//字符串的替换
	str2:=strings.Replace("hellogogogo","go","golang",-1)//-1表示全部替换，若为n，则代表替换n个
	fmt.Println(str2)
	//指定某个字符为分割符，并按分割符进行拆分成字符串
	arr:=strings.Split("go-java-python","-")
	fmt.Println(arr)
	//大小写切换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("gO"))
	//字符串左右两边的空格去掉
	fmt.Println(strings.TrimSpace("    go     "))
	//将字符串左右两边指定的字符去掉,同理有TrimRight,TrimLeft
	fmt.Println(strings.Trim("yjavagokyly","y"))
	//判断字符串是否以指定字符串开头,同理有判断结尾：HasSuffix
	fmt.Println(strings.HasPrefix("hhtp://go.com","http"))
}