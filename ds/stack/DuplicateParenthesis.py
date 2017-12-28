s="((a+(b))+(c+d))" 
#s="((a+b)+((c+d)))"
print ("Q: Is this ",s, " expression have duplicate parenthesis? ")
def dup(s):
    l1=[]
    top=''
    for c in s:
        if (c==')'):
            top=l1.pop()
            if(top=='('):
                return True
            else:
                while(top!='('):
                    top=l1.pop()
        else:
            l1.append(c)
    return False
if(dup(s) is True):
    print("Yes")
else:
    print("A: No")