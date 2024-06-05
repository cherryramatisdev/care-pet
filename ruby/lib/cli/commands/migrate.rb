require 'dry/cli'

module CLI
  module Commands
    class Migrate < Dry::CLI::Command
      desc "Create the carepet keyspace and tables"

      def call(*)
        # TODO implement...
        puts 'To be implemented'
      end
    end
  end
end
