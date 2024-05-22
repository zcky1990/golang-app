const path = require('path');
// const HtmlWebpackPlugin = require('html-webpack-plugin');

const ENV = process.env.NODE_ENV || 'development';

// Set the output directory based on the environment
const outputDirectory = ENV === 'production' ? 'static/dist/javascript' : 'static/dev/javascript';

const { VueLoaderPlugin } = require('vue-loader');

module.exports = {
    entry: 
    {
        home: './static/javascript/home.js',
    },
    output: {
        path: path.resolve(__dirname, outputDirectory),
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
    //use this plugin if you want to generate html file
    // new HtmlWebpackPlugin({
    //     template: path.resolve(__dirname, 'app/views/home/index.html')
    // }),
    new VueLoaderPlugin()
],

//we are using watch option instead dev server, so we only need to tun our go app, then whenever any changes happen in our js it automaticaly build file
//see build:dev and build:prod on package.json
//to run npm run build:dev or npm run build:prod
watch: true,
watchOptions: {
    ignored: /node_modules/,
},
//we didnt use devServer to hotreload
// devServer: {
//     static: {
//         directory: path.resolve(__dirname, 'dist'),
//     },
//     compress: false,
//     port: 9000
//     }
};
