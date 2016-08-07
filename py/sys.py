import sys 

str1 = "abc"
str2 = 'abc'
str3 = '''abc
def'''

def main() :
    print sys.argv
    print "%s %s %s" % (str1, str2, str3)

if __name__ == '__main__' :
    main()
