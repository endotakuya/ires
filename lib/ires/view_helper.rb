module Ires
  module ViewHelper

    # Image resize
    # return [image_tag]
    def ires_tag(path:, width:, height:, mode: "resize", **option)
      full_path = image_full_path(path.to_s)

      # Reszie image 
      case mode
      when "resize"
        @image = Ires::Service.resizeImage(full_path, width, height, image_dir)
      when "crop"
        @image = Ires::Service.cropImage(full_path, width, height, image_dir)
      when "resize_to_crop"
        @image = Ires::Service.resizeToCropImage(full_path, width, height, image_dir)  
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

    def image_full_path(path)
      root = Rails.root.to_s
      if path.include?(root) || path.include?("http")
        return path
      else
        return File.join(image_dir, path)
      end
    end

  end
end