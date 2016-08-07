# diff from abstract factory, builder define all create_xxx method for 
# subclass, in another word, builder hold the entries of object it create.
# abstract class s, concrete factory s <- a, s <- b
import abc

class AbstractGuiFactory(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def __str__(self):
        pass

    def create_text(self, text, fontsize=12):
        return self.Text(text, fontsize)

class TextGuiFactory(AbstractGuiFactory):
    def __str__(self):
        return 'TextGuiFactory'

    class Text:
        def __init__(self, text, fontsize):
            self._text = text

class SvgGuiFactory(AbstractGuiFactory):
    def __str__(self):
        return 'SvgGuiFactory'

    class Text:
        def __init__(self, text, fontsize):
            self._text = text
            self._fontsize = fontsize

def main():
    f1 = TextGuiFactory()
    print(f1.create_text('Plant Text', 16))
    f2 = SvgGuiFactory()
    print(f2.create_text('Svg Text', 16))

if __name__ == '__main__':
    main()
