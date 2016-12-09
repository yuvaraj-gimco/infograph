

export class NavModel {

  getAlertingNav(subPage) {
    return {
      section: {
        title: 'Alerting',
        url: '/plugins',
        icon: 'icon-gf icon-gf-alert'
      },
      navItems: [
        {title: 'Alert List', active: subPage === 0, url: 'alerting/list', icon: 'fa fa-list-ul'},
        {title: 'Notifications', active: subPage === 1, url: 'alerting/notifications', icon: 'fa fa-bell-o'},
      ]
    };
  }

  getDashboardNav(dashboard) {
    return {
      section: {
        title: dashboard.title,
        icon: 'icon-gf icon-gf-dashboard'
      },
      navItems: [
        {title: 'Settings', click: '', icon: 'fa fa-cog'},
        {title: 'Templating', click: '', icon: 'fa fa-code'},
        {title: 'Annotations', click: '', icon: 'fa fa-bolt'},
        {title: 'View JSON', click: '', icon: 'fa fa-eye'},
        {title: 'Save As', click: '', icon: 'fa fa-save'},
      ]
    };
  }
}


var navModel = new NavModel();
export {navModel};
