str1 = ")gdfgdsf(bfdgfdgf(bsedfg(gbdfs)hds)hbghfg))Bgh(GDFGDFG"

str2 = "gdfgdsf(bfdgfdgf(bsedfg(gbdfs)hds)hbghfg))Bgh(GDFGDFG"

str3 = "gdfgdsf(bfdgfdgf(bsedfg(gbdfs)hds)hbghfg))BghGDFGDFG"

str4 = "gdfgdsf())(bfdg"

str5 = "gdfgdsf(bfdgfdgf(bsedfg(gb()dfs)hds)hbghfg)BghGDFGDFG"


def validate_paranthesis(str0):
    if (str0.count("(")) != (str0.count(")")):
        return False

    ## stack
    s = []

    for ch in str0:
        if ch == '(':
            s.append('(')

        if ch == ')':
            if (len(s) == 0):
                return False
            else:
                s.pop()
    if (len(s) == 0):
        return True


if validate_paranthesis(str5):
    print("Sahi Hai")
else:
    print("Galat Hai")
