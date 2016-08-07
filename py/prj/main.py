import os
import sys



if __name__ == "__main__" and __package__ is None:
    # [parent_dir]/prj/[__file__]
    parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    os.sys.path.insert(1, parent_dir)
    mod = __import__('prj')
    sys.modules["prj"] = mod
    __package__='prj'

    from .complex.Complex import *

    a = Complex(1, 2)
    b = Complex(1, -1)
    print(compmulti(a, b))
