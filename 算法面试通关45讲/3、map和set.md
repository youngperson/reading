# 哈希表
- 也叫散列表,来源于数组。存储的是key-val的映射集合
- 哈希表发生冲突，就使用链表进行存储
- 哈希表是数组的一种扩展,哈希表是数组和链表的结合体
- 数组的扩容,无非就是创建一个新的数组,将元素赋值进去之后将新数组的地址返回即可(数组越大扩容效率会低)
- 哈希表的扩容,可以选择尽量少让元素重新分配的哈希哈方式(哈希函数)

# 集合
- 相当于是哈希表中只需要key
- 实现方式有哈希表和树

# 哈希表和二叉搜索
- 哈希表查找快一些
- 二次搜索树中的元素是有序的

# 242 有效的字母异位词
```
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
示例 1:
输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。
进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
```

## 解法
```
func isAnagram(s string, t string) bool {
    same := make(map[rune]int, 0) 
    for _, v := range s {
        if _, found := same[v]; found {
            same[v] = same[v]+1
        } else {
            same[v] = 1
        }
    }

    for _, v := range t {
        if _, found := same[v]; found {
            same[v] = same[v]-1
        } else {
            return false
        }
    }

    for _, v := range same {
        if v != 0 {
            return false
        }
    }
    return true
}
时间复杂度O(n),空间复杂度O(n)
记忆：借助1个map先对字符进行+1计数+对另外一个字符串中的字符进行-1计数,最终map中的技术都为0才是字母异位
```

# 1 两数之和
```
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

## 解法
```
func twoSum(nums []int, target int) []int {
    usedindex := make(map[int]int, 0)   // key为val,val为index
    rtn := make([]int, 0)
    for index, v := range nums {
        findVal := target - v 
        if val, found := usedindex[findVal]; found {
            rtn = append(rtn, val, index)
        } else {
            usedindex[v] = index
        }
    }
    return rtn
}
记忆：借助map把下表作为val存储起来+forloop数组
```