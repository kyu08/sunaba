-- テーブル（配列・辞書）
-- 配列としてのテーブル
-- インデックスは1から始まる(?!)
local fruits = { "apple", "banana", "cherry" }
print(fruits[1])
print(fruits[2])

-- 辞書としてのテーブル
-- person.name と person["name"] は同じ
local person = {
	name = "taro",
	age = 25,
}
print(person.name)
print(person.age)

-- 長さ演算子 #
local arr = { 1, 2, 3, 4, 5 }
print(#arr)

-- ipairs: 配列部分を順番にイテレート（1, 2, 3...）
for i, v in ipairs(fruits) do
	print(i, v)
end

-- pairs: 全てのキーをイテレート（順不同）
for k, v in pairs(person) do
	print(k, v)
end

--[[
練習問題:
1. 配列を作成し、インデックス1でアクセスしてみてください
2. 辞書を作成し、ドット記法とブラケット記法の両方でアクセスしてみてください
3. # で配列の長さを取得してみてください
4. ipairs で配列をループしてみてください
5. pairs で辞書をループしてみてください
]]
