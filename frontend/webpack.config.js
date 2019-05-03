const path = require('path');
const ManifestPlugin = require('webpack-manifest-plugin');
const IS_DEV_SERVER = process.argv[1].indexOf('webpack-dev-server') >= 0

module.exports = {
  mode: 'development',
  // エントリーポイントの設定
  entry: {
    app: './src/js/app.ts',
  },
  module: {
    rules: [
      {
        test: /\.ts$/, // 拡張子 .ts の場合
        use: "ts-loader", // TypeScript をコンパイルする
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
    new ManifestPlugin()
  ],
  devServer: {
    inline: true,
    contentBase: path.join(__dirname, '../public/assets'),
    watchContentBase: true,
  },
};
