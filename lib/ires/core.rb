require 'ffi'
require 'ires/os'

module Ires
  module Core
    extend FFI::Library
    
    # NOTE: ires.so is golang object
    ffi_lib File.expand_path("../../shared/#{Ires::Os.current}/ires.so", File.dirname(__FILE__))
    
    # resize func
    attach_function :iresImagePath, [:string, :int, :int, :string, :string, :string], :string
  end
end
