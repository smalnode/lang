import abc

class BarRender(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def render(self, caption, pairs):
        pass

class BarCharter:
    def __init__(self, render):
        if not isinstance(render, BarRender):
            raise TypeError('BarCharter only accept BarRender')
        self.__render = render

    def render(self, caption, pairs):
        self.__render.render(caption, pairs)

class TextBarRender(BarRender):
    def render(self, caption, pairs):
        scalefactor = 40
        maxvalue = max([v for _, v in pairs])
        for t, v in pairs:
            print('*' * int(v / maxvalue * scalefactor), t)

class Text2BarRender(BarRender):
    def render(self, caption, pairs):
        scalefactor = 40
        maxvalue = max([v for _, v in pairs])
        for t, v in pairs:
            print('=' * int(v / maxvalue * scalefactor), t)



def main():
    barCharter = BarCharter(TextBarRender())
    barCharter.render('Week', [('Mon', 100), ('Tue', 255), ('Wed', 34), ('Thr', 49), ('Fri', 145), ('Sat', 200), ('Sun', 70)])
    barCharter2 = BarCharter(Text2BarRender())
    barCharter2.render('Week', [('Mon', 1), ('Tue', 2), ('Wed', 3), ('Thr', 4), ('Fri', 5), ('Sat', 6), ('Sun', 7)])

if __name__ == '__main__':
    main()
