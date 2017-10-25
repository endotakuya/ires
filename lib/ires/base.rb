module Ires
  module Base
    
    # Image path
    # @return [String]
    def ires_image_path(ires_element)
      Ires::Init.iresImagePath(
        ires_element[:path],
        ires_element[:width],
        ires_element[:height],
        ires_element[:mode],
        ires_element[:dir],
        ires_element[:expire])
    end

  end    
end
