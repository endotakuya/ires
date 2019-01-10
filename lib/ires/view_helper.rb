require 'ires/core'
require 'ires/service'
require 'ires/mode'
require 'ires/type'
require 'action_view/helpers'

module Ires
  module ViewHelper
    # Image resize
    # @return [image_tag]
    def ires_tag(path, width: nil, height: nil, type: Type::ALL, mode: Mode::RESIZE, expire: 30.days, **option)
      image_path = Ires::Service.path(
        path,
        width: width || 0,
        height: height || 0,
        mode: mode,
        type: type,
        expire: expire
      )

      # Set image_tag
      image_tag(image_path, option)
    end
  end
end
