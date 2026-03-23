-- 文字列操作

-- 文字列の長さ #
local s = "str"
print(#s)

-- string.sub（部分文字列）
-- string.sub(s, 開始, 終了)
-- 負のインデックスは後ろから
local s = "hello world"
print(string.sub(s, 1, 5)) -- 1も5も含まれる
print(string.sub(s, 7)) -- 1も5も含まれる

-- string.find（検索）
-- 開始位置と終了位置を返す
local s = "hello world"
local start, finish = string.find(s, "world")
print(start, finish)

-- string.format（フォーマット）
-- %s: 文字列, %d: 整数, %f: 小数
print(string.format("Name: %s, Age: %d", "Taro", 25))
print(string.format("%.2f", 3.14159))

-- パターンマッチング基礎
-- %d: 数字, %a: 英字, %s: 空白
-- +: 1回以上, *: 0回以上
local s = "hello123world"
print(string.match(s, "%d+"))

--[[
練習問題:
1. 文字列の長さを # で取得してみてください
2. string.sub で文字列の一部を取り出してみてください
3. string.find で特定の文字列を検索してみてください
4. string.format で変数を埋め込んだ文字列を作ってみてください
5. string.match でパターンマッチングを試してみてください
]]
