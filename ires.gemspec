$:.push File.expand_path("../lib", __FILE__)

# Maintain your gem's version:
require "ires/version"

# Describe your gem and declare its dependencies:
Gem::Specification.new do |s|
  s.name        = "ires"
  s.version     = Ires::VERSION
  s.authors     = ["enta0701"]
  s.email       = ["endo.takuya.0701@gmail.com"]
  s.homepage    = "https://github.com/endotakuya/ires"
  s.summary     = "Ires is image resizer gem."
  s.description = "Ires is image resizer gem."
  s.license     = "MIT"

  s.files = Dir["{app,config,db,lib}/**/*", "MIT-LICENSE", "Rakefile", "README.md"]
end
