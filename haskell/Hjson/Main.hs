module Main where

--import PutJSON
--import PrettyStub
import SimpleJSON
import PrettyJSON
import Prettify

a = JNumber 1.2
b = JNull
c = JBool False
d = JObject [("d", (JArray [(JNumber 1), (JNumber 2), (JNumber 3)]))]
e = JArray [JNull, JNull, (JBool True), (JNumber 1)]
s = JString "deadquin"
v = JArray [a, b, c, d, e]
o = JObject [("\\name", s), ("data", v)]

main = putStrLn (pretty 30 (renderJValue o))
