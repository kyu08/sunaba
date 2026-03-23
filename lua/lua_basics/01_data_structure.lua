local function myprint(var)
	print(var, "->", type(var))
end

-- 基本的なデータ型は以下の4つ。
-- 1. nil
-- 2. bool
-- 3. number(整数も少数もnumber)
-- 4. string

local nil_var = nil
myprint(nil_var)

local bool = false
myprint(bool)

local number = 10.1
myprint(number)

local str = "foo"
myprint(str)

-- cは未定義なのでnilが表示される
local a, b, c = 1, 2
myprint(c)

global_var = 10
myprint(global_var)
