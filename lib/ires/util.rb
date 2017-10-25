require 'rbconfig'

module Ires
  module Util

    # Reszie image directory
    # return [String]
    def image_dir
      @image_dir ||= Pathname.new(Rails.root).join('public').to_s
    end

    # Image full path or HTTP URL
    # return [String]
    def image_full_path(path)
      root = Rails.root.to_s
      if path.include?(root) || path.include?('http')
        path
      else
        File.join(image_dir, path)
      end
    end

    # Reszie image directory
    # return [none(ffi)]
    def current_os
      if ['darwin', 'linux'].include?(os)
        os
      else
        logger.fatal "Ires is not supported by this #{os}"
        nil
      end
    end

    # Search OS
    # return [String]
    def os
      @os ||= (
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
      )
    end

    # Expiration date (default: 7.days)
    # ex. "20170101"
    # return [String]
    def expiration_date(expire)
      (Date.today + expire).strftime('%Y%m%d')
    end

    # Check file or URI
    # return [Bool]
    def target_resource?(path)
      !File.exist?(path) && !path.include?("http")
    end
    
  end
end