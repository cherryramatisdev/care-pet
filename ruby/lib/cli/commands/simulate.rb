require 'dry/cli'

module CLI
  module Commands
    class Simulate < Dry::CLI::Command
      desc "Generate a pet health data and pushes it into the storage"

      def call(*)
        # TODO implement...
        puts 'To be implemented'
      end
    end
  end
end
