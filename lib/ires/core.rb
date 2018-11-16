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
    #   mode:   :int
    #   dir:    :string
    #   expire: :string
    attach_function :iresImagePath, %i[string int int int int string string], :string
  end
end
