-- モジュールの読み込みと使用

-- require でモジュールを読み込む
-- 同じディレクトリなら require("モジュール名") でOK
local mymodule = require("09_modules/mymodule")
-- NOTE: requireのパスはlua foo.luaを実行したディレクトリから見た相対パスを指定する必要がある。
-- なので 09_modules直下でlua main.luaを実行するなら上記のようにrequire("mymodule")で問題ないが、
-- もう一つ上のディレクトリでlua ./09_modules/main.luaを用に実行するのであれば
-- require("09_modules/mymodule")のようにする必要がある。

-- モジュールの関数を呼び出す
print(mymodule.greet("taro"))

-- package.path で検索パスを確認できる
print("package.path:", package.path)

--[[
練習問題:
1. mymodule.lua に greet(name) 関数を追加してください
2. main.lua で require して greet を呼び出してください
3. mymodule.lua にプライベート関数を追加し、パブリック関数から呼び出してみてください
4. main.lua からプライベート関数にアクセスできないことを確認してください

実行方法:
  cd 09_modules && lua main.lua
]]
