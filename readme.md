# ifuzz

练习编程项目

# 1.0版本

使用 -i 来指定输入字典的处理方式默认为"chain"，例如：-i "zip" -i "product"

chain：将输入的两个字典合为一个一维数组

zip：实现了python中zip函数，将传入的两个字典数组进行zip处理后返回一个二维数组

```
>>> a = [1,2,3]
>>> b = [4,5,6]
>>> zipped = zip(a,b)     # 打包为元组的列表
[(1, 4), (2, 5), (3, 6)]
```

product： 实现了python中product函数,返回A和B中的元素组成的笛卡尔积的元组

使用 -a 来设计攻击模式（AttackType）例如：`Cluster bomb` ，`Sniper`，`Battering ram`

Sniper：对单个FUZZ参数进行FUZZ -n指定FUZZ的数量，对于Sniper -n只能为1

Cluster bomb：实现传入两个FUZZ参数，-i 只能指定`chain`，对传入的两个FUZZ实现一个字典两个FUZZ的笛卡尔积

Pitchfork：实现传入两个FUZZ参数，-i 只能指定`zip`，对传入的两个字典依cmd输入顺序，对相应FUZ0Z,FUZ1Z进行爆破

