require 'ires/core'
require 'ires/service'
require 'ires/mode'
require 'ires/type'
require 'action_view/helpers'

module Ires
  module ViewHelper
    # Image resize
    # @return [image_tag]
    def ires_tag(path:, width: nil, height: nil, type: Type::ALL, mode: Mode::RESIZE, expire: 30.days, **option)
      raise ArgumentError, "Either width or height is required" if width.nil? && height.nil?

      image = Ires::Service.path(
        path: path,
        width: width || 0,
        height: height || 0,
        mode: mode,
        type: type,
        expire: expire
      )
      return nil if image.nil?

      # Set image_tag
      image_tag(image, option)
    end
  end
end
