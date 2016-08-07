module PutJSON where

import Data.List (intercalate)
import SimpleJSON

renderValue :: JValue -> String
renderValue (JString s) = show s
renderValue (JNumber n) = show n
renderValue (JBool True) = "true" 
renderValue (JBool False) = "false" 
renderValue JNull = "null"
renderValue (JObject o) = 
    "{ " ++ pairs o ++ " }"
    where pairs [] = ""
          pairs xt = intercalate ", " (map renderPair xt) 
          renderPair (key, value)  = show key ++ ": " ++ renderValue value

renderValue (JArray a) =  
    "[ " ++ values a ++ " ]"
    where values [] = ""
          values xt = intercalate ", " (map renderValue xt)

putJValue :: JValue -> IO ()
putJValue v = putStrLn (renderValue v)
