module Ires
  class Engine < ::Rails::Engine
    ActiveSupport.on_load :action_view do
      ActionView::Base.send(:include, Ires::ViewHelper)
    end
  end
end


# require 'ires/helpers'
# module Ires
#   class Engine < ::Rails::Engine
#     isolate_namespace Ires
#     initializer 'ires.action_view_helpers' do
#       ActiveSupport.on_load :action_view do
#         include Ires::Helpers
#       end
#     end
#   end
# end