-- ループ

-- for（数値ループ）
-- 1から5までが表示される(4じゃないんだ...)
for i = 1, 5 do
	print(i)
end
print("\n--------------------\n")

-- for i = 開始, 終了, ステップ do ... end
for i = 10, 1, -2 do
	print(i)
end
print("\n--------------------\n")

-- while
local i = 1
while i <= 5 do
	print(i)
	i = i + 1
end
print("\n--------------------\n")

-- repeat...until（最低1回は実行される）
local i = 1
repeat
	print(i)
	i = i + 1
until i > 5
print("\n--------------------\n")

-- break（ループを抜ける）
for i = 1, 10 do
	if i == 5 then
		break
	end
	print(i)
end
-- ※ Luaには continue がない
-- ↑まじか

print("\n--------------------\n")
--[[
練習問題:
1. forで1から10まで出力してみてください
2. forで10から1まで逆順に出力してみてください（ステップに-1を指定）
3. whileで1から5まで出力してみてください
4. repeat...untilで同じことをやってみてください
5. breakを使って、5になったらループを抜けてみてください
]]
