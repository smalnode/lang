from enum import Enum

class EventType(Enum):
    Null = 0
    Mouse = 1
    Key = 2
    Sys = 3
    Fatal = 4

class NullHandler:
    def __init__(self, successor=None):
        self.__successor = successor

    def handle(self, event):
        if event.kind == EventType.Null:
            pass
        else:
            self.passtonext(event)

    def passtonext(self, event):
        if self.__successor is not None:
            self.__successor.handle(event)


class KeyHandler(NullHandler):
    def handle(self, event):
        if event.kind == EventType.Key:
            print("Key handled: %s" % (event.value))
        else:
            self.passtonext(event)

class MouseHandler(NullHandler):
    def handle(self, event):
        if event.kind == EventType.Mouse:
            print("Mouse handled: %s" % (event.value))
        else:
            self.passtonext(event)

class SysHandler(NullHandler):
    def handle(self, event):
        if event.kind == EventType.Sys:
            print("Sys handled: %s" % (event.value))
        else:
            self.passtonext(event)

class Event:
    def __init__(self, kind, value):
        self.kind = kind
        self.value = value

def main():
    handler = MouseHandler(KeyHandler(SysHandler(NullHandler())))
    e = Event(EventType.Sys, "low battery")
    f = Event(EventType.Mouse, "left clicked")
    g = Event(EventType.Fatal, "fatal error")
    handler.handle(e)
    handler.handle(f)
    handler.handle(g)
    
if __name__ == '__main__':
    main()
