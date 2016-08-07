import os, time

def main():
    print "From parent: my pid is ", os.getpid()
    i = 5

    pid = os.fork()
    if pid:
        print "From parent, my pid after fork is ", os.getpid()
        time.sleep(0.2)
        print "From parent, global i = ", i
    else:
        print "From new process, my pid is ", os.getpid()
        print "From child, global i = ", i
        i = 10 # a local variable copy-on-write

    time.sleep(0.5)
    print "From common code path"

if __name__ == "__main__":
    main()
