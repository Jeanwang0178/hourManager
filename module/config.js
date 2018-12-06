/** EasyWeb spa v3.0.4 data:2018-12-04 */

layui.define(function (exports) {

    var config = {
        base_server: 'json/', // 接口地址，实际项目请换成http形式的地址
        tableName: 'easyweb',  // 存储表名
        pageTabs: false,   // 是否开启多标签
        viewPath: 'components', // 视图位置
        // 获取缓存的token
        getToken: function () {
            return layui.data(config.tableName).token;
        },
        // 清除token
        removeToken: function () {
            layui.data(config.tableName, {
                key: 'token',
                remove: true
            });
        },
        // 缓存token
        putToken: function (token) {
            layui.data(config.tableName, {
                key: 'token',
                value: token
            });
        },
        // 当前登录的用户
        getUser: function () {
            return layui.data(config.tableName).login_user;
        },
        // 缓存user
        putUser: function (user) {
            layui.data(config.tableName, {
                key: 'login_user',
                value: user
            });
        },
        // 获取用户所有权限
        getUserAuths: function () {
            var authorities = config.getUser().authorities;
            var auths = [];
            for (var i = 0; i < authorities.length; i++) {
                auths.push(authorities[i].authority);
            }
            return auths;
        },
        // ajax请求的header
        getAjaxHeaders: function () {
            var headers = [];
            var token = config.getToken();
            headers.push({
                name: 'Authorization',
                value: 'Bearer ' + token.access_token
            });
            return headers;
        },
        // ajax请求结束后的处理，返回false阻止代码执行
        ajaxSuccessBefore: function (res) {
            if (res.code == 401) {
                config.removeToken();
                layer.msg('登录过期', {icon: 2, time: 1500}, function () {
                    location.reload();
                }, 1000);
                return false;
            } else if (res.code == 403) {
                layer.msg('没有访问权限', {icon: 2});
            } else if (res.code == 404) {
                layer.msg('404目标不存在', {icon: 2});
            }
            return true;
        },
        // 路由不存在处理
        routerNotFound: function (r) {
            location.replace('#/template/error/error-404');
        }
    };

    exports('config', config);
});
