const path = require('path');
const ManifestPlugin = require('webpack-manifest-plugin');

module.exports = {
  mode: 'development',
  // エントリーポイントの設定
  entry: {
    app: './src/js/app.js',
  },
  // 出力の設定
  output: {
    // 出力するファイル名
    filename: '[name].[contenthash].js',
    // 出力先のパス（絶対パスを指定する必要がある）
    path: path.join(__dirname, '../public/js')
  },
  plugins: [
    new ManifestPlugin()
  ],
};
