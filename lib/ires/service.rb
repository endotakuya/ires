require "ffi"
module Ires
  module Service
    extend FFI::Library
    # NOTE: ires.so is golang object
    lib_path = File.expand_path("../../../ires.so",  __FILE__)
    ffi_lib lib_path
    
    # resize func
    attach_function :resizeImage,       [:string, :int, :int, :string], :string
    attach_function :cropImage,         [:string, :int, :int, :string], :string
    attach_function :resizeToCropImage, [:string, :int, :int, :string], :string  
  end
end