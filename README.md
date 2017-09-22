# Ires
`Ires` is image resizer gem.

## Usage

```erb
<!-- Usually -->
<%= ires_tag( path: "image_01.jpg", width: 90, height: 120 ) %>

<!-- Using image_tag options -->
<%= ires_tag( path: "http://example.com/image_02.jpg", width: 200, height: 200, mode: "crop", alt: "example image" ) %>
```

### Select mode

| info                       |     　　　mode 　　　  |
|:---------------------------|:--------------------:|
| Resize                     | resize (default)     |
| Cropping                   | crop                 |
| Rsize after Cropping       | rsize_to_crop        | 

### Saved directory

```
.
└── public
    └── ires
        ├── crop
        │   └── 300x220
        │       └── image_300x220.jpg
        ├── original
        │   └── original
        │       └── image.jpg
        ├── resize
        │   ├── 200x150
        │   │   └── image_200x150.jpg
        │   ├── 300x220
        │   │   └── image_300x220.jpg
        │   ├── 300x400
        │   │   └── image_300x400.jpg
        │   ├── 400x300
        │   │   └── image_400x300.jpg
        │   └── 90x120
        │       └── image_90x120.jpg
        └── resize_to_crop
            └── 300x220
                └── image_300x220.jpg
```

`original` directory where downloaded images are saved.

## Installation
Add this line to your application's Gemfile:

```ruby
gem 'ires'
```

And then execute:
```bash
$ bundle
```

Or install it yourself as:
```bash
$ gem install ires
```

## Development

環境はDockerで準備しています

```shell
$ docker build -t ires:v1 .

# コンテナに入る
$ docker run -it -v $(pwd):/go/src/ires -p 3000:3000 ires-go:v3 /bin/bash
```

## Gemテスト

以下、コンテナ内の作業になります

### 1. Go（shared objectの作成）

パッケージ管理は[dep](https://github.com/golang/dep)を使っています

```shell
# パッケージの依存関係を解決
$ dep ensure

# shared object として出力する
$ go build -buildmode=c-shared -o shared/ires.so main.go 
```

### 2. Railsアプリの起動

```shell
$ test/dummy/bin/rails s -b 0.0.0.0
```

## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
