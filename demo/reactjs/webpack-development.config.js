module.exports = require("./make-webpack-config")({
    api: 'http://localhost:3000',
    engines: ['auth', 'cms', 'hr', 'ops', 'reading', 'site', 'team'],
    version: '0.0.1'
});
