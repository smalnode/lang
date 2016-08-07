module SimpleJSON (
    JValue(..)
  , getString
  , getInt
  , getDouble
  , getBool
  , getObject
  , getArray
  , isNull
) where

data JValue = JString String
    | JNumber Double
    | JBool Bool
    | JNull
    | JObject [(String, JValue)]
    | JArray [JValue]
    deriving (Eq, Ord, Show)

getString :: JValue -> Maybe String
getString (JString v) = Just v
getString _ = Nothing

getInt :: JValue -> Maybe Int
getInt (JNumber v) = Just (truncate v)
getInt _ = Nothing

getDouble :: JValue -> Maybe Double
getDouble (JNumber v) = Just v
getDouble _ = Nothing

getBool :: JValue -> Maybe Bool
getBool (JBool v) = Just v
getBool _ = Nothing

getObject :: JValue -> Maybe [(String, JValue)]
getObject (JObject v) = Just v
getObject _ = Nothing

getArray :: JValue -> Maybe [JValue]
getArray (JArray v) = Just v
getArray _ = Nothing

isNull :: JValue -> Bool
isNull v = v == JNull
