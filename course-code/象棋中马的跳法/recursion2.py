# /*
#  * @Author: zijian.su 
#  * @Date: 2021-11-26 13:29:16 
#  * @Last Modified by:   zijian.su 
#  * @Last Modified time: 2021-11-26 13:29:16 
#  */

# recursion solution
def f(x, y, step, a, b):
    if x>=10 or x<0 or y>=9 or y<0:
        return 0
    if step == 0:
        if x==a and y==b:
            return 1
        else:
            return 0
    return f(x+2, y+1, step-1, a, b)+f(x-2, y+1, step-1, a, b)+f(x+1, y+2, step-1, a, b)+f(x-1,y+2, step-1, a, b)\
           +f(x+1, y-2, step-1, a, b)+f(x-1, y-2, step-1, a, b)+f(x+2, y-1, step-1, a, b)+f(x-2, y-1, step-1, a, b)

if __name__ == '__main__':
    x, y, k = map(int, input().split())
    print(f(0, 0, k, x, y))
