-- 関数
-- 関数定義
local function greet(name)
	print("hi, " .. name .. "!")
end

greet("taro")

-- 複数戻り値
local function minmax(a, b)
	return math.min(a, b), math.max(a, b)
end

local min, max = minmax(10, 2)
print("min: " .. min .. ", max:" .. max)

-- 可変長引数（...）
local function sum(...)
	local total = 0
	for _, v in ipairs({ ... }) do
		total = total + v
	end
	return total
end

print(sum(1, 2, 3, 4))

-- クロージャ（関数が外部の変数を保持）
local function counter()
	local count = 0
	return function()
		count = count + 1
		return count
	end
end

local c = counter()
print(c())
print(c())

--[[
練習問題:
1. 引数を受け取って挨拶を出力する関数を作ってみてください
2. 2つの値を受け取り、小さい方と大きい方を両方返す関数を作ってみてください
3. 可変長引数で渡された全ての数値を足し算する関数を作ってみてください
4. カウンターのクロージャを作り、呼ぶたびに1ずつ増える関数を作ってみてください
]]
