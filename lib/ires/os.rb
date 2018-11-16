module Ires
  module Os
    class << self
      # Reszie image directory
      # @return [none(ffi)]
      def current
        if %w[darwin linux].include?(os)
          os
        else
          logger.fatal "Ires is not supported by this #{os}"
          nil
        end
      end

      private

      # Search OS
      # @return [String]
      def os
        @os ||= begin
          host_os = RbConfig::CONFIG['host_os']
          case host_os
          when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            'windows'
          when /darwin|mac os/
            'darwin'
          when /linux/
            'linux'
          when /solaris|bsd/
            'unix'
          else
            'unknown'
          end
        end
      end
    end
  end
end
