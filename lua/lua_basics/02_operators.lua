-- 演算子
-- 算術演算子
print(10 + 3)

-- 比較演算子
print(1 == 1)
print(1 ~= 2)

-- 論理演算子
print(true and false)
print(true or false)
print(not true)

-- 短絡評価（and/or は値そのものを返す）
print(nil or "default")

-- 文字列連結
print("hello" .. ", " .. "world")

--[[
練習問題:
1. 10 / 3 と 10 // 3 の違いを確認してみてください
2. ~=（等しくない）を使って比較してみてください
3. nil or "デフォルト値" を実行し、デフォルト値のパターンを試してみてください
4. 数値と文字列を .. で連結してみてください
]]
