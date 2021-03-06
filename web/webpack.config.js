const path = require('path');
const webpack = require('webpack');

module.exports = {
    // context: path.resolve(__dirname, "src"),
    entry: {
        app: './src/js/index.js'
    },
    output: {
        path: __dirname + '/public/js',
        filename: 'app.js'
    },
    plugins: [
        new webpack.ProvidePlugin({
            riot: 'riot'
        }),
        new webpack.DefinePlugin({
            API_SERVER: "'https://mirage.syfm.space'"
        })
    ],
    node: {
        fs: 'empty'
    },
    module: {
        rules: [
            {
                test: /\.sol/,
                loader: 'truffle-solidity-loader'
            },
            {
                test: /\.tag$/,
                enforce: 'pre',
                loader: 'riot-tag-loader'
            },
            {
                test: /\.css$/,
                use: [{ loader: 'style-loader' }, { loader: 'css-loader' }]
            },
            {
                test: /\.js$|\.tag$/,
                enforce: 'post',
                loader: 'babel-loader',
                exclude: /node_modules/,
                query: {
                    presets: ['es2015-riot']
                }
            }
        ]
    },
    devtool: 'inline-source-map',
    devServer: {
        compress: true,
        contentBase: 'public',
        publicPath: '/js',
        port: 3000,
        historyApiFallback: {
            disableDotRule: true
        }
    }
};
