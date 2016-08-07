module Whs.Wlib.Wsf where

import Data.Char

firstOrEmpty :: [[Char]] -> [Char]
firstOrEmpty lst = if not (null lst) then head lst else "empty"

reverse2 :: [a] -> [a]
reverse2 lst = 
    if null lst then 
        []
    else
        reverse2 (tail lst) ++ [ head lst ]

maxmin :: Ord a => [a] -> (a, a)
maxmin lst = 
    case lst of
    [x] -> (x, x)
    x:xt -> 
        (if x > t_max then x else t_max
         , if x < t_min then x else t_min)
        where {
            (t_max, t_min) = maxmin xt
        }

fabonacci :: Integer -> Integer
fabonacci n = 
    case n of
    0 -> 0
    1 -> 1
    _ -> fabonacci (n - 1) + fabonacci (n - 2)

ifabonacci :: Integer -> Maybe Integer
ifabonacci n = 
    if n < 0 then
    Nothing
    else
    case n of
    0 -> Just 0
    1 -> Just 1
    _ -> let Just f1 = ifabonacci (n - 1) 
             Just f2 = ifabonacci (n - 2) 
         in Just (f1 + f2)

data Client = 
    GovOrg String
    | Company String Integer Person String
    | Individual Person Bool
    deriving Show

data Person = 
    Person String String
    deriving Show

data ClientR = 
    GovOrgR { clientRName :: String }
    | CompanyR { clientRName :: String
                , companyId :: Integer
                , person :: PersonR
                , duty :: String }
    | IndividualR { person :: PersonR }
    deriving Show

data PersonR = 
    PersonR { firstName :: String
            , lastName :: String }
    deriving Show

greet :: ClientR -> String
greet IndividualR { person = PersonR { firstName } } = "Hi, " ++ firstName
greet CompanyR { person = PersonR { firstName = fn}} = "Hello, " ++ fn
greet GovOrgR { } = "Welcome"

nameInCapital :: PersonR -> PersonR
nameInCapital p@(PersonR { firstName = i:r }) = 
    let newName = (toUpper i):r in
    p { firstName = newName } 
nameInCapital p@(PersonR { firstName = "" }) = p

clientName :: Client -> String
clientName client = 
    case client of
    GovOrg name -> name
    Company name _ _ _ -> name
    Individual (Person fName lName) _ -> fName ++ " " ++ lName

companyName :: Client -> Maybe String
companyName client = case client of
    Company name _ _ _ -> Just name
    _ -> Nothing

(+++) :: [a] -> [a] -> [a]
list1 +++ list2 = 
    case list1 of
    [] -> list2
    (x:xt) -> x : (xt +++ list2)

wnull :: [a] -> Bool
wnull lst = 
    case lst of
    [] -> True
    _ -> False

whead :: [a] -> a
whead lst = 
    case lst of
    x:xt -> x

wtail :: [a] -> [a]
wtail lst =
    case lst of
    x:xt -> xt

wfst :: (a, a) -> a
wfst (a, b) = a

wsnd :: (a, a) -> a
wsnd (a, b) = b

sorted :: [Integer] -> Bool
sorted lst  =
    case lst of
    [] -> True
    x:r@(y:_) -> x <= y && sorted (r)
    [_] -> True

binom :: Integer -> Integer -> Integer
binom n k
    | k == 0 = 1
    | n == k = 1
    | otherwise = ((binom (n - 1) (k - 1)) + (binom (n - 1) k))

multipleOf :: Integer -> Integer -> Bool
multipleOf m n = (mod m n) == 0

specialMultiples :: Integer -> String
specialMultiples n  
    | multipleOf n 2 = show n ++ " is multiple of 2 "
    | multipleOf n 3 = show n ++ " is multiple of 5 "
    | multipleOf n 5 = show n ++ " is multiple of 3 "
    | otherwise = show n ++ " is a beautiful number"

ackermann :: Integer -> Integer -> Integer
ackermann m n
    | m == 0 = n + 1
    | m > 0 && n == 0 = ackermann (m - 1) 1
    | m > 0 && n > 0 = ackermann (m - 1) (ackermann m (n - 1))

responsibility :: Client -> String
responsibility (Company _ _ _ r) = r
responsibility _ = "Unknown"

specialClient :: Client -> Bool
specialClient (clientName -> "Deqin Wang") = True
specialClient (responsibility -> "Developer") = True
specialClient _ = False

data ConnType = TCP | UDP 
data UseProxy = NoProxy | Proxy String
data TimeOut = NoTimeOut | TimeOut Integer

data ConnOptions = ConnOptions {
    connType :: ConnType
    , connSpeed :: Integer
    , proxy :: UseProxy
    , connCaching :: Bool
    , connKeepAlive :: Bool
    , timeOut :: TimeOut 
    } 

connect' :: String -> ConnOptions -> String
connect' url options =
    "connected url: " ++ url

connDefault :: ConnOptions
connDefault = ConnOptions TCP 0 NoProxy True True NoTimeOut
