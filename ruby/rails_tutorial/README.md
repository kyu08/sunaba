## `rails _7.0.4.3_ new hello_app`したら依存関係のエラーが出た
```sh
❯ rails _7.0.4.3_ new hello_app
/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/activesupport-7.0.4.3/lib/active_support/logger_thread_safe_level.rb:12:in `<module:LoggerThreadSafeLevel>': uninitialized constant ActiveSupport::LoggerThreadSafeLevel::Logger (NameError)

    Logger::Severity.constants.each do |severity|
          ^^^^^^^^^^
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/activesupport-7.0.4.3/lib/active_support/logger_thread_safe_level.rb:9:in `<module:ActiveSupport>'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/activesupport-7.0.4.3/lib/active_support/logger_thread_safe_level.rb:8:in `<top (required)>'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/activesupport-7.0.4.3/lib/active_support/logger_silence.rb:5:in `<top (required)>'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/activesupport-7.0.4.3/lib/active_support/logger.rb:3:in `<top (required)>'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/activesupport-7.0.4.3/lib/active_support.rb:29:in `<top (required)>'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/railties-7.0.4.3/lib/rails/command.rb:3:in `<top (required)>'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/railties-7.0.4.3/lib/rails/cli.rb:12:in `<top (required)>'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from <internal:/Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/3.2.0/rubygems/core_ext/kernel_require.rb>:86:in `require'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2.6/lib/ruby/gems/3.2.0/gems/railties-7.0.4.3/exe/rails:10:in `<top (required)>'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2/bin/rails:25:in `load'
	from /Users/kyu08/.local/share/mise/installs/ruby/3.2/bin/rails:25:in `<main>'
```

Goはこの辺のストレスがないのでいきなり :ha: となった。

## scaffold
以下のコマンドで一気に関連するファイルが作成される。

```sh
$ rails generate scaffold User name:string email:string
      invoke  active_record
      create    db/migrate/<タイムスタンプ>_create_users.rb
      create    app/models/user.rb
      invoke    test_unit
      create      test/models/user_test.rb
      create      test/fixtures/users.yml
      invoke  resource_route
       route    resources :users
      invoke  scaffold_controller
      create    app/controllers/users_controller.rb
      invoke    erb
      create      app/views/users
      create      app/views/users/index.html.erb
      create      app/views/users/edit.html.erb
      create      app/views/users/show.html.erb
      create      app/views/users/new.html.erb
      create      app/views/users/_form.html.erb
      invoke    test_unit
      create      test/controllers/users_controller_test.rb
      create      test/system/users_test.rb
      invoke    helper
      create      app/helpers/users_helper.rb
      invoke      test_unit
      invoke    jbuilder
      create      app/views/users/index.json.jbuilder
      create      app/views/users/show.json.jbuilder
      create      app/views/users/_user.json.jbuilder
```

さらにその後`rails db:migrate`を実行して`localhost:3000/users/new`にアクセスしてブラウザをポチポチしたらユーザーが作成できた。

「??????」となったけど`config/database.yml`を見るとSQLiteに永続化しているらしい。悪い意味で隠蔽具合がすごすぎる。

SimpleとEasyの違いを感じる。(GoはSimpleでRailsはEasyだと感じている)

しかも`http://localhost:3000/users.json`にアクセスしたらuserテーブルの情報が全部書いてある。カジュアル情報漏洩やめて。（scaffoldが実際にどのくらい現場で使われているのか知らないけど）



