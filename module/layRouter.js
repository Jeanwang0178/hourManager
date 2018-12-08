/** EasyWeb spa v3.0.4 data:2018-12-04 */

layui.define(function (exports) {
    var Regex = [];
    var viewIndex, lashPath = location.hash;

    var onhashchange = function () {
        var currentRouter = layui.router();
        var view = '';
        for (var i = 0; i < currentRouter.path.length; i++) {
            view += ('/' + currentRouter.path[i]);
        }

        if (!view || view == '/') {
            if (viewIndex) {
                location.replace('#' + viewIndex);
            }
        } else {
            for (var i = 0; i < Regex.length; i++) {
                if (currentRouter.href.match(Regex[i])) {
                    view = Regex[i];
                    break;
                }
            }

            if (router.pop)
                router.pop.call(this, currentRouter);

            router.lash = currentRouter.href;

            if (router[view]) {
                router[view].call(this, currentRouter);
            } else if (router.notFound) {
                router.notFound.call(this, currentRouter);
            }
        }
    };

    var router = {
        isRefresh: false,
        init: function (o) {
            viewIndex = o.index;

            if (o.pop && typeof o.pop == 'function')
                router.pop = o.pop;

            if (o.notFound && typeof o.notFound == 'function')
                router.notFound = o.notFound;

            onhashchange();
            return this;
        },
        reg: function (r, u) {
            if (!r)
                return;

            if (u == undefined)
                u = function () {
                };

            if (r instanceof RegExp) { //正则注册
                router[r] = u;
                Regex.push(r);
            } else if (r instanceof Array) { //数组注册
                for (var i in r) {
                    this.reg.apply(this, [].concat(r[i]).concat(u));
                }
            } else if (typeof r == 'string') { //关键字注册
                if (typeof u == 'function')
                    router[r] = u;
                else if (typeof u == 'string' && router[u])
                    router[r] = router[u];
            }

            return this;
        },
        go: function (u) {
            location.hash = '#' + u;
        },
        refresh: function () {
            router.isRefresh = true;
            onhashchange();
        }
    };

    setInterval(function () {
        if (lashPath != location.hash) {
            lashPath = location.hash;
            onhashchange();
        }
    }, 100);

    exports('layRouter', router);
});
