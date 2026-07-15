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


