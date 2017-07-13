-- Learn You a Haskell for Great Good!
-- Starting Out
-- mtso 2017

doubleMe x = x + x

doubleUs x y = x*2 + y*2

doubleSmallNumber x =
    if x > 100
        then x
        else x*2

doubleSmallNumber' x = (if x > 100 then x else x*2) + 1

conanO'Brien = "It's a-me, Conan O'Brien"

