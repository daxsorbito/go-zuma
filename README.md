# Zuma game coding challenge

The game would require two inputs, the `board` and `balls`. Each `ball` would be inserted into the board. First the ball would be inserted into the `board` where it has a matching character. Then if the matching character's subsequent count is more than 2 then remove that set of characters from the board. If `ball` is not found in the `board` append the `ball` into the `board`. Repeat the process until all `ball` in the `balls` are consumed.

1. Board: "WWRRRBB", Balls: "RBW"
   - Input: WWRRRBB, RBW
   - Output: ""
   - Process: WWRR[R]BB -> WWBB -> WWBB[B] -> WW -> [W]WW -> ""
1. Board: "QWERRDDSA", Balls: "RDSSEAQA"
   - Input: QWERRDDSA, RDSSEAQA
   - Output: QQWWEE
   - Process: QWERR[R]DDSA -> QWEDDSA -> QWED[D]DSA -> QWESA -> QWES[S]A -> QWESSA -> QWESS[S]A -> QWEA -> QWE[E]A -> QWEEA -> QWEEA[A] -> QWEEAA -> Q[Q]WEEAA -> QQWEEAA -> QQWEEAA[A] -> QQWEE
1. Board: "ASDFFDSAFFFFRT", Balls: "AADSDFSF"
   - Input: ASDFFDSAFFFFRT, AADSDFSF
   - Output: DSART
   - Process: [A]ASDFFDSAFFFFRT -> AASDFFDSAFFFFRT -> [A]AASDFFDSAFFFFRT -> SD[D] -> FDSAFFFFRT -> SDDFFDSAFFFFRT -> [S]SDDFFDSAFFFFRT -> SSDDFFDSAFFFFRT -> SS[D] ->DDFFDSAFFFFRT-> SSFFDSAFFFFRT -> SS[F]FFDSAFFFFRT -> SSDSAFFFFRT -> [S]SSDSAFFFFRT -> DSAFFFFRT -> DSA[F]FFFFRT -> DSART
