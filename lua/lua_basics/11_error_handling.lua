-- エラーハンドリング
-- error（エラーを投げる）
---@param a number
---@param b number
---@return number
local function divide(a, b)
	if b == 0 then
		error("divided by zero")
	end
	return a / b
end
print(divide(10, 2))
-- print(divide(10, 0)) -- divided by zero
print("-----------------")

-- assert（条件を満たさなければエラー）
---@param name string
local function greet(name)
	assert(name, "name is required")
	print("Hello, " .. name)
end
-- greet(nil) -- name is required

-- pcall（エラーをキャッチして処理を継続）
-- 戻り値: ok(bool), result(成功時は戻り値、失敗時はエラーメッセージ)
local ok, result = pcall(function()
	return divide(10, 0)
end)

if ok then
	print("Result: " .. result)
else
	print("Error: ", result)
end

-- またはこう書くこともできる
-- 外部ライブラリの関数を呼ぶときはerrorをthrowしているかどうかを確認する必要がありそう。
-- また、標準ライブラリでは以下のnumber?, string?を返してそもそもエラーを返り値として
-- 扱うデザインパターンもあるらしい。(Goっぽい)
local ok, result = pcall(divide, 10, 0)
if ok then
	print("Result: " .. result)
else
	print("Error: ", result)
end

-- xpcall（スタックトレース付きでエラーをキャッチ）
local ok, result = xpcall(function()
	error("something wrong")
end, debug.traceback)
print(result)

--[[
練習問題:
1. 0で割ろうとしたら error() を投げる関数を作ってください
2. assert() を使って引数のバリデーションをしてください
3. pcall() でエラーをキャッチし、成功/失敗で分岐してください
4. xpcall() と debug.traceback でスタックトレースを表示してください
]]
