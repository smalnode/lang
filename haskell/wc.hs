-- file: wc.hs
-- edited by deadquin
-- date: 20150203

main = interact wordCount
    where wordCount input = show (length (lines input)) ++ "\n"
