require 'ffi'
require 'ires/os'

module Ires
  module Core
    extend FFI::Library
    # NOTE: ires.so is golang object
    ffi_lib File.expand_path("../../shared/#{Ires::Os.current}/ires.so", File.dirname(__FILE__))

    # resize func
    # Type:
    #   path:   :string
    #   width:  :int
    #   height: :int
    #   type:   :int
    #   dir:    :string
    #   expire: :string
    params = %i[string int int int string string]
    attach_function :resizeImagePath, params, :string
    attach_function :cropImagePath, params, :string
    attach_function :resizeToCropImagePath, params, :string
  end
end
