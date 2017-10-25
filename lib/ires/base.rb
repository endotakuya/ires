module Ires
  module Base
    
    # Image path
    # @return [String | nil]
    def ires_image_path(ires_element)
      case ires_element[:mode]
      when 'resize'
        @image = Ires::Init.resizeImage(
          ires_element[:path],
          ires_element[:width],
          ires_element[:height],
          ires_element[:dir],
          ires_element[:expire])
      when 'crop'
        @image = Ires::Init.cropImage(
          ires_element[:path],
          ires_element[:width],
          ires_element[:height],
          ires_element[:dir],
          ires_element[:expire])
      when 'resize_to_crop'
        @image = Ires::Init.resizeToCropImage(
          ires_element[:path],
          ires_element[:width],
          ires_element[:height],
          ires_element[:dir],
          ires_element[:expire]) 
      end
      @image ||= nil
    end

  end    
end
