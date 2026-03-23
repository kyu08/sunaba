-- 条件分岐

-- if / elseif / else / end

-- 真偽値の扱い
-- Luaでは nil と false のみがfalsy、それ以外は全てtruthy
local n = nil
if n or false then
	print("never")
else
	print("hi")
end

-- 三項演算子の代用: condition and "yes" or "no"

--[[
練習問題:
1. 変数の値に応じて「positive」「negative」「zero」を出力してみてください
2. 0, "", nil それぞれを if で判定し、truthyかfalsyか確認してみてください
3. 三項演算子の代用パターンを試してみてください
]]
