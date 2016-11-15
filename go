#! /usr/bin/env ruby
#
# The preamble down to the check_ruby_version line was generated by version
# 0.1.8 of the go_script gem. This preamble tries to load bundler state
# from Gemfile.lock if present, then the go_script gem for core functionality.
#
# This means the ./go script will run `gem install` or `bundle install` as
# necessary when invoked for the first time, or when dependencies change. It
# also means that commands invoked within the ./go script do not need to be
# prefixed with `bundle exec`.

require 'English'

Dir.chdir File.dirname(__FILE__)

# If a require statement fails, the script calls try_command_and_restart to
# try to install the necessary gems before re-executing itself with the
# original arguments.
def try_command_and_restart(command)
  exit $CHILD_STATUS.exitstatus unless system command

  # If the RUBYOPT environment variable is set, bundler will presume that it
  # has already run. Hence, filtering out RUBYOPT forces bundler to load its
  # state during the re-execution.
  env = {}.merge(ENV)
  env.delete('RUBYOPT')
  exec(env, RbConfig.ruby, *[$PROGRAM_NAME].concat(ARGV))
end

begin
  require 'bundler/setup' if File.exist? 'Gemfile'
rescue LoadError
  try_command_and_restart 'gem install bundler'
rescue SystemExit
  try_command_and_restart 'bundle install'
end

begin
  require 'go_script'
rescue LoadError
  try_command_and_restart 'gem install go_script' unless File.exist? 'Gemfile'
  abort "Please add \"gem 'go_script'\" to your Gemfile"
end

extend GoScript
check_ruby_version '2.2.4'

command_group :dev, 'Development commands'

def_command :update_gems, 'Update Ruby gems' do |gems|
  update_gems gems
end

def_command :test, 'Execute automated tests' do |args|
  exec_cmd "rake test #{args_to_string args}"
end

def_command :lint, 'Run style-checking tools' do |files|
  lint_ruby files
end

def_command :build, 'Test and build the gem' do |args|
  test
  exec_cmd "rake build #{args_to_string args}"
end

def_command :release, 'Test, build, and release a new gem' do
  test
  exec_cmd 'rake release'
end

execute_command ARGV
