# Makefileだと定義したコマンドはsub shellで実行されるため、interactive shellのディレクトリを変更することができない。
# その回避策として、実行するコマンドを単にpbcopyし、interactive shell側でそれをcmd + vして実行する運用にしてみる。
.PHONY: new-command
new-command:
	echo "cp -rf x-y ${ARG} && cd ${ARG} && nvim main.c" | pbcopy

.PHONY: new
new:
	cp -rf x-y ${ARG}
