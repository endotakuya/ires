require 'ires/base'
require 'ires/util'

module Ires
  module Service
    extend Ires::Util
    extend Ires::Base

    # Resize image path
    # @return [String]
    def self.path(path:, width:, height:, mode: 'resize', expire: 30.days)
      full_path = image_full_path(path.to_s)

       # if no image or could not find file path then perform the same action as 'image_tag'
      return nil if target_resource?(full_path)

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

  end
end