require 'rbconfig'

module Ires
  module Util
    class << self
      # Reszie image directory
      # return [none(ffi)]
      def current_os
        if ["darwin", "linux"].include?(os)
          return os
        else
          logger.fatal "Ires is not supported by this #{os}"
          return nil
        end
      end
        
      # Search OS
      # return [String]
      def os
        @os ||= (
          host_os = RbConfig::CONFIG['host_os']
          case host_os
          when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            "windows"
          when /darwin|mac os/
            "darwin"
          when /linux/
            "linux"
          when /solaris|bsd/
            "unix"
          else
            "unknown"
          end
        )
      end

    end
  end
end