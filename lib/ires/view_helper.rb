require 'ires/init'
require 'ires/service'
require 'action_view/helpers'

module Ires
  module ViewHelper

    # Image resize
    # @return [image_tag]
    def ires_tag(path:, width:, height:, mode: 'resize', expire: 30.days, **option)
      image = Ires::Service.path(path: path, width: width, height: height, mode: mode, expire: expire)
      return nil if image.nil?

      # Set image_tag
      image_tag(image, option)
    end

  end
end