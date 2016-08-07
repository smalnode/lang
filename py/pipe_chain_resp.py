from enum import Enum
from asyncio import coroutine

class EventType(Enum):
    Null = 0
    Mouse = 1
    Key = 2
    Sys = 3
    Fatal = 4

@coroutine
def key_handler(successor=None):
    while True:
        event = (yield)
        if event is not None and  event.kind == EventType.Key:
            print("Key handled: %s" % (event.value))
        else:
            if successor is not None:
                successor.send(event)

@coroutine
def mouse_handler(successor=None):
    while True:
        event = (yield)
        if event is not None and event.kind == EventType.Mouse:
            print("Mouse handled: %s" % (event.value))
        else:
            if successor is not None:
                successor.send(event)

@coroutine
def sys_handler(successor=None): 
    while True:
        event = (yield)
        if event is not None and event.kind == EventType.Sys:
            print("Sys handled: %s" % (event.value))
        else:
            if successor is not None:
                successor.send(event)

class Event:
    def __init__(self, kind, value):
        self.kind = kind
        self.value = value

def main():
    e = Event(EventType.Sys, "low battery")
    f = Event(EventType.Mouse, "left clicked")
    g = Event(EventType.Fatal, "fatal error")
    kh = key_handler()
    kh.send(None)
    mh = mouse_handler(kh)
    mh.send(None)
    sh = sys_handler(mh)
    sh.send(None)
    pipe = sh
    pipe.send(e)
    pipe.send(f)
    pipe.send(f)
    pipe.send(e)
    pipe.send(f)
    pipe.send(f)
    pipe.send(e)
    pipe.send(f)
    pipe.send(f)
    pipe.send(f)
    pipe.send(f)

if __name__ == '__main__':
    main()
