from html.parser import HTMLParser
from urllib.request import urlopen

def main() :
    resp = urlopen('http://thz.la/forum.php?mod=forumdisplay&fid=220&page=1')
    parser = HTMLParser()
    parser.feed()

if __name__ == '__main__' :
    main()
    
