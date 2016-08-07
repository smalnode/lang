# abstract factory create objects that are used to assemble large objec, all 
# these objects are of one "family"
# first concrete factory a define all create_xxx method, the others only define
# object they created, factory a, a <- b, a <- c
class DiagramFactory:
    def create_text(self, text, fontsize=12):
        return self.Text(text, fontsize)

    class Text:
        def __init__(self, text, fontsize):
            self._text = text
            self._fontsize = fontsize

        def __str__(self):
            return ("Text('%s', %d)" 
                    % (self._text, self._fontsize))

class SvgDiagramFactory(DiagramFactory):
    class Text:
        def __init__(self, text, fontsize):
            self._text = text
            self._fontsize = fontsize
            self._color = 'BLACK'

        def __str__(self):
            return ("SvgText('%s', %d, %s)" 
                    % (self._text, self._fontsize, self._color))

def main():
    diagramFactory = DiagramFactory()
    t1 = diagramFactory.create_text('Plain Text', 16)
    svgFactory = SvgDiagramFactory()
    t2 = svgFactory.create_text('Svg Text', 12)

    print("DiagramFactory.create_text('Plain Text', 16) = ", t1)
    print("SvgDiagramFactory.create_text('Svg Text', 12) = ", t2)

if __name__ == '__main__':
    main()
