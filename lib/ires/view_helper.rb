module Ires
  module ViewHelper

    # Image resize
    # return [image_tag]
    def ires_tag(path:, width:, height:, mode:, **option)
      # Reszie image 
      case mode
      when "resize"
        @image = Ires::Service.resizeImage(path.to_s, width, height, image_dir)
      when "crop"
        @image = Ires::Service.cropImage(path.to_s, width, height, image_dir)
      when "resize_to_crop"
        @image = Ires::Service.resizeToCropImage(path.to_s, width, height, image_dir)  
      end

      return nil if @image.nil?

      # Set image_tag
      image_tag(@image, option)
    end

    private
    def image_dir
      # Reszie image directory
      # return [String]
      @image_dir ||= Pathname.new(Rails.root).join("public").to_s
    end

  end
end