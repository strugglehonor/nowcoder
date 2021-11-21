#  * @Author: zijian.su 
#  * @Date: 2021-11-21 19:27:59 
#  * @Last Modified by: zijian.su 
#  * @Last Modified time: 2021-11-21 19:27:59 
N = int(input())
arr = input()
nums = [int(x) for x in arr.split()]
Map = dict()
Map[0] = -1
maxlen = 0
for i in range(N):
    if nums[i] < 0:
        nums[i] = -1
    if nums[i] > 0:
        nums[i] = 1
# targer为0（目标和）
sumval, maxlen, target = 0, 0, 0
for j in range(N):
    sumval += nums[j]
    if sumval not in Map.keys():
        Map[sumval] = j
    if sumval-target in Map.keys():
        t = Map[sumval-target]
        maxlen = max(maxlen, j-t)
print(maxlen)