const path = require('path');
const ManifestPlugin = require('webpack-manifest-plugin');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const IS_DEV_SERVER = process.argv[1].indexOf('webpack-dev-server') >= 0

module.exports = {
  mode: 'development',
  // エントリーポイントの設定
  entry: {
    app: ['./src/js/app.ts', './src/css/app.scss']
  },
  module: {
    rules: [
      {
        test: /\.ts$/, // 拡張子 .ts の場合
        use: "ts-loader", // TypeScript をコンパイルする
      },
      {
        test: /\.scss$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
            options: {
              hmr: IS_DEV_SERVER,
            },
          },
          {
            loader: "css-loader",
            options: {
              // CSS内のurl()メソッドの取り込みを禁止する
              url: false,
              sourceMap: true,
              // ↓に Sass, PostCSSの2つあるので2を指定
              importLoaders: 2
            }
          },
          {
            loader: "postcss-loader",
            options: {
              sourceMap: true,
              plugins: [
                require("autoprefixer")({
                  grid: true
                })
              ]
            }
          },
          {
            loader: "sass-loader",
            options: {
              sourceMap: true
            }
          }
        ]
      }
    ]
  },
  resolve: { // import 文で .ts ファイルを解決する
    extensions: [".ts", ".js"],
  },
  // 出力の設定
  output: {
    // 出力するファイル名
    filename: IS_DEV_SERVER ? '[name].js' : '[name].[hash].js',
    // 出力先のパス（絶対パスを指定する必要がある）
    path: path.join(__dirname, '../public/assets'),
    publicPath: "/assets/",
  },
  plugins: [
    new ManifestPlugin(),
    new MiniCssExtractPlugin({
      filename: IS_DEV_SERVER ? '[name].css' : '[name].[hash].css',
      chunkFilename: IS_DEV_SERVER ? '[id].css' : '[id].[hash].css',
    })
  ],
  devServer: {
    inline: true,
    contentBase: path.join(__dirname, '../public/assets'),
    watchContentBase: true,
  },
};
