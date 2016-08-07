module Book where

import Data.Char
import Data.Bits (shiftL, (.|.), (.&.))

maybeString :: Maybe t -> [Char]
maybeString (Just _) = "Just"
maybeString  Nothing = "Nothing"

main :: IO()
main = putStrLn "Hello, haskell!"

quickSort :: Ord a => [a] -> [a]
quickSort [] = []
quickSort (x:xt) = 
    quickSort [ a | a <- xt, a < x] 
    ++ [x] ++
    quickSort [ a | a <- xt, a >= x ]

prodList lst = 
    [product [1..a] | a <- lst]

data Customer = 
    Customer { 
        customerId :: Int, 
        customerName :: Name, 
        customerAddress :: Address 
    } deriving (Show)

data Name = 
    Name {
        firstName :: String,
        lastName :: String
    } deriving (Show)

data Address = 
    Address [String]
    deriving (Show)

greet :: Customer -> String
greet customer = 
    let name = (customerName customer) in
    "Hello, " ++ (lastName name) ++ (firstName name) ++ "!\n"

data List a = 
    Cons a (List a)
    | Nil
    deriving (Show)

insert :: a -> List a -> List a
insert a list = 
    Cons a list

safeTail :: [a] -> Maybe [a]
safeTail lst = 
    if null lst || null (tail lst) then
        Nothing
    else
        Just (tail lst)

shadow = 
    let x = 1 in
        ((let x = "A" in x), x)

shadow2 x = 
    let x = 1 in
        x

lend amount balance =
    if newBalance  < 100 then
        Nothing
    else
        Just newBalance
    where
        newBalance = balance - amount

pluralise :: String -> [Int] -> [String]
pluralise word counts = 
    map plural counts
    where 
        plural 0 = "no " ++ word ++ "s"
        plural 1 = "one " ++ word
        plural n = show n ++ " " ++ word ++ "s"

fileName = "Book.hs"

splitLines :: String -> [String]
splitLines str
    | null str = []
    | otherwise = let (pre, suf) = break isLineBreak str
                  in pre : case suf of
                                ('\r': '\n' : rest) -> splitLines rest
                                ('\n' : rest) -> splitLines rest
                                _ -> []

isLineBreak :: Char -> Bool
isLineBreak c = c == '\r' || c == '\n'

myAll :: (a -> Bool) -> [a] -> Bool
myAll f [] = True
myAll f (x:xt) = (f x) && (myAll f xt)

myAny :: (a -> Bool) -> [a] -> Bool
myAny f [] = False
myAny f (x:xt) = (f x) || (myAny f xt)

add3 :: Num a => a -> a -> a -> a
add3 a b c = a + b + c

splitWith f lst
    | null lst = []
    | otherwise = let t = (break f lst) in 
                    fst t : let s = snd t in
                                case s of
                                    x:xt -> [x] : splitWith f xt 
                                    _ -> []

strToInt :: String -> Int
strToInt str = strToInt' 0 str

strToInt' :: Int -> String -> Int
strToInt' acc (x:xt) = let acc' = acc * 10 + digitToInt x in
                        strToInt' acc' xt
strToInt' acc _ = acc

squareList :: Num a => [a] -> [a]
squareList (x:xt) = x * x : squareList xt
squareList _ = []

upperCase :: String -> String
upperCase (x:xt) = toUpper x : upperCase xt
upperCase _ = [] 

mySum :: Num a => [a] -> a
mySum lst = mySum' 0 lst

mySum' acc (x:xt) = mySum' (acc + x) xt 
mySum' acc _ = acc

base = 65521

adler32 xs = helper 1 0 xs
    where helper a b (x:xt) = let a' = mod (a + (ord x .&. 0xff)) base
                                  b' = mod (a' + b) base in
                                helper a' b' xt
          helper a b _ = (shiftL b 16) .|. a

adler32_foldl xs = let (a, b) = foldl step (1, 0) xs in
                    (shiftL b 16) .|. a
                   where step (a, b) x = let a' = a + (ord x .&. 0xff)
                                         in (mod a' base, mod (a' + b) base) 

myfoldl :: (a -> b -> a) -> a -> [b] -> a
myfoldl step zero (x:xt) = myfoldl step (step zero x) xt
myfoldl _ zero [] = zero

myfoldr :: (a -> b -> b) -> b -> [a] -> b
myfoldr step zero (x:xt) = step x (myfoldr step zero xt)
myfoldr _ zero [] = zero

myfilter p lst = foldr step [] lst
                where step x xt
                            | p x = x : xt 
                            | otherwise = xt

myFoldl :: (a -> b -> a) -> a -> [b] -> a
myFoldl f z xs = (foldr step id xs) z
                    where step x g a = g (f a x)

myMap :: (a -> b) -> [a] -> [b]
myMap f l  = foldr step [] l
                where step x xt = f x : xt

foldlSum :: Num a => [a] -> a
foldlSum xt = myfoldl step 0 xt
                where step acc x = acc + x

anyOdd lst = any (\e -> (mod e 2) == 1) lst

foldAny :: (a -> Bool) -> [a] -> Bool
foldAny f xt = foldl step False xt
                where step b y = b || f y 

dropAny :: (a -> Bool) -> [a] -> [a]
dropAny f xt = foldr step [] xt
                where { step y yt
                    | f y = yt
                    | otherwise = y : yt 
                }

qs (x:xt) = qs [a | a <- xt, a <= x] ++
            [x] ++
            qs [a | a <- xt, a > x] 
qs _ = []

myTails :: [a] -> [[a]]
myTails xt@(_:t) = xt : myTails t
myTails _ = [[]]

data MyBool = MyFalse
    | MyTrue

class MyEq a where
    isEqual, isNotEqual :: a -> a -> Bool
    isEqual x y = not (isNotEqual x y)
    isNotEqual x y = not (isEqual x y)

instance Show MyBool where
    show MyFalse = "iFalse"
    show MyTrue = "iTrue"

instance Eq MyBool where
    (==) MyFalse MyTrue = True
    (==) MyTrue MyFalse = True
    (==) _ _ = False
    (/=) MyFalse MyTrue = True
    (/=) MyTrue MyFalse = True
    (/=) _ _ = False

instance MyEq MyBool where
    isEqual (MyFalse) (MyTrue) = True
    isEqual (MyTrue) (MyFalse) = True
    isEqual _ _ = False
