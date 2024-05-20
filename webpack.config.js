const path = require('path');
// const HtmlWebpackPlugin = require('html-webpack-plugin');
const { VueLoaderPlugin } = require('vue-loader');

module.exports = {
    entry: 
    {
        home: './static/javascript/home.js',
    },
    output: {
        path: path.resolve(__dirname, 'static/dist/javascript'),
        filename: '[name].js'
    },
    module: {
    rules: [{
        test: /\.vue$/,
        use: 'vue-loader'
    },
    {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
    },
    {
        test: /\.(png|jpg|gif|svg)$/,
        use: 'file-loader'
    }
    ]
},
plugins: [
    // new HtmlWebpackPlugin({
    //     template: path.resolve(__dirname, 'app/views/home/index.html')
    // }),
    new VueLoaderPlugin()
],
devServer: {
    static: {
        directory: path.resolve(__dirname, 'static/dist/javascript'),
    },
    compress: false,
    port: 9000
    }
};
