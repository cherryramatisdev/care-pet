require 'dry/cli'
require_relative 'commands/version'
require_relative 'commands/serve'
require_relative 'commands/migrate'
require_relative 'commands/simulate'

module CLI
  extend Dry::CLI::Registry

  register "version", CLI::Commands::Version, aliases: ["v", "-v", "--version"]
  register "migrate", CLI::Commands::Migrate, aliases: ["v"]
  register "serve", CLI::Commands::Serve, aliases: ["s"]
  register "simulate", CLI::Commands::Simulate
end
