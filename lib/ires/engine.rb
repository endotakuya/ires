module Ires
  class Engine < ::Rails::Engine
    ActiveSupport.on_load :action_view do
      ActionView::Base.send(:include, Ires::ViewHelper)
    end
  end
end
