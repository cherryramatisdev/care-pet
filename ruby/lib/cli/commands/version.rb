require 'dry/cli'

module CLI
  module Commands
    class Version < Dry::CLI::Command
      desc "CLI Version"

      def call(*)
        puts "0.0.1"
      end
    end
  end
end
