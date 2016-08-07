main = do 
    str <- getLine
    let d = (read str)::Double 
    putStrLn (show (truncate (2 * d)))
