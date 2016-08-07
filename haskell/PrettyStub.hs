module PrettyStub where

data Doc = ToBeDefined
        deriving (Show)

string :: String -> Doc
string str = undefined

double :: Double -> Doc
double num = undefined

text :: String -> Doc
text str = undefined

char :: Char -> Doc
char c = undefined

(<>) :: Doc -> Doc -> Doc
a <> b = undefined

hcat :: [Doc] -> Doc
hcat d = undefined

fsep :: [Doc] -> Doc
fsep xs = undefined
