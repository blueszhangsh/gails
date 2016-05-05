var path = require("path");
var webpack = require("webpack");
var StatsPlugin = require("stats-webpack-plugin");
var HtmlWebpackPlugin = require('html-webpack-plugin');
var BowerWebpackPlugin = require("bower-webpack-plugin");


module.exports = function(options) {
    var entry = options.engines.reduce(function(obj, en) {
        obj[en] = path.join(__dirname, "app", "engines", en, "main");
        return obj
    }, {});

    entry.vendor = [
        'jquery',
        'bootstrap'
    ];

    var plugins = [
        // new webpack.ProvidePlugin({
        //   $: "jquery",
        //   jQuery: "jquery"
        // }),
        new webpack.optimize.CommonsChunkPlugin({
            name: 'vendor'
        }),
        new webpack.DefinePlugin({
            VERSION: JSON.stringify(options.version),
            API: JSON.stringify(options.api),
        }),
        new BowerWebpackPlugin()
    ];

    var loaders = [{
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        loader: "babel"
    }, {
        test: /\.(png|jpg|jpeg|gif|ico|svg|ttf|woff|woff2|eot)$/,
        loader: "file"
    }, {
        test: /\.css$/,
        loaders: ['style', 'css']
    }, {
        test: /\.less$/,
        loaders: ['style', 'css', 'less']
    }];

    var htmlOptions = {
        inject: true,
        template: 'app/index.html'
    };
    if (options.compress) {
        htmlOptions.minify = {
            collapseWhitespace: true,
            removeComments: true
        };

        plugins.push(new webpack.optimize.UglifyJsPlugin({
            compress: {
                drop_console: true,
                drop_debugger: true,
                dead_code: true,
                unused: true,

                warnings: false
            },
            output: {
                comments: false
            }
        }));

        plugins.push(new webpack.optimize.DedupePlugin());
        plugins.push(new webpack.optimize.OccurrenceOrderPlugin(true));
        plugins.push(new webpack.DefinePlugin({
            "process.env": {
                NODE_ENV: JSON.stringify("production")
            }
        }));
        plugins.push(new webpack.NoErrorsPlugin());
    } else {
        plugins.push(new StatsPlugin('stats.json', {
            chunkModules: true,
            exclude: [/node_modules/, /bower_components/]
        }));
    }

    options.engines.forEach(function(en) {

        plugins.push(new HtmlWebpackPlugin(Object.assign({},
            htmlOptions, {
                title: en,
                filename: en + ".html",
                chunks: ['vendor', en]
            }
        )));
    });


    return {
        entry: entry,
        output: {
            publicPath: "/",
            path: path.resolve(__dirname, "public"),
            filename: "public/bundle.js"
        },
        module: {
            loaders: loaders
        },
        plugins: plugins,
        output: {
            publicPath: "/",
            path: path.resolve(__dirname, "public"),
            filename: options.compress ? "[id]-[chunkhash].js" : '[name].js'
        },
        devServer: {
            //historyApiFallback: true,
            port: 4200
        }
    };
};
