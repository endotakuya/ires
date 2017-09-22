require "ffi"
module Ires
  module Service
    extend FFI::Library
    # NOTE: ires.so is golang object
    ffi_lib File.expand_path("../../shared/ires.so", File.dirname(__FILE__))
    
    # resize func
    attach_function :resizeImage,       [:string, :int, :int, :string], :string
    attach_function :cropImage,         [:string, :int, :int, :string], :string
    attach_function :resizeToCropImage, [:string, :int, :int, :string], :string  
  end
end