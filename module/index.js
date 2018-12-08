/** EasyWeb spa v3.0.4 data:2018-12-04 */

layui.define(['layer', 'config', 'admin', 'laytpl', 'element', 'form', 'layRouter'], function (exports) {
    var $ = layui.jquery;
    var layer = layui.layer;
    var config = layui.config;
    var admin = layui.admin;
    var element = layui.element;
    var form = layui.form;
    var layRouter = layui.layRouter;
    var bodyDOM = '.layui-layout-admin>.layui-body';
    var tabDOM = bodyDOM + '>.layui-tab';
    var tabFilter = 'admin-pagetabs';
    var navFilter = 'admin-side-nav';
    var titleDOM = bodyDOM + '>.layui-body-header>.layui-body-header-title';

    var index = {
        // 使用递归循环注册
        regRouter: function (menus) {
            $.each(menus, function (i, data) {
                if (data.url && data.url.indexOf('#') == 0) {
                    var hashPath = index.getHashPath(data.url);
                    layRouter.reg(hashPath, function (r) {
                        index.loadView({
                            menuId: r.href,
                            menuPath: config.viewPath + hashPath + '.html',
                            menuName: data.name
                        });
                    });
                }
                if (data.subMenus) {
                    index.regRouter(data.subMenus);
                }
            });
        },
        // 路由加载组件
        loadView: function (param) {
            var menuId = param.menuId;
            var menuPath = param.menuPath;
            var menuName = param.menuName;
            var contentDom = bodyDOM, loadingDOM = bodyDOM;
            var flag;  // 选项卡是否添加
            // 判断是否开启了选项卡功能
            if (config.pageTabs) {
                var tabId = index.getHashPath('#' + menuId);
                $(tabDOM + '>.layui-tab-title>li').each(function (index) {
                    if ($(this).attr('lay-id') === tabId) {
                        flag = true;
                        return false;
                    }
                });
                if (!flag) {
                    element.tabAdd(tabFilter, {
                        title: '<span class="title">' + menuName + '</span>',
                        content: '<div lay-id="' + tabId + '" lay-url="' + menuId + '"></div>',
                        id: tabId
                    });
                } else if (menuId != $(tabDOM + '>.layui-tab-content>.layui-tab-item>div[lay-id="' + tabId + '"]').attr('lay-url')) {
                    $(tabDOM + '>.layui-tab-content>.layui-tab-item>div[lay-id="' + tabId + '"]').attr('lay-url', menuId);
                    flag = false;
                }
                contentDom = tabDOM + '>.layui-tab-content [lay-id="' + tabId + '"]';
                loadingDOM = $(contentDom).parent();
                if (!param.noChange) {
                    element.tabChange(tabFilter, tabId);
                }
                admin.rollPage('auto');
                $(window).resize();
            }
            if (!flag || layRouter.isRefresh) {
                admin.showLoading(loadingDOM);
                $(contentDom).load(menuPath, function () {
                    layRouter.isRefresh = false;
                    element.render('breadcrumb');
                    setTimeout(function () {
                        admin.removeLoading(loadingDOM);
                    }, 300);
                });
            }
            $('.layui-table-tips-c').trigger('click'); // 切换tab关闭表格内浮窗
            admin.activeNav(menuId);
            // 移动设备切换页面隐藏侧导航
            if (admin.getPageWidth() <= 750) {
                admin.flexible(true);
            }
        },
        // 加载主页
        loadHome: function (param) {
            var hashPath = index.getHashPath(param.url);
            if (config.pageTabs) {
                $(bodyDOM).addClass('tab-open');
                index.loadView({
                    menuId: hashPath,
                    menuPath: config.viewPath + hashPath + '.html',
                    menuName: param.name,
                    noChange: true
                });
            }
            layRouter.init({
                index: hashPath,
                notFound: function (r) {
                    config.routerNotFound && config.routerNotFound(r);
                }
            });
        },
        // 打开新页面
        openNewTab: function (param) {
            index.regRouter([{
                name: param.title,
                url: param.url
            }]);
            index.go(param.url.substring(1));
        },
        // 关闭选项卡
        closeTab: function (menuId) {
            element.tabDelete(tabFilter, menuId);
        },
        // 跳转页面
        go: function (hash) {
            layRouter.go(hash);
        },
        // 获取hash的view路径
        getHashPath: function (hash) {
            var layRouter = layui.router(hash);
            var hashPath = '';
            for (var i = 0; i < layRouter.path.length; i++) {
                hashPath += ('/' + layRouter.path[i]);
            }
            return hashPath;
        },
        // 设置Tab标题
        setTabTitle: function (param) {
            var tabId = param.url;
            var title = param.title;
            if (!tabId) {
                tabId = index.getHashPath();
            }
            if (config.pageTabs) {
                $(tabDOM + '>.layui-tab-title>li[lay-id="' + tabId + '"] .title').text(title);
            } else {
                $(titleDOM).text(title);
                $(titleDOM).next('.layui-breadcrumb').find('cite').last().text(title);
            }
        }
    };

    // tab选项卡切换监听
    element.on('tab(' + tabFilter + ')', function (data) {
        var layId = $(this).attr('lay-id');
        index.go($(tabDOM + '>.layui-tab-content>.layui-tab-item>div[lay-id="' + layId + '"]').attr('lay-url'));
    });

    // 监听侧导航栏点击事件
    element.on('nav(' + navFilter + ')', function (elem) {
        var $that = $(elem);
        if ($('.layui-side .layui-nav-tree').attr('lay-accordion') == 'true' && $that.parent().hasClass('layui-nav-item')) {
            if ($that.parent().hasClass('layui-nav-itemed') || $that.parent().hasClass('layui-this')) {
                $('.layui-layout-admin>.layui-side .layui-nav .layui-nav-item').removeClass('layui-nav-itemed');
                $that.parent().addClass('layui-nav-itemed');
            }
        } else {
            admin.setNavHoverCss($that.parentsUntil('.layui-nav-item').parent().children().eq(0));
        }
    });

    // 是否开启footer
    var openFooter = layui.data(config.tableName).openFooter;
    if (openFooter != undefined && openFooter == false) {
        $('body.layui-layout-body').addClass('close-footer');
    }

    // 是否开启多标签
    var openTab = layui.data(config.tableName).openTab;
    if (openTab != undefined) {
        config.pageTabs = openTab;
    }

    exports('index', index);
});
