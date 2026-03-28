-- 型付け（LuaLS型アノテーション）
-- LuaLS（Lua Language Server）を使うとコメントで型を書ける
-- VSCodeなどで補完やエラーチェックが効くようになる

-- 基本の型アノテーション
---@type string
local str = "hi"

---@type number
local num = 100

---@boolean
local is_foo = false

---@type string[]
local members = { "taro", "jiro" }

-- 関数の型
---@param name string
---@param age number
---@return string
local function greet(name, age)
	return string.format("Hello %s, your are %d", name, age)
end

print(greet("taro", 10))

-- nilを許容する型（?をつける）
---@type string?
local maybe_name = "foo"
print(maybe_name)
print("------------------------------")

-- クラス風の型定義
---@class Person
---@field name string
---@field age number

---@type Person
local person = { name = "john", age = 10 }
print(person)
print("------------------------------")

-- エイリアス（型の別名）
---@alias UserId number

---@param id UserId
local function getUser(id)
	print("foo")
end

--[[
練習問題:
1. 変数に ---@type で型アノテーションをつけてみてください
2. 関数に ---@param と ---@return をつけてみてください
3. ---@class でオブジェクトの型を定義してみてください
4. 間違った型の値を代入して、LuaLSが警告を出すか確認してみてください
]]
