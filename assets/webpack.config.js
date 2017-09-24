var path = require('path');
var webpack = require('webpack');

// HTML
/*const HtmlWebpackPlugin = require('html-webpack-plugin');*/
//const HtmlWebpackPluginConfig = new HtmlWebpackPlugin({
    //template: './src/index.html',
    //filename: 'index.html',
    //inject: 'body'
/*})*/

module.exports = {
    devtool: 'source-map',
    entry: [
        './src/index.js'
    ],
    output: {
        path: path.join(__dirname, 'dist'),
        filename: 'bundle.js',
    },
    module: {
        loaders: [{
            test: /.jsx?$/,
            loader: 'babel-loader',
            include: [
                path.join(__dirname, 'src')
            ],
            exclude: /node_modules/,
            query: {
                presets: ['es2015', 'react']
            }
        }]
    }
    //plugins: [HtmlWebpackPluginConfig]
};
