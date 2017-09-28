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


### Specify cache expiration

Default: **30days**

```erb
<%= ires_tag( path: "/image.jpg", width: 400, height: 300, expire: 7.days ) %>
```

### Saved directory

```
.
└──  public
    ├── image.jpg
    └── ires
        ├── crop
        │   ├── 150x150
        │   │   └── 20171012_image.jpg
        │   ├── 200x170
        │   │   └── 20171019_image.jpg
        │   ├── 400x300
        │   │   └── 20171028_image.jpg
        │   └── 640x480
        │       └── 20171005_image.jpg
        ├── original
        │   └── original
        ├── resize
        │   ├── 150x150
        │   │   └── 20171012_image.jpg
        │   ├── 200x170
        │   │   └── 20171019_image.jpg
        │   ├── 400x300
        │   │   └── 20171028_image.jpg
        │   └── 640x480
        │       └── 20171005_image.jpg
        └── resize_to_crop
            ├── 150x150
            │   └── 20171012_image.jpg
            ├── 200x170
            │   └── 20171019_image.jpg
            ├── 400x300
            │   └── 20171028_image.jpg
            └── 640x480
                └── 20171005_image.jpg
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

## Caution

- It works only with `linux` and `darwin` now.
- Can build only linux（.so）in this docker.


## Development

環境はDockerで準備しています

```shell
$ docker build -t ires:v1 .

# コンテナに入る
$ docker run -it -v $(pwd):/go/src/github.com/endotakuya/ires -p 3000:3000 ires:v1 /bin/bash
```

## Gemテスト

以下、コンテナ内の作業になります

### 1. Go（shared objectの作成）

パッケージ管理は[dep](https://github.com/golang/dep)を使っています

```shell
# パッケージの依存関係を解決
$ dep ensure

# shared object として出力する
$ CGO_ENABLED=1 GOOS=linux go build -v -buildmode=c-shared -o shared/linux/ires.so ext/main.go
```
※ 現状のDockerでは、linux環境のみbuildができます  
※ 他の環境でbuildしたい場合はGCCを追加するか、ホスト側でGoを導入してbuildしてください🙇

### 2. Railsアプリの起動

```shell
$ test/dummy/bin/rails s -b 0.0.0.0
```

## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
