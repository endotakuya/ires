require 'net/http'
require 'action_view/helpers'

module Ires
  module ViewHelper

    # Image resize
    # return [image_tag]
    def ires_tag(path:, width:, height:, mode: 'resize', expire: 30.days, **option)
      full_path = image_full_path(path.to_s)

      # if no image or could not find file path then perform the same action as 'image_tag'
      return image_tag(path, option) if !File.exist?(full_path) && !full_path.include?("http")

      # Expiration date (default: 7.days)
      # ex. "20170101"
      expiration_date = (Date.today + expire).strftime('%Y%m%d')
    
      # Reszie image 
      case mode
      when 'resize'
        @image = Ires::Service.resizeImage(
          full_path,
          width,
          height,
          image_dir,
          expiration_date)
      when 'crop'
        @image = Ires::Service.cropImage(
          full_path,
          width,
          height,
          image_dir,
          expiration_date)
      when 'resize_to_crop'
        @image = Ires::Service.resizeToCropImage(
          full_path,
          width,
          height,
          image_dir,
          expiration_date)  
      end

      return nil if @image.nil?

      # Set image_tag
      image_tag(@image, option)
    end

    private
    # Reszie image directory
    # return [String]
    def image_dir
      @image_dir ||= Pathname.new(Rails.root).join('public').to_s
    end

    def image_full_path(path)
      root = Rails.root.to_s
      if path.include?(root) || path.include?('http')
        path
      else
        File.join(image_dir, path)
      end
    end

  end
end