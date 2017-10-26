require 'ires/core'
require 'ires/os'

module Ires
  class Service
    class << self

      # Resize image path
      # @return [String]
      def path(path:, width:, height:, mode: 'resize', expire: 30.days)

        os = Ires::Os.current
        return nil if os.nil?

        full_path = image_full_path(path.to_s)

        # if no image or could not find file path then perform the same action as 'image_tag'
        return nil if invalid_path?(full_path)
        
        expiration_date = expiration_date(expire)
        dir = image_dir

        ires_element = {
          path:   full_path,
          width:  width,
          height: height,
          mode:   mode,
          dir:    dir,
          expire: expiration_date
        }
        ires_image_path(ires_element)
      end

      private

      # Image full path or HTTP URL
      # @return [String]
      def image_full_path(path)
        root = Rails.root.to_s
        if path.include?(root) || path.include?('http')
          path
        else
          File.join(image_dir, path)
        end
      end

      # Reszie image directory
      # @return [String]
      def image_dir
        @image_dir ||= Pathname.new(Rails.root).join('public').to_s
      end

      # Check file or URI
      # @return [Bool]
      def invalid_path?(path)
        !File.exist?(path) && !path.include?("http")
      end

      # Expiration date (default: 7.days)
      # ex. "20170101"
      # @return [String]
      def expiration_date(expire)
        (Date.today + expire).strftime('%Y%m%d')
      end

      # Image path
      # @return [String]
      def ires_image_path(ires_element)
        Ires::Core.iresImagePath(
          ires_element[:path],
          ires_element[:width],
          ires_element[:height],
          ires_element[:mode],
          ires_element[:dir],
          ires_element[:expire])
      end

    end
  end
end