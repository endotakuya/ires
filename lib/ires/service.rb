require 'ires/core'
require 'ires/os'
require 'ires/mode'
require 'ires/type'

module Ires
  class Service
    class << self
      # Resize image path
      # @return [String]
      def path(path, width: nil, height: nil, type: Type::ALL, mode: Mode::RESIZE, expire: 30.days)
        raise StandardError, "#{path} is not string." unless path.kind_of?(String)
        raise ArgumentError, "Either width or height is required." if width.nil? && height.nil?
        os = Ires::Os.current
        return nil if os.nil?

        raise StandardError, "Nil location provided. Can't build URI." if path.nil?
        return path if path.empty?

        full_path = image_full_path(path.to_s)

        # if no image or could not find file path then perform the same action as 'image_tag'
        return nil if invalid_path?(full_path)

        expiration_date = expiration_date(expire)
        dir = image_dir

        ires_element = {
          path:   full_path,
          width:  width || 0,
          height: height || 0,
          mode:   mode,
          type:   type,
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
        !File.exist?(path) && !path.include?('http')
      end

      # Expiration date (default: 7.days)
      # ex. "20170101"
      # @return [String]
      def expiration_date(expire)
        (Time.zone.today + expire).strftime('%Y%m%d')
      end

      # Image path
      # @return [String]
      def ires_image_path(ires_element)
        mode = ires_element[:mode]
        ires_element.delete(:mode)
        case mode
        when Mode::RESIZE
          Core.resizeImagePath(*ires_element.values)
        when Mode::CROP
          Core.cropImagePath(*ires_element.values)
        when Mode::RESIZE_TO_CROP
          Core.resizeToCropImagePath(*ires_element.values)
        else
          Core.resizeImagePath(*ires_element.values)
        end
      end
    end
  end
end
