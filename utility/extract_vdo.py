import os
import sys

vdosubfixs = ['.wmv', '.avi', '.rmvb', '.mp4', '.mkv', '.jpg']

def isVdoFile(fl):
    for subfix in vdosubfixs:
        if fl.endswith(subfix):
            return True
    return False

if  len(sys.argv) > 1:
    path = sys.argv[1]
    if not os.path.exists(path):
        raise ValueError("path [%s] doesn't exists" % (path)) 
else:
    raise ValueError("no path is specified")

subdirs = [ subdir for subdir in os.listdir(path) 
        if os.path.isdir("%s/%s" % (path, subdir))]

for subdir in subdirs:
    abssubdir = "%s/%s" % (path, subdir)
    files = [fl for fl in os.listdir(abssubdir) 
            if not os.path.isdir("%s/%s" % (abssubdir, fl)) and isVdoFile(fl)] 
    for fl in files:
        oldfile = "%s/%s" % (abssubdir, fl)
        newfile = "%s/%s" % (path, fl)
        try:
            os.rename(oldfile, newfile)
        except PermissionError as pe:
            print(pe.strerror)
        except FileExistsError:
            pass
