-- メタテーブル
-- テーブルの振る舞いをカスタマイズする仕組み
-- __index（存在しないキーへのアクセス時に呼ばれる）
local defaults = { x = 100, y = 100 }
local point = { x = 10 }
-- 存在しないキー(ここではy)にアクセスされたときは__indexに登録したテーブルの中身が返る
setmetatable(point, { __index = defaults })

print(point.x) -- 10
print(point.y) -- 100

-- __newindex（存在しないキーへの代入時に呼ばれる）
local t = {}
-- NOTE: バリデーションをしたいときとか↓のようにログを残したいときに便利らしい
setmetatable(t, {
	__newindex = function(table, key, value)
		print("Setting [" .. key .. "] = " .. value)
		-- ここも自分で呼ぶ必要があるので注意
		rawset(table, key, value)
	end,
})
t.foo = "bar"
print(t.foo)

-- OOP風のクラス
-- Dog.__index = Dog と function Dog:method() の組み合わせ
---@class Dog
---@field name string
local Dog = { name = "no name dog" }
Dog.__index = Dog

-- NOTE: これを定義しておくとprint(dog)のようにインスタンスを渡したときの
-- 出力をカスタマイズできる。(GoでいうStringer interfaceの実装と同じ?)
---@param t Dog
Dog.__tostring = function(t)
	return string.format("Hi! I am %s!!!", t.name)
end

function Dog.new(name)
	if name == nil then
		-- NOTE: setmetatableをしないとbarkなどのメソッドが使えるようにならないので注意
		return setmetatable({}, Dog)
	else
		return setmetatable({ name = name }, Dog)
	end
end

function Dog:bark()
	print("I'm: " .. self.name)
end

local d = Dog.new()
-- d.bark()でも呼べてしまうのややこしい...
d:bark()
local d2 = Dog.new("shiba")
d2:bark()

print(d)
print(d2)

--[[
練習問題:
1. __index を使ってデフォルト値を持つテーブルを作ってみてください
2. __add を使って2つのテーブル（ベクトル）を足し算できるようにしてください
3. Dog クラスを参考に、Cat クラスを作ってみてください（meow()メソッド付き）
4. __tostring を使って、print(obj) でカスタム文字列を出力してみてください
]]
