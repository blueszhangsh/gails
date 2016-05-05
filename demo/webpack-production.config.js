module.exports = require("./make-webpack-config")({
    compress: true,
    engines: ['auth', 'cms', 'hr', 'ops', 'reading', 'site', 'team'],
    api: 'http://www.change-me.com',
    version: '0.0.1'
});