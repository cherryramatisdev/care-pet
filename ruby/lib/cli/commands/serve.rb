require 'dry/cli'

module CLI
  module Commands
    class Serve < Dry::CLI::Command
      desc "REST API service for tracking pets health state"

      def call(*)
        # TODO implement...
        puts 'To be implemented'
      end
    end
  end
end
